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

		post := model.PostModel{
			Title:       gf.Sentence(6),
			Description: gf.Paragraph(1, 3, 12, " "),
			UserID:      user.ID,
		}
		if err := db.Create(&post).Error; err != nil {
			return err
		}

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
