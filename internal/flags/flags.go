package flags

import (
	flag "github.com/spf13/pflag"
)

// Flags структура для хранения настроек
type Flags struct {
	RunAddr   string
	ResultURL string
}

// ParseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func ParseFlags() *Flags {
	// Create a Config instance
	flg := &Flags{}
	flag.StringVarP(&flg.RunAddr, "a", "a", "localhost:8080", "Адрес запуска HTTP-сервера.")
	flag.StringVarP(&flg.ResultURL, "b", "b", "localhost:8080", "Адрес результирующего сокращённого URL.")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
	return flg
}
