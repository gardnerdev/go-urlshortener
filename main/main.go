package main

import (
	"fmt"
	"net/http"

	"github.com/gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/bitcoin": "https://www.youtube.com/watch?v=s4g1XFU8Gto",
		"/luna":    "https://www.youtube.com/watch?v=KqpGMoYZMhY",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on: 8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", gm)
	return mux
}

func gm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GM ;)")
}
