package main

import (
	"backend/internal/infrastructure/models"
	"backend/internal/infrastructure/repositories"
	"backend/internal/infrastructure/seeder"
	"backend/internal/interfaces/controllers"
	"backend/internal/usecases"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"backend/internal/infrastructure/database"
	"backend/internal/infrastructure/router"
	"backend/internal/infrastructure/server"
)

// @title Simple Blog API
// @version 1.0
// @description This is a simple blog API.
// @host localhost:8080
// @BasePath /
func main() {
	log.SetOutput(os.Stdout)
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load(".env")
	}

	db, err := database.InitDB()
	if err != nil {
		panic("データベース接続失敗: " + err.Error())
	}

	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Track{}, &models.Tag{}, &models.PostTag{}, &models.Like{}); err != nil {
		log.Fatalf("マイグレーション失敗: %v", err)
	}

	if err := seeder.Seed(db); err != nil {
		log.Fatalf("シード注入失敗: %v", err)
	}

	postRepo := repositories.NewPostRepository(db)
	postService := usecases.NewPostUsecase(postRepo)
	postHandler := controllers.NewPostController(postService)

	userRepo := repositories.NewUserRepository(db)
	userService := &usecases.UserUsecase{Repo: userRepo}
	userHandler := controllers.NewUserController(userService)

	tagRepo := repositories.NewTagRepository(db)
	tagService := usecases.NewTagUsecase(tagRepo, postRepo)
	tagHandler := controllers.NewTagController(tagService)

	spotifyService := usecases.NewSpotifyUsecase()
	spotifyHandler := controllers.NewSpotifyController(spotifyService)

	likeRepo := repositories.NewLikeRepository(db)
	likeService := usecases.NewLikeUsecase(likeRepo, postRepo)
	likeHandler := controllers.NewLikeController(likeService)

	oauthService := usecases.NewOAuthUsecase(userRepo)
	oauthHandler := controllers.NewOAuthController(oauthService)

	e := echo.New()
	router.InitRouter(e, *postHandler, *userHandler, *tagHandler, *spotifyHandler, likeHandler, *oauthHandler)

	server.StartServer(e)
}
