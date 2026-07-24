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
	Query 		string	`json:"query"`
	Language 	string	`json:"language"`
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

type error struct{
	Code	int	`json:"code"`
	Message	string	`json:"message"`
}
type ErrorResponse struct{
	Success	bool	`json:"success"`
	ErrMsg	error	`json:"error"`
}

type success struct{
	Code	int	`json:"code"`
	Message	string	`json:"message"`
}
type SuccessResponse struct{
	Success		bool	`json:"success"`
	SuccessMsg	success	`json:"scx"`
}