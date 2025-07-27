package main

import (
	"simple-blog/backend/internal/handler"
	"simple-blog/backend/internal/middleware"
	"simple-blog/backend/internal/model"
	"simple-blog/backend/internal/repository"
	"simple-blog/backend/internal/seeder"
	"simple-blog/backend/internal/service"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Simple Blog API
// @version 1.0
// @description This is a simple blog API.
// @host localhost:8080
// @BasePath /
func main() {
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load(".env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	var db *gorm.DB
	var err error
	maxRetries := 10

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("DB接続に成功しました")
			break
		}
		fmt.Printf("DB接続リトライ中... (%d/%d): %v\n", i+1, maxRetries, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		panic("データベース接続失敗: " + err.Error())
	}

	if err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Track{}, &model.Tag{}, &model.PostTag{}); err != nil {
		log.Fatalf("マイグレーション失敗: %v", err)
	}

	if err := seeder.Seed(db); err != nil {
		log.Fatalf("シード注入失敗: %v", err)
	}

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userHandler := &handler.UserHandler{Service: userService}

	tagRepo := repository.NewTagRepository(db)
	tagService := service.NewTagService(tagRepo, postRepo)
	tagHandler := handler.NewTagHandler(tagService)

	spotifyService := service.NewSpotifyService()
	spotifyHandler := handler.NewSpotifyHandler(spotifyService)

	likeRepo := repository.NewLikeRepository(db)
	likeService := service.NewLikeService(likeRepo, postRepo)
	likeHandler := handler.NewLikeHandler(likeService)

	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.PATCH, echo.DELETE},
		AllowCredentials: true,
	}))

	e.POST("/auth/login", userHandler.Login)
	e.POST("/auth/register", userHandler.Register)

	authGroup := e.Group("")
	authGroup.Use(middleware.AuthMiddleware(middleware.AuthMiddlewareConfig{}))

	authGroup.POST("/auth/refresh", userHandler.RefreshToken)
	authGroup.GET("/auth/me", userHandler.GetMe)

	authGroup.GET("/posts", postHandler.GetAllPosts)
	authGroup.GET("/posts/:id", postHandler.GetPostByID)
	authGroup.GET("/posts/page/:page", postHandler.GetPostsByPage)
	authGroup.POST("/blogs", postHandler.CreatePost)
	authGroup.PUT("/blogs/:id", postHandler.UpdatePost)
	authGroup.DELETE("/blogs/:id", postHandler.DeletePost)

	// User routes (authentication required)
	authGroup.GET("/users", userHandler.GetAllUsers)
	authGroup.GET("/users/:id", userHandler.GetUserByID)
	authGroup.GET("/users/:id/posts", userHandler.GetUserPosts)

	// Like routes (authentication required)
	authGroup.POST("/posts/:postID/like", likeHandler.LikePost)
	authGroup.DELETE("/posts/:postID/unlike", likeHandler.UnlikePost)
	authGroup.GET("/posts/:postID/liked", likeHandler.IsPostLikedByUser)

	// Spotify routes (authentication required)
	authGroup.GET("/spotify/search", spotifyHandler.SearchTracks)

	// Tag routes (authentication required)
	tagGroup := authGroup.Group("/tags")
	tagGroup.POST("", tagHandler.CreateTag)
	tagGroup.GET("", tagHandler.GetAllTags)
	tagGroup.GET("/:id", tagHandler.GetTagByID)
	tagGroup.PUT("/:id", tagHandler.UpdateTag)
	tagGroup.DELETE("/:id", tagHandler.DeleteTag)
	tagGroup.POST("/posts/:postID", tagHandler.AddTagsToPost)
	tagGroup.DELETE("/posts/:postID", tagHandler.RemoveTagsFromPost)
	tagGroup.GET("/posts/:postID", tagHandler.GetTagsByPostID)

	e.Logger.Fatal(e.Start(":8080"))
}