package storage
import (
	"database/sql"
	"sync"
)

type InMemoryStorage struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		data: make(map[string]string),
	}
}

func (s *InMemoryStorage) Save(shortURL, originalURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[shortURL] = originalURL
	return nil
}
func (s *InMemoryStorage) Get(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	originalURL, exists := s.data[shortURL]
	if !exists {
		return "", sql.ErrNoRows
	}
	return originalURL, nil
}
