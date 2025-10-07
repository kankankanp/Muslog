package seeder

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"time"

	"github.com/kankankanp/Muslog/internal/infrastructure/model"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const SeedValue int64 = 20240801

func Seed(db *gorm.DB) error {
	log.Println("Clearing existing data...")

	// 現在のテーブル名に合わせて初期化（従来名 + many2many: post_tags）
	if err := db.Exec("TRUNCATE TABLE band_applications, band_recruitments, post_tags, tracks, posts, users, tags, likes, messages, communities RESTART IDENTITY CASCADE").Error; err != nil {
		return err
	}

	r := rand.New(rand.NewSource(SeedValue))
	gf := gofakeit.New(uint64(SeedValue))

	log.Println("Running deterministic seeding...")

	const guestUUID = "c88793d0-7afd-4acc-b15c-9a11dd4382a0"

	users := make([]model.UserModel, 0, 60)
	var tagModels []model.TagModel
	posts := make([]model.PostModel, 0, 50)

	{
		hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		demo := model.UserModel{
			ID:       guestUUID,
			Name:     "ゲストユーザー",
			Email:    "user@example.com",
			Password: string(hashed),
		}
		if err := db.Where("email = ?", demo.Email).FirstOrCreate(&demo).Error; err != nil {
			return err
		}
		users = append(users, demo)
	}

	tagNames := []string{"ロック", "ポップス", "ジャズ", "メタル", "ヒップホップ", "クラシック", "エレクトロ", "アコースティック"}
	tagModels = make([]model.TagModel, len(tagNames))
	for i, name := range tagNames {
		tagModels[i] = model.TagModel{Name: name}
	}
	if err := db.Create(&tagModels).Error; err != nil {
		return err
	}

	for i := 0; i < 50; i++ {
		name := gf.Name()
		email := gf.Email()

		// パスワードのハッシュはソルトで毎回変わる（正常）。
		// 平文は固定にしておく（必要なら "password-XX" などに）
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := model.UserModel{
			ID:       uuid.NewString(),
			Name:     name,
			Email:    email,
			Password: string(hashedPassword),
		}
		if err := db.Create(&user).Error; err != nil {
			return err
		}
		users = append(users, user)

		tagCount := r.Intn(3) + 1
		indices := r.Perm(len(tagModels))[:tagCount]
		selectedTags := make([]model.TagModel, 0, tagCount)
		for _, idx := range indices {
			selectedTags = append(selectedTags, tagModels[idx])
		}

		post := model.PostModel{
			Title:       gf.Sentence(6),
			Description: gf.Paragraph(1, 3, 12, " "),
			UserID:      user.ID,
			Tags:        selectedTags,
		}
		if err := db.Create(&post).Error; err != nil {
			return err
		}
		posts = append(posts, post)

		numTracks := r.Intn(3) + 3
		for j := 0; j < numTracks; j++ {
			seed := fmt.Sprintf("%s-%d-%d", user.ID, i, j)
			img := fmt.Sprintf("https://picsum.photos/seed/%s/300/300", url.PathEscape(seed))

			track := model.TrackModel{
				SpotifyID:     gf.UUID(),
				Name:          gf.Word(),
				ArtistName:    gf.Name(),
				AlbumImageUrl: img,
				PostID:        post.ID,
			}
			if err := db.Create(&track).Error; err != nil {
				return err
			}
		}
	}

	log.Println("Seeding likes...")
	if len(posts) > 0 {
		likeTargets := len(posts) * 3
		likeKeys := make(map[string]struct{}, likeTargets)
		postLikeCounts := make(map[uint]int)
		attempts := 0
		for len(likeKeys) < likeTargets && attempts < likeTargets*5 {
			attempts++
			post := posts[r.Intn(len(posts))]
			user := users[r.Intn(len(users))]
			key := fmt.Sprintf("%s-%d", user.ID, post.ID)
			if _, exists := likeKeys[key]; exists {
				continue
			}
			if err := db.Create(&model.LikeModel{PostID: post.ID, UserID: user.ID}).Error; err != nil {
				return err
			}
			likeKeys[key] = struct{}{}
			postLikeCounts[post.ID]++
		}
		for postID, cnt := range postLikeCounts {
			if err := db.Model(&model.PostModel{}).Where("id = ?", postID).Update("likes_count", cnt).Error; err != nil {
				return err
			}
		}
	}

	log.Println("Seeding communities and messages...")
	communityCount := 12
	communities := make([]model.CommunityModel, 0, communityCount)
	for i := 0; i < communityCount; i++ {
		creator := users[r.Intn(len(users))]
		community := model.CommunityModel{
			ID:          uuid.NewString(),
			Name:        fmt.Sprintf("%sコミュニティ", gf.Company()),
			Description: gf.Paragraph(1, 2, 18, " "),
			CreatorID:   creator.ID,
			CreatedAt:   time.Now().Add(-time.Duration(r.Intn(1440)) * time.Minute),
		}
		if err := db.Create(&community).Error; err != nil {
			return err
		}
		communities = append(communities, community)
	}
	for _, community := range communities {
		messageCount := r.Intn(10) + 5
		for j := 0; j < messageCount; j++ {
			sender := users[r.Intn(len(users))]
			message := model.MessageModel{
				ID:          uuid.NewString(),
				CommunityID: community.ID,
				SenderID:    sender.ID,
				Content:     gf.Sentence(12),
				CreatedAt:   time.Now().Add(-time.Duration(r.Intn(720)) * time.Hour),
			}
			if err := db.Create(&message).Error; err != nil {
				return err
			}
		}
	}

	log.Println("Seeding band recruitments...")
	if len(users) == 0 {
		return fmt.Errorf("no users available for band recruitment seeding")
	}

	genres := []string{"ロック", "J-POP", "ジャズ", "メタル", "フォーク", "エレクトロ"}
	locations := []string{"東京", "大阪", "名古屋", "福岡", "札幌", "オンライン"}
	skillLevels := []string{"初心者歓迎", "中級以上", "経験3年以上"}
	partsPool := [][]string{
		{"Vo", "Gt"},
		{"Gt", "Ba"},
		{"Gt", "Dr"},
		{"Key", "Vo"},
		{"Ba", "Dr"},
		{"Vo"},
		{"Gt"},
	}

	for i := 0; i < 25; i++ {
		creator := users[r.Intn(len(users))]
		parts := partsPool[r.Intn(len(partsPool))]
		deadlineOffset := r.Intn(60) - 20
		deadline := time.Now().AddDate(0, 0, deadlineOffset)
		realDeadline := &deadline
		if r.Intn(4) == 0 {
			realDeadline = nil
		}
		status := "open"
		if deadlineOffset < 0 && r.Intn(3) == 0 {
			status = "closed"
		}

		recruitment := model.BandRecruitmentModel{
			ID:              uuid.NewString(),
			Title:           gf.Sentence(4),
			Description:     gf.Paragraph(1, 2, 20, " "),
			Genre:           genres[r.Intn(len(genres))],
			Location:        locations[r.Intn(len(locations))],
			RecruitingParts: pq.StringArray(parts),
			SkillLevel:      skillLevels[r.Intn(len(skillLevels))],
			Contact:         fmt.Sprintf("%s@example.com", gf.Word()),
			Deadline:        realDeadline,
			Status:          status,
			UserID:          creator.ID,
		}
		if err := db.Create(&recruitment).Error; err != nil {
			return err
		}

		applications := r.Intn(4)
		createdApplicants := make(map[string]struct{})
		for j := 0; j < applications; j++ {
			applicant := users[r.Intn(len(users))]
			if applicant.ID == creator.ID {
				continue
			}
			if _, exists := createdApplicants[applicant.ID]; exists {
				continue
			}
			createdApplicants[applicant.ID] = struct{}{}
			application := model.BandApplicationModel{
				ID:                uuid.NewString(),
				BandRecruitmentID: recruitment.ID,
				ApplicantID:       applicant.ID,
				Message:           gf.Sentence(10),
			}
			if err := db.Create(&application).Error; err != nil {
				return err
			}
		}
	}

	log.Println("✅ Deterministic seeding completed")
	return nil
}
