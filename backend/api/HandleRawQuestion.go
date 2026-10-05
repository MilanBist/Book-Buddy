package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

type RawQuestionStore interface{
	GetTop5Chats(userId int, bookId int) ([]models.Messages, error)
	AddConversation(userQuery, response string, userId, bookId int) (int, error)
}

type RawQuestionEmbeddings interface{
	GenerateEmebedding(splittedDocx string) ([]float32, error)
	FindBestEmbeddings(point []float32, bookName string) ([]string, error)
}

type QuestionResponse interface{
	AskLLM(docx []string, preChats []models.Messages, query string, language string, w http.ResponseWriter, r *http.Request, userId int, bookId int) (string, error)
}

type QuestionHanlder struct{
	StoreQuestion 	RawQuestionStore
	AnswerQuestion 	RawQuestionEmbeddings
	Response 		QuestionResponse
}

func(h *QuestionHanlder) HandleRawQuestion(w http.ResponseWriter, r *http.Request){
	// get user data and decode it
	var userQuery *models.Question
	json.NewDecoder(r.Body).Decode(&userQuery)

	// check the userData
	fmt.Println("[HANDLING RAW QUESTION]: User query",userQuery)

	// if the user query is empty send user error as bad request
	if userQuery.Query == ""{
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: "No query.",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}


	// using bge-m3 is multilingual embedding generator so mostly same embedding is generated for the same thing in different language
	ans, err := h.AnswerQuestion.GenerateEmebedding(userQuery.Query)
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

	requiredDocx, err := h.AnswerQuestion.FindBestEmbeddings(ans, userQuery.BookName)
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
	top5Chats, err := h.StoreQuestion.GetTop5Chats(userId, userQuery.BookId)

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
	newResponse, err := h.Response.AskLLM(requiredDocx,top5Chats, userQuery.Query, userQuery.Language, w, r, userId, userQuery.BookId)
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

	_, err = h.StoreQuestion.AddConversation(userQuery.Query, newResponse, userId, userQuery.BookId)
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