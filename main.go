package main

import (
	"fmt"
	"net/http"
)

func main() {
	defaultMux := defaultMux()
	urlsToPath := map[string]string{
		"/google":  "https://google.com",
		"/youtube": "https://youtube.com",
	}

	http.ListenAndServe("localhost:8080", MapUrlstoPath(urlsToPath, defaultMux))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	return mux
}

func MapUrlstoPath(urlsToPath map[string]string, fallback *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path
		if url, ok := urlsToPath[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}

}
