# file-upload

GO言語のmime/multipartパッケージでファイルをアップロード

## How to Use It

サーバーを起動する
```
go run server/server.go
```

クライントを別のターミナルで実行する

```
go run client/client.go
```

* サーバーのターミナルでリクエストのデバッグ情報が出力される
* `uploaded_1543223947765718937.txt`のような名前のあるファイルが作成される
