package storage

import "context"

// Client は画像などのバイナリアセットを管理する抽象インターフェースです。
type Client interface {
	Upload(ctx context.Context, folder, fileName, contentType string, data []byte) (string, error)
	Delete(ctx context.Context, imageURL string) error
}
