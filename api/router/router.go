package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"indexer.com/indexer/api/handler/search_handler"
)

func InitializeRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	r.Post("/search", search_handler.PostSearch)
	return r
}
