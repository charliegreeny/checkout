package cart

import (
	"encoding/json"
	"net/http"

	"github.com/charliegreeny/checkout/pkg/model"
	"github.com/go-chi/chi"
)

type Handler struct {
	service model.Service[*Entity]
}

func NewHandler(s model.Service[*Entity]) *Handler {
	return &Handler{s}
}

func (c Handler) CreateCartHandler(w http.ResponseWriter, r *http.Request) {

}

func (c Handler) AddItemToCart(w http.ResponseWriter, r *http.Request) {

}

func (c Handler) GetCartHandler(w http.ResponseWriter, r *http.Request) {
	e, err := c.service.GetById(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "cart not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(e)
}
