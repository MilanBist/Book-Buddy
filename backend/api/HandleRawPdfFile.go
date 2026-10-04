package api

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

type StoreFileToLocation interface{
	StoreFile(fileName string, file multipart.File)(string, error)
}

type StoreRawFile interface{
	AddBookIdentity(bookTitle string, userId int) (int, int, error)
}

type Extract interface{
	ExtractData(file *multipart.File, filePath string) (int, error)
}

type RawPdfHandler struct{
	GetFileLocation 	StoreFileToLocation
	StoreBook 			StoreRawFile
	ExtractPdf 			Extract
}

func(h *RawPdfHandler) HandleRawPdfFile(w http.ResponseWriter, r *http.Request){
	// handle raw pdf files
	file,header, err := r.FormFile("document")
	if err != nil{
		http.Error(w, "Error in receiving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// get the file name
	fileName := header.Filename

	// check if the extension is pdf or not
	if strings.ToLower(filepath.Ext(fileName)) != ".pdf" {
		http.Error(w, "Only PDF files are allowed", http.StatusBadRequest)
		return
	}


	fullPath, err := h.GetFileLocation.StoreFile(fileName, file)

	// send this file to the pipeline ->
	// i. Extract the text from the pdf
	// ii. Converting the extracted text to a single string
	// iii. Split this string
	// iv. Create embeddings and store in vector db

	status, err := h.ExtractPdf.ExtractData(&file, fullPath)
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
	fileName = filePaths[len(filePaths)-1]

	// get the file name
	newFilePath := strings.Split(fileName, ".")
	fileName = newFilePath[0]

	userId := r.Context().Value("userId").(int)
	// add to the database
	status, bookId, err := h.StoreBook.AddBookIdentity(fileName, userId)
	if err != nil{
		fmt.Println("[PDF HANDLER]: Error in adding the book to db.")
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
		Data: models.BookData{
			BookId: bookId,
			BookName: fileName,
		},
	}
	json.NewEncoder(w).Encode(&response)
}