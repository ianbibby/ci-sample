package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", handler(true))
	http.ListenAndServe(":8080", mux)
}

func handler(b bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if b {
			w.Write([]byte("foo"))
			return
		}
		w.Write([]byte("bar"))
	})
}
