package handler

import (
	"net/http"
	"strconv"

	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
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

// =======================
// Profile Image Upload
// =======================
func (h *ImageHandler) UploadProfileImage(c echo.Context) error {
	userID := c.Param("userId")

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "image file is required",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "failed to open image file",
		})
	}
	defer src.Close()

	user, err := h.userRepo.FindByID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.CommonResponse{
			Message: "user not found",
		})
	}

	if user.ProfileImageUrl != "" {
		if err := h.imageUsecase.DeleteImage(c.Request().Context(), user.ProfileImageUrl); err != nil {
			c.Logger().Errorf("failed to delete old profile image from S3: %v", err)
		}
	}

	imageUrl, err := h.imageUsecase.UploadImage(c.Request().Context(), src, file, "profile-images")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "failed to upload profile image",
			Error:   err.Error(),
		})
	}

	user.ProfileImageUrl = imageUrl
	if err := h.userRepo.Update(c.Request().Context(), user); err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "failed to update user profile image URL",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ImageUploadResponse{
		Message:  "Profile image uploaded successfully",
		ImageURL: imageUrl,
	})
}

// =======================
// Post Header Image Upload
// =======================
func (h *ImageHandler) UploadPostHeaderImage(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("postId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "invalid post ID",
		})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "image file is required",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "failed to open image file",
		})
	}
	defer src.Close()

	post, err := h.postRepo.FindByID(c.Request().Context(), uint(postID))
	if err != nil {
		return c.JSON(http.StatusNotFound, response.CommonResponse{
			Message: "post not found",
		})
	}

	if post.HeaderImageUrl != "" {
		if err := h.imageUsecase.DeleteImage(c.Request().Context(), post.HeaderImageUrl); err != nil {
			c.Logger().Errorf("failed to delete old post header image from S3: %v", err)
		}
	}

	imageUrl, err := h.imageUsecase.UploadImage(c.Request().Context(), src, file, "post-header-images")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "failed to upload post header image",
			Error:   err.Error(),
		})
	}

	post.HeaderImageUrl = imageUrl
	if err := h.postRepo.Update(c.Request().Context(), post); err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "failed to update post header image URL",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ImageUploadResponse{
		Message:  "Post header image uploaded successfully",
		ImageURL: imageUrl,
	})
}

// =======================
// In-Post Image Upload
// =======================
func (h *ImageHandler) UploadInPostImage(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "image file is required",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "failed to open image file",
		})
	}
	defer src.Close()

	imageUrl, err := h.imageUsecase.UploadImage(c.Request().Context(), src, file, "in-post-images")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "failed to upload in-post image",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ImageUploadResponse{
		Message:  "Image uploaded successfully",
		ImageURL: imageUrl,
	})
}
