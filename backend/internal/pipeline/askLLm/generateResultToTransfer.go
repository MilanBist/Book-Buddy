package askllm

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	dbqueries "github.com/MilanBist/AI-Powered-Book-Answerer/db/dbQueries"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GenerateResultAndSendToFrontend(query, prompt string, h *models.Server, w http.ResponseWriter, r *http.Request, userId, bookId int, db *pgxpool.Pool)(error){
	// get the ollma model for the answer generation
	model := h.Config.OllamaTranslationModel
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive") // optional, often harmless

	flusher := w.(http.Flusher)

	llm, err := ollama.New(
		ollama.WithModel(model),
	)
	if err != nil{
		return errors.New("Error in connecting to the model.")
	}

	// make a final chunk
	var finalChunk bytes.Buffer

	_, err = llms.GenerateFromSinglePrompt(
		context.Background(),
		llm,
		prompt,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			// I will get the chunk of the data now just send this one to the frontend
			// return to the server
			_, err := w.Write(chunk)

			if err != nil{
				return errors.New("Error in writing chunk to frontend.")
			}
			finalChunk.Write(chunk)
			flusher.Flush()
			return nil
		}))

	if err != nil{
		return errors.New("Error in completing the response.")
	}

	// add the conversation to the database
	dbqueries.AddConversation(query, finalChunk.String(), userId, bookId, db)
	return nil
}