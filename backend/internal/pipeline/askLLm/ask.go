package askllm

import (
	"fmt"
	"net/http"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)


func AskLLM(docx []string, preChats []models.Messages, query string, h *models.Server, language string, w http.ResponseWriter, r *http.Request, userId, bookId int, db *pgxpool.Pool) (error){

	promptTemplate := QueryAnswerPrompt(docx, query,preChats, language)	
	fmt.Println("Final prompt: ", promptTemplate)
	err := GenerateResultAndSendToFrontend(query, promptTemplate, h, w, r, userId, bookId, db)
	if err != nil{
		return err
	}
	return nil
}