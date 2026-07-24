package askllm

import (
	"net/http"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)


func AskLLM(docx []string, query string, h *models.Server, language string, w http.ResponseWriter, r *http.Request) (error){
	// get the proper prompt add thi docx to the prompt and then send 
	// whole of the augementation to the LLM and generate the result
	// get the data from the LLM and respond to the frontend

	promptTemplate := QueryAnswerPrompt(docx, query, language)
	// look how the prompt template seems?
	
	err := GenerateResultAndSendToFrontend(promptTemplate, h, w, r)

	if err != nil{
		return err
	}
	return nil
}