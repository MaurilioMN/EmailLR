package endpoint

import (
	"campaing/internal/contract"
	internalerrors "campaing/internal/internalErrors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {

	var request contract.NewCampaignDto
	render.DecodeJSON(r.Body, &request)
	id, err := h.CampaignService.Create(request)

	if err != nil {

		if errors.Is(err, internalerrors.Errinternal) {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		} else {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"id": id})
			return
		}

	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})


}
