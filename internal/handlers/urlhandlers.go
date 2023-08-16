package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nabbat/url-shortener-server.git/cmd/config"
	"github.com/nabbat/url-shortener-server.git/internal/shotener_maker"
	urlstorage "github.com/nabbat/url-shortener-server.git/internal/storage"
	"io"
	"log"
	"net/http"
)

type RedirectHandlerInterface interface {
	HandleRedirect(storage *urlstorage.URLStorage) http.HandlerFunc
}

type ShortenURLHandlerInterface interface {
	HandleShortenURL(storage *urlstorage.URLStorage, c *config.Config) http.HandlerFunc
}

type RedirectHandler struct{}

func (rh *RedirectHandler) HandleRedirect(storage *urlstorage.URLStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "invalid request type", http.StatusBadRequest)
			return
		}

		// Получаем идентификатор из URL-пути
		vars := mux.Vars(r)
		shortURL := vars["idShortenURL"]

		// Получаем оригинальный URL
		originalURL := storage.GetOriginalURL(shortURL)

		if originalURL == "" {
			http.Error(w, "Ссылка не найдена", http.StatusBadRequest)
			return
		}
		// Устанавливаем заголовок Location и возвращаем ответ с кодом 307
		w.Header().Set("Location", originalURL)
		w.WriteHeader(http.StatusTemporaryRedirect)

	}
}

type ShortenURLHandler struct{}

func (sh *ShortenURLHandler) HandleShortenURL(storage *urlstorage.URLStorage, c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Читаем тело запроса (URL)
		defer r.Body.Close()
		urlBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
			return
		}

		// Генерируем уникальный идентификатор сокращённого URL
		shortURL := shotener_maker.GenerateID(urlBytes)

		// Добавляем соответствие в словарь
		storage.AddURL(shortURL, string(urlBytes))

		// Отправляем ответ с сокращённым URL
		shortenedURL := fmt.Sprintf("%s/%s", c.ResultURL, shortURL)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		if _, err := io.WriteString(w, shortenedURL); err != nil {
			log.Fatal(err)
		}
	}
}

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
