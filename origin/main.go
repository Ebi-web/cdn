package main

import (
	"log"
	"net/http"
)

func main() {
	// 静的ファイルのディレクトリを設定
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	http.Handle("/404", http.NotFoundHandler())

	// サーバーをポート8080で起動
	log.Println("Listening on http://localhost:8080...")
	err := http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
