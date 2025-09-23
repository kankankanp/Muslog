package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// supabaseClient はSupabase Storageを利用したストレージクライアントです。
type supabaseClient struct {
	httpClient      *http.Client
	baseURL         string
	bucket          string
	apiKey          string
	publicURLPrefix string
}

// NewSupabaseClient はSupabase向けのストレージクライアントを生成します。

func NewSupabaseClient(baseURL, bucket, apiKey string, httpClient *http.Client) Client {
	normalizedBase := normalizeSupabaseBaseURL(baseURL)
	trimmedBucket := strings.Trim(strings.TrimSpace(bucket), "/")
	trimmedKey := strings.TrimSpace(apiKey)
	client := httpClient
	if client == nil {
		client = &http.Client{Timeout: 30 * time.Second}
	}
	return &supabaseClient{
		httpClient:      client,
		baseURL:         normalizedBase,
		bucket:          trimmedBucket,
		apiKey:          trimmedKey,
		publicURLPrefix: fmt.Sprintf("%s/storage/v1/object/public/%s/", normalizedBase, trimmedBucket),
	}
}

func (c *supabaseClient) Upload(ctx context.Context, folder, fileName, contentType string, data []byte) (string, error) {
	key := buildObjectKey(folder, fileName)
	uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", c.baseURL, c.bucket, encodeKeyForURL(key))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uploadURL, bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("failed to create Supabase upload request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("apikey", c.apiKey)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	// 上書きアップロードを許容しておく
	req.Header.Set("x-upsert", "true")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to Supabase: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
		return "", fmt.Errorf("supabase upload failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	return c.publicURLPrefix + encodeKeyForURL(key), nil
}

func (c *supabaseClient) Delete(ctx context.Context, imageURL string) error {
	if !strings.HasPrefix(imageURL, c.publicURLPrefix) {
		return fmt.Errorf("invalid Supabase public URL: %s", imageURL)
	}
	encodedKey := strings.TrimPrefix(imageURL, c.publicURLPrefix)
	decodedKey, err := url.PathUnescape(encodedKey)
	if err != nil {
		return fmt.Errorf("failed to decode Supabase object key: %w", err)
	}

	deleteURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", c.baseURL, c.bucket, encodeKeyForURL(decodedKey))
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, deleteURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create Supabase delete request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("apikey", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete file from Supabase: %w", err)
	}
	defer resp.Body.Close()

	// 404を許容して冪等性を確保
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusNotFound {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
		return fmt.Errorf("supabase delete failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	return nil
}

func encodeKeyForURL(key string) string {
	segments := strings.Split(key, "/")
	encoded := make([]string, 0, len(segments))
	for _, segment := range segments {
		if segment == "" {
			continue
		}
		encoded = append(encoded, url.PathEscape(segment))
	}
	return strings.Join(encoded, "/")
}

func normalizeSupabaseBaseURL(raw string) string {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return ""
	}
	parsed, err := url.Parse(trimmed)
	if err != nil || parsed.Host == "" {
		return strings.TrimRight(trimmed, "/")
	}
	scheme := parsed.Scheme
	if scheme == "" {
		scheme = "https"
	}
	return strings.TrimRight(fmt.Sprintf("%s://%s", scheme, parsed.Host), "/")
}
