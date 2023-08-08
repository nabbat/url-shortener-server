package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nabbat/url-shortener-server.git/cmd/config"
	"io"
	"log"
	"net/http"
)

// Словарь для хранения соответствий между сокращёнными и оригинальными URL
// TODO Создать хранилище
var urlMap = map[string]string{}

// Перенаправляем по полной ссылке
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid request type", http.StatusBadRequest)
		return
	}
	// Добавляем тестовое соответствие в словарь
	urlMap["aHR0cH"] = "https://practicum.yandex.ru/"
	// Получаем идентификатор из URL-пути
	id := r.URL.Path[1:]

	// Получаем оригинальный URL из словаря

	if originalURL, found := urlMap[id]; found {
		// Устанавливаем заголовок Location и возвращаем ответ с кодом 307
		w.Header().Set("Location", originalURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	http.Error(w, "Ссылка не найдена", http.StatusBadRequest)

}

func shortenURLHandler(w http.ResponseWriter, r *http.Request, c *config.Config) {
	// Читаем тело запроса (URL)
	urlBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
		return
	}

	// Преобразуем в строку
	url := string(urlBytes)

	// Генерируем уникальный идентификатор сокращённого URL
	id := generateID(url)

	// Добавляем соответствие в словарь
	urlMap[id] = url

	// Отправляем ответ с сокращённым URL
	shortenedURL := fmt.Sprintf("http://%s/%s", c.ResultURL, id)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	if _, err := io.WriteString(w, shortenedURL); err != nil {
		log.Fatal(err)
	}
}

// Простая функция для генерации уникального идентификатора
func generateID(fullURL string) string {
	encodedStr := base64.URLEncoding.EncodeToString([]byte(fullURL))
	// Возвращаем первые 6 символов закодированной строки
	if len(encodedStr) > 6 {
		return encodedStr[:6]
	}
	return encodedStr
}

func main() {
	c := config.SetEnv()

	// Run server
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenURLHandler(w, r, c)
	}).Methods("POST")
	r.HandleFunc("/{idShortenURL}", redirectHandler).Methods("GET")
	fmt.Println("RunAddr: ResultURL: ", c.RunAddr, c.ResultURL)
	fmt.Println("Running server on", c.RunAddr)
	err := http.ListenAndServe(c.RunAddr, r)
	if err != nil {
		panic(err)
	}
}
