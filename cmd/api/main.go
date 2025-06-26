package main

import (
	"campaing/internal/domain/campaign"
	"campaing/internal/endpoint"
	"campaing/internal/infraestrutura/database"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// Utilizando o FrameWork GO-Chi para rotas web para essa Api
func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoint.Handler{
		CampaignService: campaignService,
	}

	//Localhost:3000/Campaign
	r.Post("/Campaign", endpoint.HandlerError(handler.CampaignPost))
	r.Get("/Campaign", endpoint.HandlerError(handler.CampaignGet))

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
