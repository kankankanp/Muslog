package main

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"fmt"
	"os"
	"time"
	"log"

	_ "backend/docs"
	"backend/internal/model"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	swagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Simple Blog API
// @version 1.0
// @description This is a simple blog API.
// @host localhost:8080
// @BasePath /
func main() {
	// .envの読み込み（ローカル用）
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load(".env")
	}

	// DB接続情報の取得
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// データベース接続文字列の作成
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// データベース接続（リトライ付き）
	var db *gorm.DB
	var err error
	maxRetries := 10

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("✅ DB接続に成功しました")
			break
		}
		fmt.Printf("⏳ DB接続リトライ中... (%d/%d): %v\n", i+1, maxRetries, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		panic("❌ データベース接続失敗: " + err.Error())
	}

	if err := db.AutoMigrate(&model.User{}, &model.Post{}); err != nil {
			log.Fatalf("マイグレーション失敗: %v", err)
	}	
	// DI
	blogRepo := &repository.BlogRepository{DB: db}
	blogService := &service.BlogService{Repo: blogRepo}
	blogHandler := &handler.BlogHandler{Service: blogService}

	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userHandler := &handler.UserHandler{Service: userService}

	// Echoの設定
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Swagger
	e.GET("/swagger/*", swagger.WrapHandler)

	// Blog APIのルーティング
	e.GET("/blogs", blogHandler.GetAllBlogs)
	e.POST("/blogs", blogHandler.CreateBlog)
	e.GET("/blogs/:id", blogHandler.GetBlogByID)
	e.PUT("/blogs/:id", blogHandler.UpdateBlog)
	e.DELETE("/blogs/:id", blogHandler.DeleteBlog)
	e.GET("/blogs/page/:page", blogHandler.GetBlogsByPage)

	// User APIのルーティング
	e.GET("/users", userHandler.GetAllUsers)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.GET("/users/:id/posts", userHandler.GetUserPosts)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}