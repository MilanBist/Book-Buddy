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

func(h *Handler) HandleRawQuestion(w http.ResponseWriter, r *http.Request){
	// get the data from the form and work upon it

	// Answer the question by the user
	var userQuery *models.Question
	// now decode the data
	json.NewDecoder(r.Body).Decode(&userQuery)

	fmt.Println(userQuery)

	// if the query is just nil
	if userQuery.Query == ""{
		w.Write([]byte("Empty query."))
		return
	}
	// // convert the given query to english language
	// if userQuery.Language != "English"{
	// 	// first convert the language to the english language
	// }
	// check if the userQuery is of which language
	language, err := askllm.CheckLanguage(userQuery.Query, h.server)

	if err != nil{
		log.Println(err)
		return
	} 

	// after that I will get the language of the given prompt
	// if the language is english just continue if not just convert
	// the given provided language to english but semantic emotions in it
	// must be preserved

	if language != "English"{
		// now convert the given userQuery to the English string
		// convert the user query  to english
		englishQuery, err := askllm.ConvertToEnglish(userQuery.Query, h.server)
		if err != nil{
			log.Println(err)
			return
		}		

		// now set the userQuery be english query
		userQuery.Query = englishQuery
	}

	fmt.Println("UserQuery: ",userQuery)
	// now since the question asked is converted in the english

	ans, err := extractanswer.GenerateEmebedding(userQuery.Query, h.server)
	if err != nil{
		log.Fatal("Error in generating the embeddings.")
	}

	// with the help of the given ans as the embeddings now return the 
	// top 10 splitted docx realted to this one

	requiredDocx, err := extractanswer.FindBestEmbeddings(h.server.Store, ans)
	if err != nil{
		log.Fatal("Error in extracting the answer. ", err)
	}

	// generate answers based on the given docx
	llmResponse, err := askllm.AskLLM(requiredDocx, userQuery.Query, h.server)


	fmt.Println(llmResponse)

	// response the required docx with the language in which the user wants it
	if userQuery.Language != "English"{
		err = askllm.GenerateInRequiredLanguage(llmResponse, userQuery.Language, h.server, w, r)
		if err != nil{
			log.Println("Error in generating in required language.")
			return
		}
		return
	}

	// return this response to the frontend
	// result := map[string]string{"Response": llmResponse}
	// err = json.NewEncoder(w).Encode(result)
	// if err != nil{
	// 	fmt.Println("Error in encoding the response.")
	// 	return
	// }
}