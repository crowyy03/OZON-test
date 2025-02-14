package shortener
import (
	"math/rand"
	"time"
)

const data = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456890_"

const LenOfShortURL = 10

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateShortURL() string {
	shortURL := make([]byte, LenOfShortURL)
	for i := range shortURL {
		shortURL[i] = data[rand.Intn(len(data))]
	}
	return string(shortURL)
}
