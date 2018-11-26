package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	url := "http://localhost:3000/upload"
	fieldname := "file"
	filename := "hello.txt"
	file, err := os.Open(filename)
	handleError(err)

	// リクエストボディのデータを受け取るio.Writerを生成する。
	body := &bytes.Buffer{}

	// データのmultipartエンコーディングを管理するmultipart.Writerを生成する。
	// ランダムなbase-16バウンダリが生成される。
	mw := multipart.NewWriter(body)

	// ファイルに使うパートを生成する。
	// ヘッダ以外はデータは書き込まれない。
	// fieldnameとfilenameの値がヘッダに含められる。
	// ファイルデータを書き込むio.Writerが返却される。
	fw, err := mw.CreateFormFile(fieldname, filename)

	// fwで作ったパートにファイルのデータを書き込む
	_, err = io.Copy(fw, file)
	handleError(err)

	// リクエストのContent-Typeヘッダに使う値を取得する（バウンダリを含む）
	contentType := mw.FormDataContentType()

	// 書き込みが終わったので最終のバウンダリを入れる
	err = mw.Close()
	handleError(err)

	// contentTypeとbodyを使ってリクエストを送信する
	resp, err := http.Post(url, contentType, body)
	handleError(err)

	err = resp.Body.Close()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
