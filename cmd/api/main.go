package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type product struct {
	ID   int
	Name string
}

// Utilizando o FrameWork GO-Chi para rotas web para essa Api
func main() {
	r := chi.NewRouter()
	//Localhost:3000/?Name=name&id=5
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Name := r.URL.Query().Get("name")
		id := r.URL.Query().Get("id")
		w.Write([]byte(Name + id))
	})

	//Localhost:3000/Product
	r.Get("/{Rota}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "Rota")
		w.Write([]byte(param))
	})

	//Localhost:3000/json
	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		//Cria uma Lista e converte para retorna json com o go-chi/render
		obj := map[string]string{"mensagem": "sucesso"}
		render.JSON(w, r, obj)
	})


	//Localhost:3000/product
	r.Post("/json", func(w http.ResponseWriter, r *http.Request) {
		var product product
		render.DecodeJSON(r.Body, &product)
		product.ID = 5
		render.JSON(w, r, product)
	})

	//Localhost:Porta/
	http.ListenAndServe(":3000", r)
}
