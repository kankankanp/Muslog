package usecase

import (
	"bytes"
	"context"
	"errors"
	"mime/multipart"
	"net/textproto"
	"testing"

	testmock "github.com/kankankanp/Muslog/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockMultipartFile struct {
	*bytes.Reader
}

func newMockMultipartFile(data []byte) *mockMultipartFile {
	return &mockMultipartFile{Reader: bytes.NewReader(data)}
}

func (m *mockMultipartFile) Close() error {
	return nil
}

func (m *mockMultipartFile) Read(p []byte) (int, error) {
	return m.Reader.Read(p)
}

func (m *mockMultipartFile) ReadAt(p []byte, off int64) (int, error) {
	return m.Reader.ReadAt(p, off)
}

func (m *mockMultipartFile) Seek(offset int64, whence int) (int64, error) {
	return m.Reader.Seek(offset, whence)
}

type errorMultipartFile struct{}

func (e *errorMultipartFile) Read(p []byte) (int, error) {
	return 0, errors.New("read error")
}

func (e *errorMultipartFile) Close() error {
	return nil
}

func (e *errorMultipartFile) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func (e *errorMultipartFile) ReadAt(p []byte, off int64) (int, error) {
	return 0, errors.New("read error")
}

func TestImageUsecase_UploadImage(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		file        multipart.File
		header      *multipart.FileHeader
		setup       func(*testmock.MockStorageClient)
		expectedURL string
		expectedErr error
	}{
		{
			name:   "正常系: アップロード成功",
			file:   newMockMultipartFile([]byte("image-data")),
			header: createHeader("image.png", "image/png"),
			setup: func(storage *testmock.MockStorageClient) {
				storage.On("Upload", ctx, "posts", mock.AnythingOfType("string"), "image/png", []byte("image-data")).Return("http://example.com/image.png", nil).Once()
			},
			expectedURL: "http://example.com/image.png",
		},
		{
			name:   "正常系: Content-Typeが無い場合はデフォルトを使用",
			file:   newMockMultipartFile([]byte("image-data")),
			header: createHeader("image.jpg", ""),
			setup: func(storage *testmock.MockStorageClient) {
				storage.On("Upload", ctx, "posts", mock.AnythingOfType("string"), "application/octet-stream", []byte("image-data")).Return("http://example.com/image.jpg", nil).Once()
			},
			expectedURL: "http://example.com/image.jpg",
		},
		{
			name:        "異常系: ファイル読み込みエラー",
			file:        &errorMultipartFile{},
			header:      createHeader("broken.png", "image/png"),
			setup:       func(storage *testmock.MockStorageClient) {},
			expectedErr: errors.New("failed to read file: read error"),
		},
		{
			name:   "異常系: アップロード失敗",
			file:   newMockMultipartFile([]byte("image-data")),
			header: createHeader("image.png", "image/png"),
			setup: func(storage *testmock.MockStorageClient) {
				storage.On("Upload", ctx, "posts", mock.AnythingOfType("string"), "image/png", []byte("image-data")).Return("", errors.New("upload error")).Once()
			},
			expectedErr: errors.New("failed to store image: upload error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := new(testmock.MockStorageClient)
			if tt.setup != nil {
				tt.setup(storage)
			}

			usecase := NewImageUsecase(storage)
			url, err := usecase.UploadImage(ctx, tt.file, tt.header, "posts")

			if tt.expectedErr != nil {
				assert.Empty(t, url)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedURL, url)
			}

			storage.AssertExpectations(t)
		})
	}
}

func TestImageUsecase_DeleteImage(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		imageURL    string
		setup       func(*testmock.MockStorageClient)
		expectedErr error
	}{
		{
			name:     "正常系: URLが空の場合は何もしない",
			imageURL: "",
			setup:    func(storage *testmock.MockStorageClient) {},
		},
		{
			name:     "正常系: 削除成功",
			imageURL: "http://example.com/image.png",
			setup: func(storage *testmock.MockStorageClient) {
				storage.On("Delete", ctx, "http://example.com/image.png").Return(nil).Once()
			},
		},
		{
			name:     "異常系: 削除失敗",
			imageURL: "http://example.com/image.png",
			setup: func(storage *testmock.MockStorageClient) {
				storage.On("Delete", ctx, "http://example.com/image.png").Return(errors.New("delete error")).Once()
			},
			expectedErr: errors.New("failed to delete image: delete error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := new(testmock.MockStorageClient)
			if tt.setup != nil {
				tt.setup(storage)
			}

			usecase := NewImageUsecase(storage)
			err := usecase.DeleteImage(ctx, tt.imageURL)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			if tt.imageURL == "" {
				storage.AssertNotCalled(t, "Delete", mock.Anything, mock.Anything)
			} else {
				storage.AssertExpectations(t)
			}
		})
	}
}

func createHeader(filename, contentType string) *multipart.FileHeader {
	header := textproto.MIMEHeader{}
	if contentType != "" {
		header.Set("Content-Type", contentType)
	}
	return &multipart.FileHeader{
		Filename: filename,
		Header:   header,
	}
}
