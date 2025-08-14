package router

import (
	"backend/internal/interfaces/controllers"
	"backend/internal/interfaces/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func InitRouter(e *echo.Echo, postHandler controllers.PostController, userHandler controllers.UserController, tagHandler controllers.TagController, spotifyHandler controllers.SpotifyController, likeHandler controllers.LikeController, oauthHandler controllers.OAuthController) {
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.00.0.1:3000"},
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
	protected.Use(middlewares.AuthMiddleware(middlewares.AuthMiddlewareConfig{
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

	// likes
	postGroup.POST("/:postID/like", likeHandler.LikePost)
	postGroup.DELETE("/:postID/unlike", likeHandler.UnlikePost)

	// users
	userGroup := protected.Group("/users")
	userGroup.GET("", userHandler.GetAllUsers)
	userGroup.GET("/:id", userHandler.GetUserByID)
	//	userGroup.GET("/:id/posts", userHandler.GetUserPosts)

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
}
