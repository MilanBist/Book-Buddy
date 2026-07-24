package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/pipeline/extractpdf"
	"github.com/MilanBist/AI-Powered-Book-Answerer/utils"
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

	log.Println(fileName)

	// save the file to the location of certain place by saying it documents
	folderPath := "./uploadedFiles"

	entries, err := os.ReadDir(folderPath)
	log.Println(entries)
	info, err := os.Stat(folderPath)
	if err != nil{
		if info == nil{
		err = os.Mkdir("./uploadedFiles", 0755)
		if err != nil{
			fmt.Println("Error: ", err)
		}
	}
	}
	// cerate the file
	fullPath := folderPath + "/" + fileName
	copiedFile, err := os.Create(fullPath)
	if err != nil{
		log.Fatal("Error in creating the file. ", err)
	}

	// make a copy of the file inside the given folder made above
	_, err = io.Copy(copiedFile, file)
	if err != nil{
		log.Fatal("Error in copying the file.")
	}

	// send this file to the pipeline ->
	// i. Extract the text from the pdf
	// ii. Converting the extracted text to a single string
	// iii. Split this string
	// iv. Create embeddings and store in vector db

	err = extractpdf.ExtractData(&file, fullPath, h.server)
	fmt.Println("Error: ", err)
	if err != nil{
		fmt.Println("Error in extracting the pdf.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		// now add the error message
		var errResponse models.ErrorResponse
		utils.PrepareErrorMessage(err.Error())
		json.NewEncoder(w).Encode(&errResponse)
		return
	}
	// if no nil just return extraction and saving compelted
	var response models.SuccessResponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	// set the response
	response.Success = true
	response.SuccessMsg.Code = http.StatusAccepted
	response.SuccessMsg.Message = "Book embedded successfully."
	json.NewEncoder(w).Encode(&response)
}