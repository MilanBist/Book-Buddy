package api

import (
	"net/http"
	"time"

	"github.com/MilanBist/AI-Powered-Book-Answerer/config"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/middlewares"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qdrant/go-client/qdrant"
)



type Handler struct{
	// declare the models here
	server *models.Server
}

func NewServer(vectorStore *qdrant.Client, cfg *config.Config, postgres *pgxpool.Pool) *models.Server{
	r := chi.NewRouter()


	// using the middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// use the middleware for the time if more than 60 abort
	r.Use(middleware.Timeout(60 *time.Second))

	// use the cors here
	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
  }))

  	server := &models.Server{
		Router: r,
		Store: vectorStore,
		Config: cfg,
		PostgresDB: postgres,
	}

	handler := &Handler{
		server: server,
	}

	handler.setupRoutes()
	return handler.server

  // mount over the api layer

}

func(h *Handler) setupRoutes(){
	// initally test with ping pong statement
	h.server.Router.Get("/ping", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("pong"));
	})
	h.server.Router.Route(
		"/api", func(r chi.Router) {
			r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Success in API call!"))
			})

			// for the login and registration
			r.Post("/login", h.HandleLogin)
			r.Post("/register", h.HandleRegister)


			// create the protected handlers
			r.Group(func(r chi.Router){
				// add the middlewares here
				r.Use(middlewares.LoggingMiddleware)
				r.Use(middlewares.AuthMiddleware)

				
				// extract pdf + create embeddings + store in vector db
				r.Post("/handlePdf", h.HandleRawPdfFile)
				// get the question -> create embedding -> extract relevant data from the vector db
				r.Post("/extractAnswer", h.HandleRawQuestion)
			})

			
		})
}