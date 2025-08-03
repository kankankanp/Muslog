package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-blog/backend/internal/handler"
	"simple-blog/backend/internal/middleware"
	"simple-blog/backend/internal/model"
	"simple-blog/backend/internal/repository"
	"simple-blog/backend/internal/seeder"
	"simple-blog/backend/internal/service"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.SetOutput(os.Stdout)
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
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("DB接続に成功しました")
			break
		}
		fmt.Printf("DB接続リトライ中... (%d/10): %v\n", i+1, err)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		panic("データベース接続失敗: " + err.Error())
	}

	if err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Track{}, &model.Tag{}, &model.PostTag{}, &model.Like{}); err != nil {
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

	oauthService := service.NewOAuthService(userRepo)
	oauthHandler := handler.NewOAuthHandler(oauthService)

	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://192.168.3.10:3000"},
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
	postGroup.GET("", postHandler.GetAllPosts)
	postGroup.GET("/:id", postHandler.GetPostByID)
	postGroup.GET("/page/:page", postHandler.GetPostsByPage)
	postGroup.POST("", postHandler.CreatePost)
	postGroup.PUT("/:id", postHandler.UpdatePost)
	postGroup.DELETE("/:id", postHandler.DeletePost)

	// likes
	postGroup.POST("/:postID/like", likeHandler.LikePost)
	postGroup.DELETE("/:postID/unlike", likeHandler.UnlikePost)

	// users
	userGroup := protected.Group("/users")
	userGroup.GET("", userHandler.GetAllUsers)
	userGroup.GET("/:id", userHandler.GetUserByID)
	userGroup.GET("/:id/posts", userHandler.GetUserPosts)

	// spotify endpoint moved to public group above

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

	e.Logger.Fatal(e.Start(":8080"))
}