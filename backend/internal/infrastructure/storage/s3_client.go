package storage

import (
	"bytes"
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// s3Client はAWS S3を利用したストレージクライアントです。
type s3Client struct {
	client    *s3.Client
	bucket    string
	region    string
	urlPrefix string
}

// NewS3Client はS3向けのストレージクライアントを生成します。
func NewS3Client(client *s3.Client, bucket, region string) Client {
	trimmedBucket := strings.TrimSpace(bucket)
	trimmedRegion := strings.TrimSpace(region)
	return &s3Client{
		client:    client,
		bucket:    trimmedBucket,
		region:    trimmedRegion,
		urlPrefix: fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", trimmedBucket, trimmedRegion),
	}
}

func (c *s3Client) Upload(ctx context.Context, folder, fileName, contentType string, data []byte) (string, error) {
	key := buildObjectKey(folder, fileName)
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(c.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}
	return c.urlPrefix + key, nil
}

func (c *s3Client) Delete(ctx context.Context, imageURL string) error {
	if !strings.HasPrefix(imageURL, c.urlPrefix) {
		return fmt.Errorf("invalid S3 URL: %s", imageURL)
	}
	key := strings.TrimPrefix(imageURL, c.urlPrefix)
	_, err := c.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %w", err)
	}
	return nil
}

func buildObjectKey(folder, fileName string) string {
	cleanFolder := strings.Trim(path.Clean("/"+folder), "/")
	if cleanFolder == "." || cleanFolder == "" {
		return fileName
	}
	return cleanFolder + "/" + fileName
}
