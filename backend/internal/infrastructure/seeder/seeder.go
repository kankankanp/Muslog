package seeder

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"

	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"backend/internal/infrastructure/models"
)

const SeedValue int64 = 20240801

func Seed(db *gorm.DB) error {
	log.Println("Clearing existing data...")

	if err := db.Exec("TRUNCATE TABLE tracks, posts, users, post_tags, tags RESTART IDENTITY CASCADE").Error; err != nil {
		return err
	}

	r := rand.New(rand.NewSource(SeedValue))
	gf := gofakeit.New(uint64(SeedValue))

	log.Println("Running deterministic seeding...")

	{
		hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		demo := models.User{
			Name:     "ゲストユーザー",
			Email:    "user@example.com",
			Password: string(hashed),
		}
		if err := db.Where("email = ?", demo.Email).FirstOrCreate(&demo).Error; err != nil {
			return err
		}
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

		user := models.User{
			Name:     name,
			Email:    email,
			Password: string(hashedPassword),
		}
		if err := db.Create(&user).Error; err != nil {
			return err
		}

		post := models.Post{
			Title:       gf.Sentence(6),
			Description: gf.Paragraph(1, 3, 12, " "),
			UserID:      user.ID,
		}
		if err := db.Create(&post).Error; err != nil {
			return err
		}

		numTracks := r.Intn(3) + 3
		for j := 0; j < numTracks; j++ {
			seed := fmt.Sprintf("%d-%d-%d", user.ID, i, j)
			img := fmt.Sprintf("https://picsum.photos/seed/%s/300/300", url.PathEscape(seed))

			track := models.Track{
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

	log.Println("✅ Deterministic seeding completed")
	return nil
}
