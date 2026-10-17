package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

var BookId int64

// interface to check for the status
type StatusChecker interface{
	CheckStatus(bookId int64) (string, error)
}


type HandleProcess struct{
	ProcessCheck 	StatusChecker
}

// handle the processing status
func(s *HandleProcess) HandleBookProcessingStatus(w http.ResponseWriter, r *http.Request){
	json.NewDecoder(r.Body).Decode(&BookId)

	status, err := s.ProcessCheck.CheckStatus(BookId)


	if err != nil{
		fmt.Println("[BOOK STATUS CHECKER ENDPOINT]: Error in checking the status.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	switch status{
	case "PENDING":
		w.WriteHeader(http.StatusAccepted)
	case "ACCEPTED":
		w.WriteHeader(http.StatusOK)
	}


	response := models.APIResponse{
		Success: true,
		Message: "Successfull in fetching the data.",
		Data: status,
	}
	json.NewEncoder(w).Encode(&response)

}