package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	cache, _ := url.Parse("http://cache:8081")

	proxy := httputil.NewSingleHostReverseProxy(cache)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("proxying to cache...")
		proxy.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8082", nil)
}
