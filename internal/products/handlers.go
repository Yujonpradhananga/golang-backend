package products

import (
	"net/http"

	"github.com/Yujonpradhananga/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := struct {
		Products []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK, products)
}
