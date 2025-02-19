package shortener

import (
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	testCases := []struct {
		name        string
		expectedLen int
	}{
		{"Default length", LenOfShortURL},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			shortURL := GenerateShortURL()

			// Проверяем, что длина соответствует ожидаемой
			if len(shortURL) != tc.expectedLen {
				t.Errorf("Expected length %d, got %d", tc.expectedLen, len(shortURL))
			}

			// Проверяем, что строка содержит только разрешенные символы
			for _, char := range shortURL {
				if !isValidCharacter(char) {
					t.Errorf("Invalid character in short URL: %c", char)
				}
			}
		})
	}
}

// isValidCharacter проверяет, входит ли символ в разрешенные
func isValidCharacter(char rune) bool {
	for _, validChar := range data {
		if char == validChar {
			return true
		}
	}
	return false
}
