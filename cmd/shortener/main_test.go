package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_generateID(t *testing.T) {
	tests := []struct {
		name    string
		fullURL string
		want    string
	}{
		{name: "Проверим генератор на пустую строку", fullURL: "", want: ""},
		{name: "Проверим генератор на НЕпустую строку", fullURL: "https://practicum.yandex.ru/", want: "aHR0cH"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateID(tt.fullURL); got != tt.want {
				t.Errorf("generateID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortenURLHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name                string
		bodyURL             string
		responseContentType string
		targetURL           string
		want                want
	}{
		{name: "positive POST shortenURLHandler",
			bodyURL:             "https://practicum.yandex.ru/",
			responseContentType: "text/plain",
			targetURL:           "/",
			want: want{
				code:        201,
				response:    `http://localhost:8080/aHR0cH`,
				contentType: "text/plain",
			},
		},
		{name: "negative POST shortenURLHandler",
			bodyURL:             "https://practicum.yandex.ru/",
			responseContentType: "application/json",
			targetURL:           "/",
			want: want{
				code:        400,
				response:    "invalid request type\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bodyURLTesting := strings.NewReader(test.bodyURL)
			request := httptest.NewRequest(http.MethodPost, test.targetURL, bodyURLTesting)
			request.Header.Add("Content-Type", test.responseContentType)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			shortenURLHandler(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, res.StatusCode, test.want.code)
			// получаем и проверяем тело запроса
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(res.Body)
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Equal(t, string(resBody), test.want.response)
			assert.Equal(t, res.Header.Get("Content-Type"), test.want.contentType)
		})
	}
}

func Test_redirectHandler(t *testing.T) {
	type want struct {
		code     int
		response string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Проверяем redirect",
			want: want{
				code:     307,
				response: `https://practicum.yandex.ru/`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/aHR0cH", nil)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			redirectHandler(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, res.StatusCode, test.want.code)
			// получаем и проверяем тело запроса
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic(err)
				}
			}(res.Body)
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err, string(resBody))
		})
	}
}
