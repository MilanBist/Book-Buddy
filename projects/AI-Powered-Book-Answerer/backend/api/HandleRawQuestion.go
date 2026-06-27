package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	askllm "github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/askLLm"
	extractanswer "github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/extractAnswer"
)


// type llmResponse map[string]string

// question must be sent in the json format

func(h *Handler) HandleRawQuestion(w http.ResponseWriter, r *http.Request){
	// get the data from the form and work upon it

	// Answer the question by the user
	var userQuery models.Question
	// now decode the data
	json.NewDecoder(r.Body).Decode(&userQuery)


	// as the question is decoded
	// make a combined Augementation with the user query + ceratin format
	// fetch the all best 10 chunks and send it to the llm model
	// get the response here in this position and return to the frontend

	// send this data to get the vector embeddings of it

	// if the query is just nil
	if userQuery.Query == ""{
		w.Write([]byte("Empty query."))
		return
	}
	ans, err := extractanswer.GenerateEmebedding(userQuery.Query, h.server)
	if err != nil{
		log.Fatal("Error in generating the embeddings.")
	}
	// fmt.Println(ans)


	// with the help of the given ans as the embeddings now return the 
	// top 10 splitted docx realted to this one

	requiredDocx, err := extractanswer.FindBestEmbeddings(h.server.Store, ans)
	if err != nil{
		log.Fatal("Error in extracting the answer. ", err)
	}

	log.Println("Required Docx.")
	// generate answers based on the given docx
	llmResponse, err := askllm.AskLLM(requiredDocx, userQuery.Query, h.server)


	fmt.Println(llmResponse)

	// return this response to the frontend
	result := map[string]string{"Response": llmResponse}
	err = json.NewEncoder(w).Encode(result)
	if err != nil{
		fmt.Println("Error in encoding the response.")
		return
	}

	fmt.Println("[RESPONSE SENT] to the frontend.")
	
}