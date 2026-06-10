package extractpdf

import (
	"context"
	"fmt"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/tmc/langchaingo/llms/ollama"
)

func createEmbeddings(splittedDocx []string, s *models.Server) error {
	llm, err := ollama.New(
		ollama.WithModel(s.Config.OllamaModel),
	)

	if err != nil{
		fmt.Println("Error in splitting docx.")
		return err
	}

	// now generate result with the given llm
	response, err := llm.CreateEmbedding(context.Background(), 
		splittedDocx,
	)
	if err != nil{
		fmt.Println("Error in creating embeddings.")
		return err
	}

	// if not just print the response
	fmt.Println(response)
	fmt.Println(len(response[0]))
	return nil
}

func vectorStore(splittedDocx []string, s *models.Server)error{
	// first create the embedding of each of the given docx
	createEmbeddings(splittedDocx, s)
	return nil
}