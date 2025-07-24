package test

import (
	"backend/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetAllBlogs(t *testing.T) {
	e := echo.New()
	h := &handler.BlogHandler{} // Serviceはモック化推奨

	req := httptest.NewRequest(http.MethodGet, "/api/blog", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// 実行
	err := h.GetAllBlogs(c)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
} 