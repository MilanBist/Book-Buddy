package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbqueries "github.com/MilanBist/AI-Powered-Book-Answerer/db/dbQueries"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	askllm "github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/askLLm"
	extractanswer "github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/extractAnswer"
)

func(h *Handler) HandleRawQuestion(w http.ResponseWriter, r *http.Request){
	// get user data and decode it
	var userQuery *models.Question
	json.NewDecoder(r.Body).Decode(&userQuery)

	// check the userData
	fmt.Println("[HANDLING RAW QUESTION]: User query",userQuery)

	// if the user query is empty send user error as bad request
	if userQuery.Query == ""{
		http.Error(w, "Empty query", http.StatusBadRequest)
		return
	}


	// using bge-m3 is multilingual embedding generator so mostly same embedding is generated for the same thing in different language
	ans, err := extractanswer.GenerateEmebedding(userQuery.Query, h.server)
	if err != nil{
		fmt.Println("[HANDLE RAW QUESTION] Error: ", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	fmt.Println("Book name being used is: ", userQuery.BookName)
	requiredDocx, err := extractanswer.FindBestEmbeddings(h.server.Store, ans, userQuery.BookName)
	if err != nil{
		fmt.Println("[HANDLE RAW QUESTION] Error: ", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	// get the userId based on the context
	userId := r.Context().Value("userId").(int)
	// get top 10 previous chats on the basis of the given bookId and userId
	err, top5Chats := dbqueries.GetTop5Chats(userId, userQuery.BookId, h.server.PostgresDB)

	if err != nil{
		fmt.Println("[HANDLE RAW QUESTION] Error: ", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}


	fmt.Println("[HANDLE RAW QUESTION] Book Id: ", userQuery.BookId)
	// generate answers based on the given docx
	err = askllm.AskLLM(requiredDocx,top5Chats, userQuery.Query, h.server, userQuery.Language, w, r, userId, userQuery.BookId,  h.server.PostgresDB)
	if err != nil{
		fmt.Println("[HANDLE RAW QUESTION] Error: ", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
}