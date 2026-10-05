package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/MilanBist/AI-Powered-Book-Answerer/utils"
)

type LoginCheck interface{
	CheckUser(loginCredentials models.Login)(int, error)
}
// inherit token generator from the register too
type LoginHandler struct{
	Check 	LoginCheck
	Token 	TokenGenerator
}



func (h *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request){

	var loginCredentials models.Login
	json.NewDecoder(r.Body).Decode(&loginCredentials)
	isValid, errMsg := utils.ValidateUserLogin(&loginCredentials)
	if !isValid{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		response := models.APIResponse{
			Success: false,
			Message:"wrong credentials",
			Data: errMsg,
		}		
		json.NewEncoder(w).Encode(&response)
		return
	}


	fmt.Println(loginCredentials)
	// check for the user and get the id
	userId, err := h.Check.CheckUser(loginCredentials)
	if err != nil{
		// just return the error
		// get in certain format for sending to the frontend
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		response := models.APIResponse{
			Success: false,
			Message:"no user",
		}
		fmt.Println("Response might be: ", response)
		json.NewEncoder(w).Encode(&response)
		return
	}

	fmt.Println("Reaching here.");
	// if the isRegistered is true then do a thing like send with the jwt token in it
	tokenString, err := h.Token.GenerateTokens(int64(userId), loginCredentials.Email)
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
		Message: "Logged in Successfully.",
		Data: models.AuthData{
			Token: tokenString,
		},
	}

	fmt.Println("[LOGIN HANDLER] Reponse data: ", response)
	json.NewEncoder(w).Encode(&response)

}