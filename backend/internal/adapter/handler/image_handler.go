package handler

import (
	"fmt"
	"net/http"
	"strconv"

	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"

	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ImageHandler struct {
	imageUsecase usecase.ImageUsecase
	userRepo     domainRepo.UserRepository
	postRepo     domainRepo.PostRepository
}

func NewImageHandler(imageUsecase usecase.ImageUsecase, userRepo domainRepo.UserRepository, postRepo domainRepo.PostRepository) *ImageHandler {
	return &ImageHandler{
		imageUsecase: imageUsecase,
		userRepo:     userRepo,
		postRepo:     postRepo,
	}
}

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

	user, err := h.userRepo.FindByID(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	if user.ProfileImageUrl != "" {
		err = h.imageUsecase.DeleteImage(c.Request().Context(), user.ProfileImageUrl)
		if err != nil {
			c.Logger().Errorf("failed to delete old profile image from S3: %v", err)
		}
	}

	imageUrl, err := h.imageUsecase.UploadImage(c.Request().Context(), src, file, "profile-images")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to upload profile image: %v", err))
	}

	user.ProfileImageUrl = imageUrl
	err = h.userRepo.Update(c.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to update user profile image URL: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Profile image uploaded successfully", "imageUrl": imageUrl})
}

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

	post, err := h.postRepo.FindByID(c.Request().Context(), uint(postID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "post not found")
	}

	if post.HeaderImageUrl != "" {
		err = h.imageUsecase.DeleteImage(c.Request().Context(), post.HeaderImageUrl)
		if err != nil {
			c.Logger().Errorf("failed to delete old post header image from S3: %v", err)
		}
	}

	imageUrl, err := h.imageUsecase.UploadImage(c.Request().Context(), src, file, "post-header-images")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to upload post header image: %v", err))
	}

	post.HeaderImageUrl = imageUrl
	err = h.postRepo.Update(c.Request().Context(), post)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to update post header image URL: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Post header image uploaded successfully", "imageUrl": imageUrl})
}

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

	imageUrl, err := h.imageUsecase.UploadImage(c.Request().Context(), src, file, "in-post-images")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to upload in-post image: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Image uploaded successfully", "imageUrl": imageUrl})
}
