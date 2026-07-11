package api

import (
	"encoding/json"
	"net/http"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/controllers"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)


func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request){

	var registerCredentials *models.Register

	json.NewDecoder(r.Body).Decode(&registerCredentials)

	// send all of the data to check and do things with the login 
	controllers.Register(registerCredentials, h.server.SqliteDB)


}