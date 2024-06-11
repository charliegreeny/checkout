package item

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (c Handler) GetAllItemsHandler(w http.ResponseWriter, r *http.Request) {

}

func (c Handler) GetItemHandler(w http.ResponseWriter, r *http.Request) {

}
