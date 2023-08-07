package config

import (
	flag "github.com/spf13/pflag"
)

// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
var FlagRunAddr string
var FlagResultURL string

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func ParseFlags() {
	flag.StringVar(&FlagRunAddr, "a", ":8080", "Адрес запуска HTTP-сервера. По умолчанию localhost:8080")
	flag.StringVar(&FlagResultURL, "b", ":8081", "Адрес результирующего сокращённого URL. По умолчанию http://localhost:8081/{короткая ссылка}")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
