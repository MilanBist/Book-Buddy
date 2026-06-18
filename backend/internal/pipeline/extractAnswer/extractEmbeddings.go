package extractanswer

import (
	"context"
	"fmt"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GenerateEmebedding(splittedDocx string, s *models.Server) ([]float32, error) {
	// create the ollama client
	llm, err := ollama.New(
		ollama.WithModel(s.Config.OllamaModel),
	)

	if err != nil{
		fmt.Println("Error in splitting docx.")
		return nil,err
	}

	newString := []string{splittedDocx}

	// now generate result with the given llm
	embedding, err := llm.CreateEmbedding(context.Background(), newString)
	if err != nil{
		return  nil, err
	}
	singleEmbedding := embedding[0]
	
	return singleEmbedding, nil
}
