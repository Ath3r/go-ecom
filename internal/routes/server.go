package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

type Server struct {
	Router *chi.Mux
}

func CreateNewServer() *Server {
	server := &Server{}

	server.Router = chi.NewRouter()

	return server
}

func (s *Server) MountMiddlewares() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}).Handler)
	s.Router.Use(middleware.RequestID)

}

func (s *Server) MountHandlers() {
	s.MountMiddlewares()
	s.Router.Mount("/api/v1" , MountRoutes())
}

func MountRoutes() chi.Router {
	r := chi.NewRouter()
	r.Mount("/products", ProductRoutes())
	r.Mount("/auth", AuthRoutes())
	return r
}