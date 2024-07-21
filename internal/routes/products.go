package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Ath3r/go-ecom.git/internal/models"
	"github.com/go-chi/chi/v5"
)



func ProductRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", handleGetProducts)
	r.Get("/{id}", handleGetProduct)
	return r
}

func handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products := []models.Product{
		{ID: "1"},
		{ID: "2"},
		{ID: "3"},
	}
	json.NewEncoder(w).Encode(products)
}

func handleGetProduct(w http.ResponseWriter, r *http.Request) {
	vars := chi.URLParam(r, "id")
	product := models.Product{ID: vars}
	json.NewEncoder(w).Encode(product)
}
