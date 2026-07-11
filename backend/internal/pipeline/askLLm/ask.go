package askllm

import (
	"log"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)


func AskLLM(docx []string, query string, h *models.Server) (string, error){
	var aiGeneratedResult string
	// get the proper prompt add thi docx to the prompt and then send 
	// whole of the augementation to the LLM and generate the result
	// get the data from the LLM and respond to the frontend

	promptTemplate, err := QueryAnswerPrompt(docx, query)
	if err != nil{
		log.Fatal(err)
	}

	// look how the prompt template seems?
	
	aiGeneratedResult, err = GenerateResult(promptTemplate, h)
	return aiGeneratedResult, nil
	
}