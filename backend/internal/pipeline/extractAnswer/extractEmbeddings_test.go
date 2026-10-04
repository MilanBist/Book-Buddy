package extractanswer

import (
	"testing"

	"github.com/MilanBist/AI-Powered-Book-Answerer/config"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)


func TestGenerateEmbeddings(t *testing.T){
	server := &models.Server{
		Config: &config.Config{
			OllamaEmbeddingModel: "bge-m3:latest",
		},
	}

	emb := &Embeddings{
		Server: server,
	}

	docx := "My name is Milan."
	embeddingValue, err := emb.GenerateEmebedding(docx)
	if err != nil{
		t.Fatalf("Error in getting the embedding. %v", err)
	}

	if len(embeddingValue) == 0{
		t.Fatalf("Can't embed the value.")
	}
	t.Logf("Embedding dimensions: %d", len(embeddingValue))
}