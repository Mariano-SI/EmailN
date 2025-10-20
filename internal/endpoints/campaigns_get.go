package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) {
	campaigns := h.CampaignService.Get()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, campaigns)
}
