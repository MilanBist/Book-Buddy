package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/controllers"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/MilanBist/AI-Powered-Book-Answerer/utils"
)


func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request){

	var registerCredentials models.Register

	json.NewDecoder(r.Body).Decode(&registerCredentials)

	fmt.Println(registerCredentials)

	// send all of the data to check and do things with the login 
	statuCode, userId, err := controllers.Register(registerCredentials, h.server.PostgresDB)


	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statuCode)
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	// if the isRegistered is true then do a thing like send with the jwt token in it
	tokenString, err := utils.GenerateTokens(int64(userId), registerCredentials.Email)
	
	if err != nil{
		fmt.Println("[REGISTER ENDPOINT]: Error in generating the token.")
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
	fmt.Println("[REGISTER HANDLER] Reponse data: ", response)
	json.NewEncoder(w).Encode(&response)

}