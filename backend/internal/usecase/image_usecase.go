package usecase

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/kankankanp/Muslog/internal/infrastructure/storage"
)

type ImageUsecase interface {
	UploadImage(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, folder string) (string, error)
	DeleteImage(ctx context.Context, imageURL string) error
}

type imageUsecaseImpl struct {
	storage storage.Client
}

func NewImageUsecase(storageClient storage.Client) ImageUsecase {
	return &imageUsecaseImpl{storage: storageClient}
}

func (u *imageUsecaseImpl) UploadImage(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, folder string) (string, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	ext := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	imageURL, err := u.storage.Upload(ctx, folder, fileName, contentType, data)
	if err != nil {
		return "", fmt.Errorf("failed to store image: %w", err)
	}

	return imageURL, nil
}

func (u *imageUsecaseImpl) DeleteImage(ctx context.Context, imageURL string) error {
	if imageURL == "" {
		return nil
	}

	if err := u.storage.Delete(ctx, imageURL); err != nil {
		return fmt.Errorf("failed to delete image: %w", err)
	}
	return nil
}
