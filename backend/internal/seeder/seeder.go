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
	if len(users) < 2 {
		return fmt.Errorf("need at least two users to seed communities")
	}
	communityCount := 12
	communities := make([]model.CommunityModel, 0, communityCount)
	messageTemplates := []string{
		"みなさんこんにちは！最近聴いているおすすめ曲を共有しましょう。",
		"次のセッションは週末の夜でどうでしょうか？",
		"ギターのコード進行で悩んでいるのでアドバイス欲しいです。",
		"録音したデモ音源をアップしました。感想をください！",
		"ライブ情報をまとめたので確認をお願いします。",
	}
	for i := 0; i < communityCount; i++ {
		creator := users[r.Intn(len(users))]
		createdAt := time.Now().Add(-time.Duration((communityCount-i)*6) * time.Hour)
		community := model.CommunityModel{
			ID:          uuid.NewString(),
			Name:        fmt.Sprintf("%sコミュニティ", gf.Company()),
			Description: "バンド活動や音楽制作について話し合うためのコミュニティです。",
			CreatorID:   creator.ID,
			CreatedAt:   createdAt,
		}
		if err := db.Create(&community).Error; err != nil {
			return err
		}
		communities = append(communities, community)
	}

	participants := []model.UserModel{users[0], users[1]}
	timeOffsets := []time.Duration{0, 5 * time.Minute, 10 * time.Minute, 15 * time.Minute, 20 * time.Minute}
	for _, community := range communities {
		base := community.CreatedAt
		for idx, content := range messageTemplates {
			sender := participants[idx%len(participants)]
			createdAt := base.Add(timeOffsets[idx])
			message := model.MessageModel{
				ID:          uuid.NewString(),
				CommunityID: community.ID,
				SenderID:    sender.ID,
				Content:     content,
				CreatedAt:   createdAt,
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
