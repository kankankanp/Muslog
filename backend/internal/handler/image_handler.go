package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kankankanp/Muslog/internal/entity"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

// ImageHandler handles image-related HTTP requests.
type ImageHandler struct {
	imageUsecase usecase.ImageUsecase
	userRepo     UserRepository
	postRepo     PostRepository
}

// NewImageHandler creates a new ImageHandler.
func NewImageHandler(imageUsecase usecase.ImageUsecase, userRepo UserRepository, postRepo PostRepository) *ImageHandler {
	return &ImageHandler{
		imageUsecase: imageUsecase,
		userRepo:     userRepo,
		postRepo:     postRepo,
	}
}

// UploadProfileImage handles uploading a user's profile image.
// POST /api/v1/users/:userId/profile-image
func (h *ImageHandler) UploadProfileImage(c echo.Context) error {
	userID := c.Param("userId")

	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "image file is required")
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to open image file")
	}
	defer src.Close()

	// Get current user to check if they have an existing profile image
	user, err := h.userRepo.FindByID(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	// If an old image exists, delete it from S3
	if user.ProfileImageUrl != "" {
		err = h.imageUsecase.DeleteImage(c.Request().Context(), user.ProfileImageUrl)
		if err != nil {
			// Log the error but don't block the upload of the new image
			c.Logger().Errorf("failed to delete old profile image from S3: %v", err)
		}
	}

	// Upload new image
	imageUrl, err := h.imageUsecase.UploadImage(c.Request().Context(), src, file, "profile-images")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to upload profile image: %v", err))
	}

	// Update user's profile image URL in the database
	user.ProfileImageUrl = imageUrl
	err = h.userRepo.Update(c.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to update user profile image URL: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Profile image uploaded successfully", "imageUrl": imageUrl})
}

// UploadPostHeaderImage handles uploading a post's header image.
// POST /api/v1/posts/:postId/header-image
func (h *ImageHandler) UploadPostHeaderImage(c echo.Context) error {
	postIDStr := c.Param("postId")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid post ID")
	}

	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "image file is required")
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to open image file")
	}
	defer src.Close()

	// Get current post to check if it has an existing header image
	post, err := h.postRepo.FindByID(c.Request().Context(), uint(postID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "post not found")
	}

	// If an old image exists, delete it from S3
	if post.HeaderImageUrl != "" {
		err = h.imageUsecase.DeleteImage(c.Request().Context(), post.HeaderImageUrl)
		if err != nil {
			// Log the error but don't block the upload of the new image
			c.Logger().Errorf("failed to delete old post header image from S3: %v", err)
		}
	}

	// Upload new image
	imageUrl, err := h.imageUsecase.UploadImage(c.Request().Context(), src, file, "post-header-images")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to upload post header image: %v", err))
	}

	// Update post's header image URL in the database
	post.HeaderImageUrl = imageUrl
	err = h.postRepo.Update(c.Request().Context(), post)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to update post header image URL: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Post header image uploaded successfully", "imageUrl": imageUrl})
}

// UploadInPostImage handles uploading a generic image for use within a post (e.g., markdown).
// POST /api/v1/images/upload
func (h *ImageHandler) UploadInPostImage(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "image file is required")
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to open image file")
	}
	defer src.Close()

	// Upload image to a generic "in-post-images" folder
	imageUrl, err := h.imageUsecase.UploadImage(c.Request().Context(), src, file, "in-post-images")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to upload in-post image: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Image uploaded successfully", "imageUrl": imageUrl})
}

// UserRepository and PostRepository interfaces are needed for the handler to interact with the database.
// These should already exist or be defined in their respective repository files.
// For now, I'll define minimal interfaces here to avoid compilation errors.
type UserRepository interface {
	FindByID(ctx context.Context, id string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
}

type PostRepository interface {
	FindByID(ctx context.Context, id uint) (*entity.Post, error)
	Update(ctx context.Context, post *entity.Post) error
}