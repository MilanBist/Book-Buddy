package api

import (
	"encoding/json"
	"fmt"
	// "log"

	// "log"
	"net/http"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	askllm "github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/askLLm"
	extractanswer "github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/extractAnswer"
	"github.com/MilanBist/AI-Powered-Book-Answerer/utils"
)

func(h *Handler) HandleRawQuestion(w http.ResponseWriter, r *http.Request){
	// get user data and decode it
	var userQuery *models.Question
	json.NewDecoder(r.Body).Decode(&userQuery)

	// check the userData
	fmt.Println(userQuery)

	// if the query is just nil
	if userQuery.Query == ""{
		w.Write([]byte("Empty query."))
		return
	}

	language, err := askllm.CheckLanguage(userQuery.Query, h.server)
	if err != nil{
		var errmsg models.ErrorResponse
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg = utils.PrepareErrorMessage(err.Error())
		json.NewEncoder(w).Encode(&errmsg)
		return
	} 

	fmt.Println("The language is: ", language)
	
	if language != "English"{
		// now convert the given userQuery to the English string
		// convert the user query  to english
		englished, err := askllm.ConvertToEnglish(userQuery.Query, h.server)
		if err != nil{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			var errmsg models.ErrorResponse
			errmsg = utils.PrepareErrorMessage(err.Error())
			json.NewEncoder(w).Encode(&errmsg)
			return
		}		
	// 	now set the userQuery be english query
	// 	userQuery.Query = englishQuery
		userQuery.Query = englished
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("The english query is: ",userQuery.Language)
	fmt.Println()
	fmt.Println()
	fmt.Println()


	fmt.Println("UserQuery: ",userQuery)
	// now since the question asked is converted in the english

	// generating the embeddings of the user query
	ans, err := extractanswer.GenerateEmebedding(userQuery.Query, h.server)
	if err != nil{
		msg := utils.PrepareErrorMessage("Error in generating the embeddings.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	// with the help of the given ans as the embeddings now return the 
	// top 10 splitted docx realted to this one
	requiredDocx, err := extractanswer.FindBestEmbeddings(h.server.Store, ans)
	if err != nil{
		var errmsg models.ErrorResponse
		errmsg = utils.PrepareErrorMessage(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errmsg)
		return
	}


	// generate answers based on the given docx
	err = askllm.AskLLM(requiredDocx, userQuery.Query, h.server, userQuery.Language, w, r)
	if err != nil{
		var errmsg models.ErrorResponse
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg = utils.PrepareErrorMessage(err.Error())
		json.NewEncoder(w).Encode(&errmsg)
		return
	}

}