package main

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaing"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaing.Service{}
	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaing
		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			println(err)
		}
		id, err := service.Create(request)

		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]string{"id": id})

	})

	http.ListenAndServe(":3000", r)
}
