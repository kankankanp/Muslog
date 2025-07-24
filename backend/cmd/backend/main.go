package main

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// .envファイルの明示的な読み込み
	err := godotenv.Load("../../.env")  // ここで相対パスを指定して、ルートディレクトリから読み込みます
	if err != nil {
		fmt.Println("Error loading .env file")
		return
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

	// データベース接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
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

	// Blog APIのルーティング
	e.GET("/api/blog", blogHandler.GetAllBlogs)
	e.POST("/api/blog", blogHandler.CreateBlog)
	e.GET("/api/blog/:id", blogHandler.GetBlogByID)
	e.PUT("/api/blog/:id", blogHandler.UpdateBlog)
	e.DELETE("/api/blog/:id", blogHandler.DeleteBlog)
	e.GET("/api/blog/page/:page", blogHandler.GetBlogsByPage)

	// User APIのルーティング
	e.GET("/api/user", userHandler.GetAllUsers)
	e.GET("/api/user/:id", userHandler.GetUserByID)
	e.GET("/api/user/:id/posts", userHandler.GetUserPosts)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}