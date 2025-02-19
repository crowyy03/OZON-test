package storage

// Корректное сохранение и получение ссылок.
// Ошибку при запросе несуществующей ссылки.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryStorage(t *testing.T) {
	store := NewInMemoryStorage()

	err := store.Save("abcde12345", "http://for_example.com")
	assert.NoError(t, err)

	url, err := store.Get("abcde12345")
	assert.NoError(t, err)
	assert.Equal(t, "http://for_example.com", url)

	_, err = store.Get("unreal_url")
	assert.Error(t, err)
}
