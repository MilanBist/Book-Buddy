package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	dbqueries "github.com/MilanBist/AI-Powered-Book-Answerer/db/dbQueries"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/extractpdf"
)

func(h *Handler) HandleRawPdfFile(w http.ResponseWriter, r *http.Request){
	// handle raw pdf files
	file,header, err := r.FormFile("document")
	if err != nil{
		http.Error(w, "Error in receiving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// get the file name
	fileName := header.Filename

	fmt.Println("[PDF HANDLER]: BookName: ",fileName)

	// check if the extension is pdf or not
	if strings.ToLower(filepath.Ext(fileName)) != ".pdf" {
		http.Error(w, "Only PDF files are allowed", http.StatusBadRequest)
		return
	}

	// save the file to the location of certain place by saying it documents
	folderPath := "./uploadedFiles"
	entries, err := os.ReadDir(folderPath)
	log.Println(entries)
	info, err := os.Stat(folderPath)
	if err != nil{
		if info == nil{
		err = os.Mkdir("./uploadedFiles", 0755)
		if err != nil{
			fmt.Println("[PDF HANDLER] Error in creating folder: ", err)
			}
		}
	}
	// cerate the file
	fullPath := folderPath + "/" + fileName
	copiedFile, err := os.Create(fullPath)
	if err != nil{
		http.Error(w, "Error in creating the filepath", http.StatusInternalServerError)
		return
	}

	// make a copy of the file inside the given folder made above
	_, err = io.Copy(copiedFile, file)
	if err != nil{
		http.Error(w, "Error in copying the file", http.StatusInternalServerError)
		return
	}

	// send this file to the pipeline ->
	// i. Extract the text from the pdf
	// ii. Converting the extracted text to a single string
	// iii. Split this string
	// iv. Create embeddings and store in vector db

	status, err := extractpdf.ExtractData(&file, fullPath, h.server)
	fmt.Println("Error: ", err)
	if err != nil{
		fmt.Println("[PDF HANDLER]: Error in extracting the pdf.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		// now add the error message
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}


	// for the full path
	filePaths := strings.Split(fullPath, "/")
	fileName = filePaths[len(fileName)-2]

	userId := r.Context().Value("userId").(int)
	// add to the database
	status, err = dbqueries.AddBookIdentity(fileName, userId, h.server.PostgresDB)

	if err != nil{
		fmt.Println("[PDF HANDLER]: Error in extracting the pdf.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		// now add the error message
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	// if no nil just return extraction and saving compelted
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// set the response
	response := models.APIResponse{
		Success: true,
		Message: "Book added to database successfully.",
	}
	json.NewEncoder(w).Encode(&response)
}