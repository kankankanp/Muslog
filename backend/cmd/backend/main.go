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
	infraTx "github.com/kankankanp/Muslog/internal/infrastructure/db"
	dblogger "github.com/kankankanp/Muslog/internal/infrastructure/logger"
	"github.com/kankankanp/Muslog/internal/infrastructure/model"
	"github.com/kankankanp/Muslog/internal/infrastructure/repository"
	"github.com/kankankanp/Muslog/internal/middleware"
	"github.com/kankankanp/Muslog/internal/seeder"
	"github.com/kankankanp/Muslog/internal/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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
	dsn := os.Getenv("DATABASE_URL")

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
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: filteredLogger,
			NamingStrategy: schema.NamingStrategy{
				NameReplacer: strings.NewReplacer("Model", ""),
			},
		})
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

	// Ensure join table mapping for many2many(Post <-> Tag) uses post_tags (post_id, tag_id)
	if err := db.SetupJoinTable(&model.PostModel{}, "Tags", &model.PostTagModel{}); err != nil {
		log.Fatalf("failed to setup join table for Post.Tags: %v", err)
	}

	// Ensure required PostgreSQL extensions
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatalf("failed to create uuid-ossp extension: %v", err)
	}

	// Align column types for foreign keys to UUID where necessary
	// These adjustments run before AutoMigrate to avoid FK type mismatch.

	// communities.id -> uuid (if previously created as text)
	if db.Migrator().HasTable(&model.CommunityModel{}) {
		if err := db.Exec(`ALTER TABLE "communities" ALTER COLUMN "id" TYPE uuid USING "id"::uuid;`).Error; err != nil {
			// Fallback: replace column safely
			tx := db.Begin()
			_ = tx.Exec(`ALTER TABLE "communities" ADD COLUMN IF NOT EXISTS "id_tmp" uuid;`).Error
			_ = tx.Exec(`UPDATE "communities" SET "id_tmp" = CASE WHEN "id" ~* '^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$' THEN "id"::uuid ELSE uuid_generate_v4() END;`).Error
			_ = tx.Exec(`ALTER TABLE "communities" DROP CONSTRAINT IF EXISTS "communities_pkey";`).Error
			_ = tx.Exec(`ALTER TABLE "communities" DROP COLUMN IF EXISTS "id";`).Error
			_ = tx.Exec(`ALTER TABLE "communities" RENAME COLUMN "id_tmp" TO "id";`).Error
			_ = tx.Exec(`ALTER TABLE "communities" ALTER COLUMN "id" SET DEFAULT uuid_generate_v4();`).Error
			_ = tx.Exec(`ALTER TABLE "communities" ADD PRIMARY KEY ("id");`).Error
			_ = tx.Commit().Error
		}
	}

	// messages.community_id -> uuid
	if db.Migrator().HasTable(&model.MessageModel{}) && db.Migrator().HasColumn(&model.MessageModel{}, "community_id") {
		_ = db.Exec(`ALTER TABLE "messages" ALTER COLUMN "community_id" TYPE uuid USING "community_id"::uuid;`).Error
	}
	// messages.sender_id -> uuid
	if db.Migrator().HasTable(&model.MessageModel{}) && db.Migrator().HasColumn(&model.MessageModel{}, "sender_id") {
		_ = db.Exec(`ALTER TABLE "messages" ALTER COLUMN "sender_id" TYPE uuid USING "sender_id"::uuid;`).Error
	}
	// likes.user_id -> uuid
	if db.Migrator().HasTable(&model.LikeModel{}) && db.Migrator().HasColumn(&model.LikeModel{}, "user_id") {
		_ = db.Exec(`ALTER TABLE "likes" ALTER COLUMN "user_id" TYPE uuid USING "user_id"::uuid;`).Error
	}
	// communities.creator_id -> uuid
	if db.Migrator().HasTable(&model.CommunityModel{}) && db.Migrator().HasColumn(&model.CommunityModel{}, "creator_id") {
		_ = db.Exec(`ALTER TABLE "communities" ALTER COLUMN "creator_id" TYPE uuid USING "creator_id"::uuid;`).Error
	}
	// posts.user_id -> uuid
	if db.Migrator().HasTable(&model.PostModel{}) && db.Migrator().HasColumn(&model.PostModel{}, "user_id") {
		_ = db.Exec(`ALTER TABLE "posts" ALTER COLUMN "user_id" TYPE uuid USING "user_id"::uuid;`).Error
	}
	if err := db.AutoMigrate(
		&model.UserModel{},
		&model.PostModel{},
		&model.TrackModel{},
		&model.TagModel{},
		&model.PostTagModel{},
		&model.CommunityModel{},
		&model.MessageModel{},
		&model.LikeModel{},
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
	txManager := infraTx.NewGormTxManager(db)

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, postRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	tagRepo := repository.NewTagRepository(db)
	tagUsecase := usecase.NewTagUsecase(tagRepo, postRepo)
	tagHandler := handler.NewTagHandler(tagUsecase)

	postUsecase := usecase.NewPostUsecase(postRepo, txManager)
	postHandler := handler.NewPostHandler(postUsecase, tagUsecase)

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
		AllowOrigins:     []string{"https://muslog.vercel.app"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Muslog backend is running 🚀. Use /api/v1/* endpoints.")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

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
