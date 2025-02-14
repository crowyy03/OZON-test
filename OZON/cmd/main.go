package main
import (
	"log"
	"net/http"

	"github.com/crowyy03/OZON/internal/config"

	"github.com/crowyy03/OZON/internal/api"
	"github.com/crowyy03/OZON/internal/storage"
)

func main() {
	cfg, err := config.LoadConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}
	var store api.Storage

	if cfg.Storage.Type == "postgres" {
		dbStore, err := storage.NewPostgresStorage(cfg.Postgres.DSN)
		if err != nil {
			log.Fatal("Ошибка подключения к PostgreSQL:", err)
		}
		store = dbStore
		log.Println("Используется PostgreSQL")

	} else if cfg.Storage.Type == "in-memory" {
		store = storage.NewInMemoryStorage()
		log.Println("Используется in-memory хранилище")
	}

	router := api.NewRouter(store)
	log.Printf("Наш сервер запущен на http://0.0.0.0:%s\n", cfg.Server.Port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+cfg.Server.Port, router))

}
