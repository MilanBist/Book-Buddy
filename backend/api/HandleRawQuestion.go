package api

import "net/http"

func(h *Handler) HandleRawQuestion(w http.ResponseWriter, r *http.Request){
	// get the data from the form and work upon it
	w.Write([]byte("Question to be answered."))
	
}