# file-upload

GO言語のmime/multipartパッケージでファイルをアップロード

## How to Use It

### サーバー＆クライアントを起動する
```
go run server/server.go

# 別のターミナルを開く
go run client/client.go
```

* サーバーのターミナルでリクエストのデバッグ情報が出力される
* `uploaded_1543223947765718937.txt`のような名前のあるファイルが作成される

### ベンチマーク

```
go run server/server.go -in-mem

# 別のターミナルを開く
go test -bench=. ./client -benchmem
```
