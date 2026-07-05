package api

import (
	"encoding/json"
	"net/http"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/controllers"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)




func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request){

	var loginCredentials *models.Login

	json.NewDecoder(r.Body).Decode(&loginCredentials)

	// send all of the data to check and do things with the login 
	controllers.Login(loginCredentials, h.server.SqliteDB)


}