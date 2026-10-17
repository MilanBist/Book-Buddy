package api

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

type StoreFileToLocation interface{
	StoreFile(fileName string, file multipart.File)(string, error)
}

type StoreRawFile interface{
	AddBookIdentity(bookTitle string, userId int, status string) (int, error)
	AddBookJob(bookId int, status, errMsg string, startedTime, completedTime *time.Time) (int, error)
	UpdateBookStatus(bookId int, status, errMsg string ) error
	UpdateBookJobStatus(jobId int, status, errMsg string) error
}

type Extract interface{
	ExtractData(file *multipart.File, filePath string) (int, error)
}

type RawPdfHandler struct{
	GetFileLocation 	StoreFileToLocation
	StoreBook 			StoreRawFile
	ExtractPdf 			Extract
}



func timePtr(now time.Time) *time.Time{
	return &now
}

func(h *RawPdfHandler) HandleRawPdfFile(w http.ResponseWriter, r *http.Request){
	var booksId int
	var jobId int
	var userId int

	file,header, err := r.FormFile("document")
	if err != nil{
		fmt.Println("[PDF HANDLER]: Error in obtaining the book.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.APIResponse{
			Success: false,
			Message: "Can't obtain the pdf file.",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	fileName := header.Filename

	// check for the extension explicity as the .pdf
	if strings.ToLower(filepath.Ext(fileName)) != ".pdf" {
		fmt.Println("[PDF HANDLER]: Only can extract the .pdf.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := models.APIResponse{
			Success: false,
			Message: "precise only .pdf is accepted.",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	fullPath, err := h.GetFileLocation.StoreFile(fileName, file)
	if err != nil{
		fmt.Println("[BOOK PROCESSING HANDLER]: Error is : ", err)
		w.WriteHeader(http.StatusInternalServerError)

		response := models.APIResponse{
			Success: false,
			Message: "Error in storing file.",
		}

		json.NewEncoder(w).Encode(&response)
	}

	userId = r.Context().Value("userId").(int)

	// add the book's identity as pending one
	booksId, err = h.StoreBook.AddBookIdentity(fileName, userId, "pending")
	if err != nil{
		fmt.Println("[PDF HANDLER]: Error in adding the book to db.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	// get the job id as well
	jobId, err = h.StoreBook.AddBookJob(booksId, "queued", "", timePtr(time.Now()), nil)
	if err != nil{
		fmt.Println("[PDF HANDLER]: Error in adding the job to db.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}


	go func(){
		defer file.Close()
		// set as the processing one
		_ = h.StoreBook.UpdateBookStatus(booksId, "processing", "")
		_ = h.StoreBook.UpdateBookJobStatus(jobId, "processing", "")
		_, err = h.ExtractPdf.ExtractData(&file, fullPath)
		if err != nil{
			_ = h.StoreBook.UpdateBookStatus(booksId, "rejected", err.Error())
			_ = h.StoreBook.UpdateBookJobStatus(jobId, "rejected", err.Error())
		}

		_ = h.StoreBook.UpdateBookStatus(booksId, "completed", err.Error())
		_ = h.StoreBook.UpdateBookJobStatus(jobId, "completed", err.Error())

	}()

	w.WriteHeader(http.StatusAccepted)
	response := models.APIResponse{
		Success: true,
		Message: "",
		Data: jobId,
	}
	json.NewEncoder(w).Encode(&response)
}