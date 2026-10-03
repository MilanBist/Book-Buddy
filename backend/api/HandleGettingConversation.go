package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

type UserConversation interface{
	GetAllChats(userId, bookId int) ([]models.ReturningConversation, error)
}
type UserConversationHandler struct{
	Conversation 	UserConversation
}

func(h *UserConversationHandler) HandleConversation(w http.ResponseWriter, r *http.Request){
	// get the request as the bookId and the userId from the context
	userId := r.Context().Value("userId").(int)
	bookId := r.URL.Query().Get("bookId")
	bookName := r.URL.Query().Get("bookName")

	intBookId, _ := strconv.Atoi(bookId)

	fmt.Println("[HANDLE CONVERSATION]: Book information: \n Book id: ",bookId)
	fmt.Println("Book name: ", bookName)


	// after gaining the bookInformation specifically Book Id and userId
	allConversation, err := h.Conversation.GetAllChats(userId, intBookId)

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

	fmt.Println("[HANDLE CONVERSATION]: ",allConversation)

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