package routes

import (
	"github.com/go-chi/chi"
	"github.com/charliegreeny/checkout/pkg/item"
	"github.com/charliegreeny/checkout/pkg/cart"
)

func Item() chi.Router{
	r := chi.NewRouter()
	r.Get("/", item.GetAllItemsHandler)
	r.Get("/{id}", item.GetItemHandler)
	return r
}

func Cart() chi.Router{
	r := chi.NewRouter()
	r.Post("/", cart.CreateCartHandler)
	r.Post("/cart/{cartId}/item/{itemId}", cart.AddItemToCart)
	r.Get("/{id}", cart.GetCartHandler)
	return r
}