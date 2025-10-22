package endpoints

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) CampaignDelete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	err := h.CampaignService.Delete(id)
	fmt.Println("Passou aqui")

	return nil, http.StatusNoContent, err
}
