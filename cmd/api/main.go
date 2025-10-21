package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.ServiceImp{Repository: &database.CampaignRepository{}}
	handler := endpoints.Handler{CampaignService: &campaignService}

	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))
	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))

	fmt.Println("🚀 Servidor rodando na porta 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
