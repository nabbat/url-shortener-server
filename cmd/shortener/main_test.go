package main

import (
	"net/http"
	"testing"
)

func Test_generateID(t *testing.T) {
	type args struct {
		fullURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateID(tt.args.fullURL); got != tt.want {
				t.Errorf("generateID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redirectHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redirectHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_shortenURLHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortenURLHandler(tt.args.w, tt.args.r)
		})
	}
}
