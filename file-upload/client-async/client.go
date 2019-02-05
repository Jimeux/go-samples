package main

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

var (
	uploadURL = "http://localhost:3000/upload"
	fieldName = "file"
	fileName  = "hello.txt"
)

func main() {
	dir, err := os.Getwd()
	handleError(err)

	file, err := os.Open(path.Join(dir, fileName))
	handleError(err)

	uploadFileAsync(file, uploadURL, fieldName, fileName)
}

func uploadFileAsync(file io.Reader, url, fieldName, fileName string) {
	pr, pw := io.Pipe()
	// io.PipeWriterをmultipart.Writerに渡す
	mw := multipart.NewWriter(pw)

	// Goルーチンを開始し非同期でpwに書き込む
	go func() {
		// prの処理が正常に終わるように必ずpwをクローズする
		defer pw.Close()
		defer mw.Close()
		fw, err := mw.CreateFormFile(fieldName, fileName)
		if err != nil {
			return
		}
		if _, err := io.Copy(fw, file); err != nil {
			return
		}
	}()

	// io.PipeReaderをHTTPリクエストのボディに渡す
	resp, err := http.Post(url, mw.FormDataContentType(), pr)
	handleError(err)

	err = resp.Body.Close()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
