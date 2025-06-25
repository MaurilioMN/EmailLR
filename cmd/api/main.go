package main

import (
	"campaing/internal/contract"
	"campaing/internal/domain/campaign"
	"campaing/internal/infraestrutura/database"
	internalerrors "campaing/internal/internalErrors"
	"errors"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Utilizando o FrameWork GO-Chi para rotas web para essa Api
func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	//Localhost:3000/
	r.Post("/Campaign", func(w http.ResponseWriter, r *http.Request) {

		var request contract.NewCampaignDto
		render.DecodeJSON(r.Body, &request)
		id, err := service.Create(request)

		if err != nil {

			if errors.Is(err, internalerrors.Errinternal) {
				render.Status(r, 500)
				render.JSON(w, r, map[string]string{"error": err.Error()})
				return
			}else{
				render.Status(r, 400)
				render.JSON(w, r, map[string]string{"id": id})
				return
			}


		}


		render.Status(r, 201)
		render.JSON(w, r, map[string]string{"id": id})
	})

	//Localhost:Porta/
	http.ListenAndServe(":3000", r)
}

//MiddleWares:
// Um middleware é uma função intermediária que é executada
// entre a requisição do cliente e a resposta do servidor.

// func myMiddleware(next http.Handler) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		println("Antes")
// 		next.ServeHTTP(w, r)
// 		println("depois")
// 	})

// }
