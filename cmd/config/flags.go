package config

import (
	flag "github.com/spf13/pflag"
)

// Config структура для хранения настроек
type Config struct {
	RunAddr   string
	ResultURL string
}

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func ParseFlags(cfg *Config) {
	flag.StringVar(&cfg.RunAddr, "a", "localhost:8080", "Адрес запуска HTTP-сервера. По умолчанию localhost:8080")
	flag.StringVar(&cfg.ResultURL, "b", "http://localhost:8080/", "Адрес результирующего сокращённого URL. По умолчанию http://localhost:8080/{короткая ссылка}")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
