package models

import (
	"github.com/MilanBist/AI-Powered-Book-Answerer/config"
	"github.com/go-chi/chi/v5"
	"github.com/qdrant/go-client/qdrant"
)



type Server struct{
	Router	*chi.Mux
	Store	*qdrant.Client
	Config	*config.Config
}

type Question struct{
	Query string
}