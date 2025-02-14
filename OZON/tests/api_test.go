package tests
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/crowyy03/OZON/internal/api"
	"github.com/crowyy03/OZON/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestShortenAndExpand(t *testing.T) {
	store := storage.NewInMemoryStorage()
	router := api.NewRouter(store)

	requestBody := map[string]string{"url": "https://example.com"}
	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	shortURL, exists := response["short_url"]
	assert.True(t, exists)

	req, _ = http.NewRequest("GET", "/expand/"+shortURL, nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "https://example.com", response["original_url"])
}
