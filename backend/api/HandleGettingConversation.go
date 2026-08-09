package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbqueries "github.com/MilanBist/AI-Powered-Book-Answerer/db/dbQueries"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

func(h *Handler) HandleConversation(w http.ResponseWriter, r *http.Request){
	// get the request as the bookId and the userId from the context
	userId := r.Context().Value("userId").(int)


	var bookInformation models.ForConversation
	json.NewDecoder(r.Body).Decode(&bookInformation)

	// after gaining the bookInformation specifically Book Id and userId
	err, allConversation := dbqueries.GetAllChats(userId, bookInformation.BookId, h.server.PostgresDB)

	if err != nil{
		fmt.Println("[GETTING CONVERSATION ENDPOINT]: Error in getting the conversation.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	// return the response of all conversation
	response := models.APIResponse{
		Success: true,
		Message: "Conversation Data fetched successfully.",
		Data: allConversation,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(&response)

}