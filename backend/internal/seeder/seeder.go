package seeder

import (
<<<<<<< HEAD
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
=======
	"fmt"
	"log"
	"math/rand"
	"net/url"

	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"simple-blog/backend/internal/model"
)

const SeedValue int64 = 20240801

func Seed(db *gorm.DB) error {
	log.Println("🌱 Clearing existing data...")

	if err := db.Exec("TRUNCATE TABLE tracks, posts, users, post_tags, tags RESTART IDENTITY CASCADE").Error; err != nil {
		return err
	}

	r := rand.New(rand.NewSource(SeedValue))
	gf := gofakeit.New(uint64(SeedValue))

	log.Println("🌱 Running deterministic seeding...")

	{
		hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		demo := model.User{
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
>>>>>>> develop
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

<<<<<<< HEAD
		// ユーザー生成
		user := model.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
=======
		user := model.User{
			Name:     name,
			Email:    email,
>>>>>>> develop
			Password: string(hashedPassword),
		}
		if err := db.Create(&user).Error; err != nil {
			return err
		}

<<<<<<< HEAD
		// ポスト生成
		post := model.Post{
			Title:       faker.Sentence(),
			Description: faker.Paragraph(),
=======
		post := model.Post{
			Title:       gf.Sentence(6),
			Description: gf.Paragraph(1, 3, 12, " "),
>>>>>>> develop
			UserID:      user.ID,
		}
		if err := db.Create(&post).Error; err != nil {
			return err
		}

<<<<<<< HEAD
		// トラック生成（3〜5曲）
		numTracks := rand.Intn(3) + 3 // 3〜5
		for j := 0; j < numTracks; j++ {
			track := model.Track{
				SpotifyID:     faker.UUIDDigit(),
				Name:          faker.Word(),
				ArtistName:    faker.Name(),
				AlbumImageUrl: faker.URL(),
=======
		numTracks := r.Intn(3) + 3
		for j := 0; j < numTracks; j++ {
			seed := fmt.Sprintf("%d-%d-%d", user.ID, i, j)
			img := fmt.Sprintf("https://picsum.photos/seed/%s/300/300", url.PathEscape(seed))

			track := model.Track{
				SpotifyID:     gf.UUID(),
				Name:          gf.Word(),
				ArtistName:    gf.Name(),
				AlbumImageUrl: img,
>>>>>>> develop
				PostID:        post.ID,
			}
			if err := db.Create(&track).Error; err != nil {
				return err
			}
		}
	}

<<<<<<< HEAD
	log.Println("✅ Seeding completed with 50 users")
=======
	log.Println("✅ Deterministic seeding completed")
>>>>>>> develop
	return nil
}
