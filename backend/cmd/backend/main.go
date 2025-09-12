package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	appConfig "github.com/kankankanp/Muslog/config"
	"github.com/kankankanp/Muslog/internal/adapter/handler"
    "github.com/kankankanp/Muslog/internal/infrastructure/model"
    dblogger "github.com/kankankanp/Muslog/internal/infrastructure/logger"
	"github.com/kankankanp/Muslog/internal/infrastructure/repository"
	"github.com/kankankanp/Muslog/internal/middleware"
	"github.com/kankankanp/Muslog/internal/seeder"
	"github.com/kankankanp/Muslog/internal/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"gorm.io/driver/postgres"
    "gorm.io/gorm"
    glogger "gorm.io/gorm/logger"
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

	// Load application configuration
	cfg, err := appConfig.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load application configuration: %v", err)
	}

	// Use loaded config for DB connection
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

    var db *gorm.DB
    baseLogger := glogger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), glogger.Config{
        SlowThreshold:             200 * time.Millisecond,
        LogLevel:                  glogger.Warn,
        IgnoreRecordNotFoundError: true,
        Colorful:                  false,
    })
    filteredLogger := dblogger.NewCancelFilter(baseLogger)
    maxRetries := 10
    for i := 0; i < maxRetries; i++ {
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: filteredLogger})
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

	// Ensure required PostgreSQL extensions
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatalf("failed to create uuid-ossp extension: %v", err)
	}
	if err := db.AutoMigrate(
		&model.UserModel{},
		&model.PostModel{},
		&model.TrackModel{},
		&model.TagModel{},
		&model.PostTagModel{},
		&model.LikeModel{},
		&model.MessageModel{},
		&model.CommunityModel{},
	); err != nil {
		log.Fatalf("マイグレーション失敗: %v", err)
	}

	if err := seeder.Seed(db); err != nil {
		log.Fatalf("シード注入失敗: %v", err)
	}

	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(cfg.S3Region))
	if err != nil {
		log.Fatalf("failed to load AWS SDK config: %v", err)
	}
	s3Client := s3.NewFromConfig(awsCfg)

	postRepo := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepo)
	postHandler := handler.NewPostHandler(postUsecase)

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, postRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	tagRepo := repository.NewTagRepository(db)
	tagUsecase := usecase.NewTagUsecase(tagRepo, postRepo)
	tagHandler := handler.NewTagHandler(tagUsecase)

	spotifyUsecase := usecase.NewSpotifyUsecase()
	spotifyHandler := handler.NewSpotifyHandler(spotifyUsecase)

	likeRepo := repository.NewLikeRepository(db)
	likeUsecase := usecase.NewLikeUsecase(likeRepo, postRepo)
	likeHandler := handler.NewLikeHandler(likeUsecase)

	oauthUsecase := usecase.NewOAuthUsecase(userRepo)
	oauthHandler := handler.NewOAuthHandler(oauthUsecase)

	messageRepo := repository.NewMessageRepository(db)
	messageUsecase := usecase.NewMessageUsecase(messageRepo)
	messageHandler := handler.NewMessageHandler(messageUsecase)

	communityRepo := repository.NewCommunityRepository(db)
	communityUsecase := usecase.NewCommunityUsecase(communityRepo)
	communityHandler := handler.NewCommunityHandler(communityUsecase)

	imageUsecase := usecase.NewImageUsecase(s3Client, cfg.S3BucketName, cfg.S3Region)
	imageHandler := handler.NewImageHandler(imageUsecase, userRepo, postRepo)

	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	public := e.Group("/api/v1")
	public.POST("/auth/login", userHandler.Login)
	public.POST("/auth/register", userHandler.Register)
	public.GET("/auth/google", oauthHandler.GetGoogleAuthURL)
	public.GET("/auth/google/callback", oauthHandler.GoogleCallback)
	public.GET("/spotify/search", spotifyHandler.SearchTracks)
	public.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	public.GET("/posts", postHandler.GetAllPosts)
	public.GET("/posts/:id", postHandler.GetPostByID)
	public.GET("/posts/page/:page", postHandler.GetPostsByPage)

	protected := e.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(middleware.AuthMiddlewareConfig{
		Skipper: func(c echo.Context) bool {
			return c.Request().Method == http.MethodOptions
		},
	}))

	protected.POST("/refresh", userHandler.RefreshToken)
	protected.POST("/logout", userHandler.Logout)
	protected.GET("/me", userHandler.GetMe)

	// posts
	postGroup := protected.Group("/posts")
	postGroup.POST("", postHandler.CreatePost)
	postGroup.PUT("/:id", postHandler.UpdatePost)
	postGroup.DELETE("/:id", postHandler.DeletePost)
	postGroup.GET("/search", postHandler.SearchPosts)

	// likes
	postGroup.POST("/:postID/like", likeHandler.LikePost)
	postGroup.DELETE("/:postID/unlike", likeHandler.UnlikePost)

	// users
	userGroup := protected.Group("/users")
	userGroup.GET("", userHandler.GetAllUsers)
	userGroup.GET("/:id", userHandler.GetUserByID)
	userGroup.GET("/:id/posts", userHandler.GetUserPosts)
	userGroup.POST("/:userId/profile-image", imageHandler.UploadProfileImage)

	// tags
	tagGroup := protected.Group("/tags")
	tagGroup.POST("", tagHandler.CreateTag)
	tagGroup.GET("", tagHandler.GetAllTags)
	tagGroup.GET("/:id", tagHandler.GetTagByID)
	tagGroup.PUT("/:id", tagHandler.UpdateTag)
	tagGroup.DELETE("/:id", tagHandler.DeleteTag)
	tagGroup.POST("/posts/:postID", tagHandler.AddTagsToPost)
	tagGroup.DELETE("/posts/:postID", tagHandler.RemoveTagsFromPost)
	tagGroup.GET("/posts/:postID", tagHandler.GetTagsByPostID)

	// community
	communityGroup := protected.Group("/communities")
	communityGroup.GET("", communityHandler.GetAllCommunities)
	communityGroup.POST("", communityHandler.CreateCommunity)
	communityGroup.GET("/search", communityHandler.SearchCommunities)
	communityGroup.GET("/:communityId/messages", messageHandler.GetMessagesByCommunityID)

	// image
	protected.POST("/posts/:postId/header-image", imageHandler.UploadPostHeaderImage)
	protected.POST("/images/upload", imageHandler.UploadInPostImage)

	// Initialize WebSocket hub
	hub := handler.NewHub(messageUsecase)
	go hub.Run()

	// WebSocket route
	e.GET("/ws/community/:communityId", func(c echo.Context) error {
		handler.ServeWs(hub, messageUsecase, c.Response(), c.Request())
		return nil // ServeWs handles the response, so return nil
	})

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
