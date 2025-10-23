package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDb()
	campaignRepository := database.CampaignRepository{Db: db}
	campaignService := campaign.ServiceImp{Repository: &campaignRepository}
	handler := endpoints.Handler{CampaignService: &campaignService}

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)
		r.Get("/", endpoints.HandlerError(handler.CampaignGet))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
		r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Println("🚀 Servidor rodando na porta 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
