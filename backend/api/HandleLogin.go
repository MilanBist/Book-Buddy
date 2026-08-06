package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/controllers"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/MilanBist/AI-Powered-Book-Answerer/utils"
)




func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request){

	var loginCredentials *models.Login

	json.NewDecoder(r.Body).Decode(&loginCredentials)

	// send all of the data to check and do things with the login 
	statusCode,userId, err := controllers.Login(loginCredentials, h.server.PostgresDB)
	fmt.Println(userId)


	if err != nil{
		// just return the error
		// get in certain format for sending to the frontend
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	// if the isRegistered is true then do a thing like send with the jwt token in it
	tokenString, err := utils.GenerateTokens(int64(userId), loginCredentials.Email)
	if err != nil{
		fmt.Println("[LOGIN ENDPOINT]: Error in generating the token.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}


	// return the true messaage to the frontend with the token string
	response := models.APIResponse{
		Success: true,
		Message: "Registered Successfully.",
		Data: models.AuthData{
			Token: tokenString,
		},
	}


	fmt.Println("[LOGIN HANDLER] Reponse data: ", response)
	json.NewEncoder(w).Encode(&response)

}