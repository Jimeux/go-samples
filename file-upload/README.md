# file-upload

Go言語のmime/multipartパッケージでファイルをアップロード

## How to Use It

#### サーバーを起動する
```
go run server/server.go
```

#### 同期的クライントによるリクエスト
```
go run client/client.go
```

#### `io.Pipe`クライントによるリクエスト
```
go run client-async/client.go
```

* サーバーのターミナルでリクエストのデバッグ情報が出力される
* リクエストごとに`uploaded_1543223947765718937.txt`のような名前のあるファイルが作成される

## ベンチマーク

#### テスト用のサーバーを起動する
```
go run server/server.go -in-mem
```

#### ベンチマークを実行する

```
go test -bench=. ./client -benchmem
go test -bench=. ./client-async -benchmem
```
