package api

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	// detect the language of the prompt
	language, err := askllm.CheckLanguage(userQuery.Query, h.server)
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

	fmt.Println("[HANDLE RAW QUESTION]: Language detected is: ", language)
	

	// if the language is not english convert to the english and set the userQuery.query to be english prompt
	if language != "English"{
		englished, err := askllm.ConvertToEnglish(userQuery.Query, h.server)
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
		userQuery.Query = englished
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("[HANDLE RAW QUESTION]: The english query is: ",userQuery.Query)
	fmt.Println()
	fmt.Println()


	// generating the embeddings of the user query
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

	// with the help of the given ans as the embeddings now return the 
	// top 10 splitted docx realted to this one
	requiredDocx, err := extractanswer.FindBestEmbeddings(h.server.Store, ans)
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

	// generate answers based on the given docx
	err = askllm.AskLLM(requiredDocx, userQuery.Query, h.server, userQuery.Language, w, r, userId, userQuery.BookId,  h.server.PostgresDB)
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