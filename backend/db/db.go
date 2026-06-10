package db

import (
	"errors"
	"strconv"
	"github.com/MilanBist/AI-Powered-Book-Answerer/config"
	"github.com/qdrant/go-client/qdrant"
)

type Store struct{
	store	*qdrant.Client
}


func LoadVectorDatabase(cfg *config.Config) (*Store, error) {
	port, _ := strconv.Atoi(cfg.QdrantPort)
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: port,
	})
	if err != nil{
		return nil, errors.New("Error in creating client.")
	}

	//return the client created on local host and certain port
	store := &Store{
		store: client,
	}

	return store, nil
}