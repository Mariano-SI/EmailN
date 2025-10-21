package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) CampaignCancelPatch(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	err := h.CampaignService.Cancel(id)

	return nil, http.StatusOK, err
}
