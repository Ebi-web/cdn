package main

import (
	"log"
	"net/http"
)

func main() {
	// 静的ファイルのディレクトリを設定
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// サーバーをポート8080で起動
	log.Println("Listening on http://localhost:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
