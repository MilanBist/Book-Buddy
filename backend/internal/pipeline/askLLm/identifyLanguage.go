package askllm

import (
	"context"
	"fmt"
	// "strings"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GenerateLanguageResults(prompt string, h *models.Server)(string, error){
	// get the ollma model for the answer generation
	model := h.Config.OllamaTranslationModel
	fmt.Println("Reaching here to send to frontend.", prompt)

	llm, err := ollama.New(
		ollama.WithModel(model),
	)
	if err != nil{
		fmt.Println("Error in loading ollama model.")
		return "", err
	}


	result, err := llms.GenerateFromSinglePrompt(
		context.Background(),
		llm,
		prompt,
		llms.WithTemperature(0.8),
	)

	if err != nil{
		fmt.Println("Error in completing the response from the llm.")
		return "", err
	}

	return result, nil

}