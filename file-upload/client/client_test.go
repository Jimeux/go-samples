package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func uploadServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formFile, _, _ := r.FormFile("file")
		defer formFile.Close()
		_, _ = io.Copy(new(bytes.Buffer), formFile)
		w.WriteHeader(http.StatusCreated)
	}))
}

func BenchmarkUploadFile(b *testing.B) {
	// server := uploadServer()
	dir, err := os.Getwd()
	handleError(err)
	fileName := path.Join(dir,"..", "hello.txt")

	for i := 0; i < b.N; i++ {
		file, err := os.Open(fileName)
		handleError(err)
		uploadFile(file, "http://localhost:3000", "file", fileName)
		file.Close()
	}
}

func BenchmarkUploadFileAsync(b *testing.B) {
	// server := uploadServer()
	dir, err := os.Getwd()
	handleError(err)
	fileName := path.Join(dir,"..", "hello.txt")

	for i := 0; i < b.N; i++ {
		file, err := os.Open(fileName)
		handleError(err)
		uploadFileAsync(file, "http://localhost:3000", "file", fileName)
		file.Close()
	}
}
