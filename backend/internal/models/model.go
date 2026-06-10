package models

import (
	"github.com/MilanBist/AI-Powered-Book-Answerer/config"
	"github.com/MilanBist/AI-Powered-Book-Answerer/db"
	"github.com/go-chi/chi/v5"
)



type Server struct{
	Router	*chi.Mux
	Store	*db.Store
	Config	*config.Config
}