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
	DeleteImage(ctx context.Context, imageUrl string) error
}

type imageUsecase struct {
	s3Client *s3.Client
	s3Bucket string
	s3Region string
}

func NewImageUsecase(s3Client *s3.Client, s3Bucket string, s3Region string) ImageUsecase {
	return &imageUsecase{
		s3Client: s3Client,
		s3Bucket: s3Bucket,
		s3Region: s3Region,
	}
}

func (u *imageUsecase) UploadImage(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, folder string) (string, error) {
	ext := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	key := fmt.Sprintf("%s/%s", folder, fileName)

	_, err := u.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(u.s3Bucket),
		Key:         aws.String(key),
		Body:        file,
		ACL:         "public-read", // Make the object publicly readable
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	// Construct the public URL
	imageUrl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.s3Bucket, u.s3Region, key)
	return imageUrl, nil
}

// DeleteImage deletes a file from S3 given its URL.
func (u *imageUsecase) DeleteImage(ctx context.Context, imageUrl string) error {
	// Extract the key from the URL
	// Assuming URL format: https://<bucket-name>.s3.<region>.amazonaws.com/<key>
	// This parsing might need to be more robust depending on actual URL variations.
	key := ""
	prefix := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", u.s3Bucket, u.s3Region)
	if len(imageUrl) > len(prefix) && imageUrl[:len(prefix)] == prefix {
		key = imageUrl[len(prefix):]
	} else {
		return fmt.Errorf("invalid image URL format for deletion: %s", imageUrl)
	}

	_, err := u.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(u.s3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %w", err)
	}
	return nil
}
