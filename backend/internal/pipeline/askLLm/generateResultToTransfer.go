package askllm

import (
	"context"
	"fmt"
	"net/http"
	// "strings"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GenerateResultAndSendToFrontend(prompt string, h *models.Server, w http.ResponseWriter, r *http.Request)(error){
	// get the ollma model for the answer generation
	model := h.Config.OllamaModel
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive") // optional, often harmless

	flusher := w.(http.Flusher)

	fmt.Print("Reaching here to send to frontend.");

	llm, err := ollama.New(
		ollama.WithModel(model),
	)
	if err != nil{
		fmt.Println("Error in loading ollama model.")
		return err
	}


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
				// return the error in writing the chunk
				return err
			}
			fmt.Print(string(chunk))
			flusher.Flush()
			return nil
		}))

	if err != nil{
		fmt.Println("Error in completing the response from the llm.")
		return err
	}

	return nil

}