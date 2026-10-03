package extractanswer

import (
	"context"
	"errors"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/tmc/langchaingo/llms/ollama"
)

type Embeddings struct{
	Server 	*models.Server
}


func(s *Embeddings) GenerateEmebedding(splittedDocx string) ([]float32, error) {
	// create the ollama client
	llm, err := ollama.New(
		ollama.WithModel(s.Server.Config.OllamaEmbeddingModel),
	)

	if err != nil{
		return nil, errors.New("Error in splitting the docx.")
	}

	newString := []string{splittedDocx}

	// now generate result with the given llm
	embedding, err := llm.CreateEmbedding(context.Background(), newString)
	if err != nil{
		return  nil, errors.New("Error in creating the embeddings.")
	}
	singleEmbedding := embedding[0]
	
	return singleEmbedding, nil
}
