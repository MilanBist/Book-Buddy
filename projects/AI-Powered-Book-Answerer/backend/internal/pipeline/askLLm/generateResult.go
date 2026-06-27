package askllm

import (
	"context"
	"fmt"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)


func GenerateResult(prompt string, h *models.Server) (string, error){
	// get the ollma model for the answer generation
	model := h.Config.OllamaModel


	llm, err := ollama.New(
		ollama.WithModel(model),
	)
	if err != nil{
		fmt.Println("Error in loading ollama model.")
		return "", err
	}

	completion, err := llms.GenerateFromSinglePrompt(
		context.Background(),
		llm,
		prompt,
		llms.WithTemperature(0.8),
	)

	if err != nil{
		fmt.Println("Error in completing the response from the llm.")
		return "", err
	}


	return completion, err

}