package storage
import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(dsn string) (*PostgresStorage, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Save(shortURL, originalURL string) error {
	log.Printf("Сохраняем: short=%s, original=%s", shortURL, originalURL)

	_, err := s.db.Exec("INSERT INTO urls (original_url, short_url) VALUES ($1, $2)", originalURL, shortURL)

	if err != nil {
		log.Printf("Ошибка сохранения: %v", err)
	}
	return err
}

func (s *PostgresStorage) Get(shortURL string) (string, error) {
	var originalURL string
	log.Printf("Ищем URL по короткой ссылке: %s", shortURL)

	err := s.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)

	if err != nil {
		log.Printf("Ошибка поиска: %v", err)
		return "", err
	}

	log.Printf("Найден URL: %s", originalURL)
	return originalURL, nil
}
