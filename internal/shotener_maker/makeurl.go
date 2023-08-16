package shotener_maker

import "encoding/base64"

// GenerateID Функция для генерации уникального идентификатора
func GenerateID(fullURL []byte) string {
	encodedStr := base64.URLEncoding.EncodeToString(fullURL)
	// Возвращаем первые 6 символов закодированной строки
	if len(encodedStr) > 6 {
		return encodedStr[:6]
	}
	return encodedStr
}
