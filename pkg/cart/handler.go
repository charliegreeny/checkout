package cart

import (
	"encoding/json"
	"errors"
	"fmt"
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
	id := chi.URLParam(r, "id")
	e, err := c.service.GetById(id)
	if errors.As(err, &model.NotFoundErr{}) {
		http.Error(w, fmt.Sprintf("cart id %s not found", id), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("internal server error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(e)
}
