package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	Port					string
	FrontendUrl				string
	OllamaEndPoint			string
	OllamaModel				string
	OllamaEmbeddingModel	string
	QdrantPort				string
}


// return the configuration file
func LoadConfig() *Config{

	//load the .env file
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading the env file. Error: ",err)
	}
	cfg := &Config{
		Port: os.Getenv("PORT"),
		FrontendUrl: os.Getenv("FRONTEND_URL"),
		OllamaEndPoint: os.Getenv("OLLAMA_ENDPOINT"),
		OllamaModel: os.Getenv("OLLAMA_MODEL"),
		OllamaEmbeddingModel: os.Getenv("OLLAMA_EMBEDDING_MODEL"),
		QdrantPort: os.Getenv("QDRANT_PORT"),
	}

	// if models are nil then just figure them
	if cfg.OllamaEmbeddingModel == "" {
		cfg.OllamaEmbeddingModel = "nomic-embed-text"
	}
	if cfg.OllamaModel == ""{
		cfg.OllamaModel = "llama3.2"
	}

	return cfg
}