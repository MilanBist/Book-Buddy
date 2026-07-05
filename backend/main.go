package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/MilanBist/AI-Powered-Book-Answerer/api"
	"github.com/MilanBist/AI-Powered-Book-Answerer/config"
	"github.com/MilanBist/AI-Powered-Book-Answerer/db"
)


func main(){
	// add the configuration files
	cfg := config.LoadConfig()

	// add the vector data base with certain things
	vectorStore, err := db.LoadVectorDatabase(cfg)

	if err != nil{
		log.Fatal(err)
	}


	// if successfully created a new client just pass the vector store and cfg file to create server
	server := api.NewServer(vectorStore, cfg)


	// configure port
	port := ":"+cfg.Port

	fmt.Println("Server successfully started! Listening on http://localhost:8080...")

	if err = http.ListenAndServe(port, server.Router); err != nil{
		log.Fatal("Can't start the http server, Error: ",err)
	}


}