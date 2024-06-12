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
	service model.IDGetterCreator[*createInput,*Entity]
}

func NewHandler(s model.IDGetterCreator[*createInput,*Entity]) *Handler {
	return &Handler{s}
}

func (h Handler) CreateCartHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody *createInput 
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	e, err := h.service.Create(reqBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("internal server error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e.ToOutput())
}

func (h Handler) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	
}

func (h Handler) GetCartHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	e, err := h.service.GetById(id)
	if errors.As(err, &model.ErrNotFound{}) {
		http.Error(w, fmt.Sprintf("cart id %s not found", id), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("internal server error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(e.ToOutput())
}
