package main

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/seeder"
	"backend/internal/service"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
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

	if err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Track{}); err != nil {
		log.Fatalf("マイグレーション失敗: %v", err)
	}

	if err := seeder.Seed(db); err != nil {
		log.Fatalf("シード注入失敗: %v", err)
	}

	blogRepo := &repository.BlogRepository{DB: db}
	blogService := &service.BlogService{Repo: blogRepo}
	blogHandler := &handler.BlogHandler{Service: blogService}

	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userHandler := &handler.UserHandler{Service: userService}

	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	e.GET("/swagger/*", swagger.WrapHandler)

	// Auth routes
	authGroup := e.Group("/auth")
	authGroup.POST("/login", userHandler.Login)
	authGroup.POST("/refresh", userHandler.RefreshToken)
	authGroup.GET("/me", userHandler.GetMe, middleware.AuthMiddleware)

	// Blog routes
	blogGroup := e.Group("/blogs")
	blogGroup.GET("", blogHandler.GetAllBlogs)
	blogGroup.GET("/:id", blogHandler.GetBlogByID)
	blogGroup.GET("/page/:page", blogHandler.GetBlogsByPage)

	// Protected blog routes
	protectedBlogGroup := e.Group("/blogs")
	protectedBlogGroup.Use(middleware.AuthMiddleware)
	protectedBlogGroup.POST("", blogHandler.CreateBlog)
	protectedBlogGroup.PUT("/:id", blogHandler.UpdateBlog)
	protectedBlogGroup.DELETE("/:id", blogHandler.DeleteBlog)

	// User routes
	userGroup := e.Group("/users")
	userGroup.GET("", userHandler.GetAllUsers)
	userGroup.GET("/:id", userHandler.GetUserByID)
	userGroup.GET("/:id/posts", userHandler.GetUserPosts)

	e.Logger.Fatal(e.Start(":8080"))
}