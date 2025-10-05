package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockStorageClient struct {
	mock.Mock
}

func (m *MockStorageClient) Upload(ctx context.Context, folder, fileName, contentType string, data []byte) (string, error) {
	args := m.Called(ctx, folder, fileName, contentType, data)
	return args.String(0), args.Error(1)
}

func (m *MockStorageClient) Delete(ctx context.Context, imageURL string) error {
	args := m.Called(ctx, imageURL)
	return args.Error(0)
}
