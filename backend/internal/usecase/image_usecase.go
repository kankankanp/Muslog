package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type ImageUsecase interface {
	UploadImage(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, folder string) (string, error)
	DeleteImage(ctx context.Context, imageURL string) error
}

type imageUsecaseImpl struct {
	s3Client *s3.Client
	s3Bucket string
	s3Region string
}

func NewImageUsecase(s3Client *s3.Client, s3Bucket string, s3Region string) ImageUsecase {
	return &imageUsecaseImpl{
		s3Client: s3Client,
		s3Bucket: s3Bucket,
		s3Region: s3Region,
	}
}

func (u *imageUsecaseImpl) UploadImage(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, folder string) (string, error) {
	ext := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	key := fmt.Sprintf("%s/%s", folder, fileName)

	_, err := u.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(u.s3Bucket),
		Key:         aws.String(key),
		Body:        file,
		ACL:         "public-read", // 公開URLでアクセス可能にする
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	imageURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.s3Bucket, u.s3Region, key)
	return imageURL, nil
}

func (u *imageUsecaseImpl) DeleteImage(ctx context.Context, imageURL string) error {
	prefix := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", u.s3Bucket, u.s3Region)

	if len(imageURL) <= len(prefix) || imageURL[:len(prefix)] != prefix {
		return fmt.Errorf("invalid image URL format for deletion: %s", imageURL)
	}

	key := imageURL[len(prefix):]

	_, err := u.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(u.s3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %w", err)
	}
	return nil
}
