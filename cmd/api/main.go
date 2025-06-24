package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/",func (w http.ResponseWriter, r *http.Request)  {
		Name := r.URL.Query().Get("name")
		id := r.URL.Query().Get("id")
		w.Write([]byte(Name + id))
	})
	r.Get("/{Product}",func (w http.ResponseWriter, r *http.Request)  {
		param := chi.URLParam(r, "Product")
		w.Write([]byte(param))
	})
	http.ListenAndServe(":3000", r)
}
