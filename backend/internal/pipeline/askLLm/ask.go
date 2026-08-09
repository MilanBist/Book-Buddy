package askllm

import (
	"net/http"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)


func AskLLM(docx []string, preChats []models.Messages, query string, h *models.Server, language string, w http.ResponseWriter, r *http.Request, userId, bookId int, db *pgxpool.Pool) (error){

	promptTemplate := QueryAnswerPrompt(docx, query,preChats, language)
	// look how the prompt template seems?
	
	err := GenerateResultAndSendToFrontend(query, promptTemplate, h, w, r, userId, bookId, db)

	if err != nil{
		return err
	}
	return nil
}