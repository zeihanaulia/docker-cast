package main

import (
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome from customer api, using go and chi router!"))
	})
	http.ListenAndServe(":3000", r)
}