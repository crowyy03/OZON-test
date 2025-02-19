package api

// Валидный URL сокращается успешно.
// Ошибки при невалидном JSON.
// Ошибки при пустом URL.

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/crowyy03/OZON/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestShortenUrl(t *testing.T) {
	store := storage.NewInMemoryStorage()
	router := NewRouter(store)

	test := []struct {
		name       string
		input      string
		wantStatus int
	}{
		{"Valid URL", `{"url": "https://example.com"}`, http.StatusOK},
		{"Invalid JSON", `{"url": "`, http.StatusBadRequest},
		{"Empty URL", `{"url": ""}`, http.StatusBadRequest},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer([]byte(tt.input)))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.wantStatus, rr.Code)
		})
	}
}
