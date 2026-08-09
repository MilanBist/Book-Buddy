package models

import (
	"github.com/MilanBist/AI-Powered-Book-Answerer/config"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qdrant/go-client/qdrant"
)



type Server struct{
	Router		*chi.Mux
	Store		*qdrant.Client
	Config		*config.Config
	PostgresDB 	*pgxpool.Pool
}

type Question struct{
	Query 		string	`json:"query"`
	Language 	string	`json:"language"`
	BookId		int		`json:"bookId"`
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
	Code	int		`json:"code"`
	Message	string	`json:"message"`
}
type SuccessResponse struct{
	Success		bool	`json:"success"`
	SuccessMsg	success	`json:"scx"`
}

// for each and every api response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}
type AuthData struct {
	Token string `json:"token"`
}

type BookData struct{
	BookId	int 	`json:"bookId"`
	BookName string	`json:"bookName"`
}

type Messages struct{
	UserQuestion	string	`json:"userQuestion"`
	LLMResponse		string	`json:"llmResponse"`
}

type ForConversation struct{
	BookId 		int		`json:"bookId"`
	BookName	string	`json:"bookName"`
}