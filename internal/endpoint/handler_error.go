package endpoint

import (
	internalerrors "campaing/internal/internalErrors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpointFunc(w, r)
		if err != nil {

			if errors.Is(err, internalerrors.Errinternal) {
				render.Status(r, 500)
				return
			} else {
				render.Status(r, 400)
				return
			}
		}
		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
