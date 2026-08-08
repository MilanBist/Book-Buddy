package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbqueries "github.com/MilanBist/AI-Powered-Book-Answerer/db/dbQueries"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

func(h *Handler) HandleGettingBooks(w http.ResponseWriter, r *http.Request){	
	// getting the userId
	userId := r.Context().Value("userId").(int)

	// get all the books in the format of the models
	var ResponseForBooks []models.BookData

	// if error occured in fetching the data from the database
	ResponseForBooks, err := dbqueries.GetAllBooksBasedOnUserId(userId, h.server.PostgresDB)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("[GETTING BOOKS HANDLER]: Error in inserting in db. Actual error: ", err)
		response := models.APIResponse{
			Success: false,
			Message: "Can't fetch the data from the db.",
		}
		json.NewEncoder(w).Encode(&response)
	}

	fmt.Println("[GETTING BOOKS HANDLER]: Success in fetching the data.", err)
	w.WriteHeader(http.StatusAccepted)
	response := models.APIResponse{
		Success: true,
		Message: "Data successfully fetched.",
		Data: ResponseForBooks,
	}
	json.NewEncoder(w).Encode(&response)

}