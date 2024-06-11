package api

import (
	"net/http"

	"github.com/charliegreeny/checkout/pkg/cart"
	"github.com/charliegreeny/checkout/pkg/item"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func itemRoutes(i *item.Handler) chi.Router {
	r := chi.NewRouter()
	r.Get("/", i.GetAllItemsHandler)
	r.Get("/{id}", i.GetItemHandler)
	return r
}

func cartRoutes(c *cart.Handler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.CreateCartHandler)
	r.Post("/cart/{cartId}/item/{itemId}", c.AddItemToCart)
	r.Get("/{id}", c.GetCartHandler)
	return r
}

func StartRouter(i *item.Handler, c *cart.Handler) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/product", itemRoutes(i))
	r.Mount("/cart", cartRoutes(c))

	http.ListenAndServe(":8080", r)
}
