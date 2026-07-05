package models

import (
	"database/sql"
	"github.com/MilanBist/AI-Powered-Book-Answerer/config"
	"github.com/go-chi/chi/v5"
	"github.com/qdrant/go-client/qdrant"
)



type Server struct{
	Router		*chi.Mux
	Store		*qdrant.Client
	Config		*config.Config
	SqliteDB 	*sql.DB
}

type Question struct{
	Query string
}

type Login struct{
	Email string	`json:"userEmail"`
	Password string	`json:"userPassword"`
}

type Register struct{
	FirstName string	`json:"userFirstName"`
	LastName string		`json:"userLastName"`
	Address string		`json:"userAddress"`
	Email string		`json:"userEmail"`
	Password string		`json:"userPassword"`
}