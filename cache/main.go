package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	origin, _ := url.Parse("http://origin:8080")

	proxy := httputil.NewSingleHostReverseProxy(origin)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("proxying to origin...")
		proxy.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8081", nil)
}
