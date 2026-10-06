package askllm

import (
	"net/http"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

type LLMWork struct{
	Server 	*models.Server
}

func(s *LLMWork) AskLLM(docx []string, preChats []models.Messages, query string, language string, w http.ResponseWriter, r *http.Request, userId, bookId int) (string, error){

	promptTemplate := QueryAnswerPrompt(docx, query,preChats, language)	
	response, err := GenerateResultAndSendToFrontend(query, promptTemplate, s.Server, w, r, userId, bookId)
	if err != nil{
		return "",err
	}
	return response, nil
}