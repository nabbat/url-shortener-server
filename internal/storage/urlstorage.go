package storage

type URLStorage struct {
	urlMap map[string]string
}

func NewURLStorage() *URLStorage {
	storage := &URLStorage{
		urlMap: make(map[string]string),
	}
	// Добавляем тестовое соответствие
	storage.AddURL("aHR0cH", "https://practicum.yandex.ru/")
	return storage
}

// AddURL добавляет пару сокращенный URL -> оригинальный URL
func (storage *URLStorage) AddURL(shortURL, originalURL string) {
	storage.urlMap[shortURL] = originalURL
}

// GetOriginalURL возвращает оригинальный URL по сокращенному URL
func (storage *URLStorage) GetOriginalURL(shortURL string) string {
	return storage.urlMap[shortURL]
}
