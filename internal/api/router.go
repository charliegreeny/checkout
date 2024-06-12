package api

import (
	"net/http"

	"github.com/charliegreeny/checkout/pkg/cart"
	"github.com/charliegreeny/checkout/pkg/item"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func itemRoutes(r chi.Router, i *item.Handler) chi.Router {
	r.Get("/", i.GetAllItemsHandler)
	r.Get("/{id}", i.GetItemHandler)
	return r
}

func cartRoutes(r chi.Router, c *cart.Handler) chi.Router {
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
	r.Use(responseHeader)

	r.Mount("/product", itemRoutes(r, i))
	r.Mount("/cart", cartRoutes(r, c))

	http.ListenAndServe(":8080", r)
}

func responseHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}