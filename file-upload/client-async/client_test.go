package main

import (
	"os"
	"path"
	"testing"
)

func helloFileName() string {
	dir, err := os.Getwd()
	handleError(err)
	return path.Join(dir, "..", fileName)
}

func BenchmarkUploadFileAsync(b *testing.B) {
	fileName := helloFileName()

	for i := 0; i < b.N; i++ {
		file, err := os.Open(fileName)
		handleError(err)
		uploadFileAsync(file, uploadURL, fieldName, fileName)
		err = file.Close()
		handleError(err)
	}
}
