package seeder

import (
	"simple-blog/backend/internal/model"
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	log.Println("🌱 Clearing existing data...")

	if err := db.Exec("DELETE FROM tracks").Error; err != nil {
		return err
	}
	if err := db.Exec("DELETE FROM posts").Error; err != nil {
		return err
	}
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		return err
	}

	log.Println("🌱 Running random seeding...")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 50; i++ {
		// パスワードのハッシュ化
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// ユーザー生成
		user := model.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: string(hashedPassword),
		}
		if err := db.Create(&user).Error; err != nil {
			return err
		}

		// ポスト生成
		post := model.Post{
			Title:       faker.Sentence(),
			Description: faker.Paragraph(),
			UserID:      user.ID,
		}
		if err := db.Create(&post).Error; err != nil {
			return err
		}

		// トラック生成（3〜5曲）
		numTracks := rand.Intn(3) + 3 // 3〜5
		for j := 0; j < numTracks; j++ {
			track := model.Track{
				SpotifyID:     faker.UUIDDigit(),
				Name:          faker.Word(),
				ArtistName:    faker.Name(),
				AlbumImageUrl: faker.URL(),
				PostID:        post.ID,
			}
			if err := db.Create(&track).Error; err != nil {
				return err
			}
		}
	}

	log.Println("✅ Seeding completed with 50 users")
	return nil
}
