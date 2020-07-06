# generics-proposal

連結リストの実装でGo言語のジェネリクスの[ドラフト](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md#omissions) を触ってみる

## How to Use It

#### [go2go](https://go.googlesource.com/go/+/refs/heads/dev.go2go/README.go2go.md) をインストール＆ビルド

```
// ディレクトリを作成
$ mkdir go2go && cd go2go
// ソースをクローン
$ git clone https://go.googlesource.com/go goroot && cd goroot
// 最新ブランチをチェックアウト
$ git fetch && git checkout dev.go2go
// ソースからビルド
$ cd src && ./make.bash
```

#### コードを実行する
```
// 環境変数をセット
$ export PATH="$HOME/go2go/goroot/bin:$PATH"
// 実行する
$ go tool go2go run main.go2
```
