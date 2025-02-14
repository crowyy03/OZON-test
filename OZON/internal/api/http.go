package api
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/crowyy03/OZON/internal/shortener"
	"github.com/gorilla/mux"
)

type Storage interface {
	Save(shortURL, originalURL string) error
	Get(shortURL string) (string, error)
}

var store Storage

func NewRouter(s Storage) *mux.Router {
	store = s
	r := mux.NewRouter()
	r.HandleFunc("/shorten", ShortenURL).Methods("POST")
	r.HandleFunc("/expand/{shortURL}", ExpandURL).Methods("GET")
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")
	return r
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	log.Println("DEBUG: Вызван ShortenURL")
	var req struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("ERROR: Неверный запрос: %v", err)
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	if store == nil {
		log.Println("ERROR: Хранилище не инициализировано")
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	shortURL := shortener.GenerateShortURL()

	if err := store.Save(shortURL, req.URL); err != nil {
		log.Printf("Ошибка сохранения URL: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}

func ExpandURL(w http.ResponseWriter, r *http.Request) {
	log.Println("DEBUG: Вызван ExpandURL")

	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	// Используем store для получения оригинального URL
	originalURL, err := store.Get(shortURL)
	if err != nil {
		log.Printf("ERROR: Не удалось получить оригинальный URL: %v", err)
		http.Error(w, "Ссылка не найдена", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"original_url": originalURL})
}
