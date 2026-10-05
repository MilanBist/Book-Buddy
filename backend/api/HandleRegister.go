package api


import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/MilanBist/AI-Powered-Book-Answerer/utils"
)

type RegisterStore interface{
	UserExistence(email string) (bool)
	RegisterUser(registerCredentials models.Register)(int, error)
}

type TokenGenerator interface{
	GenerateTokens(userId int64, email string) (string, error)
}
type RegisterHandler struct{
	Register 	RegisterStore
	Token 		TokenGenerator
}

func (h *RegisterHandler) HandleRegister(w http.ResponseWriter, r *http.Request){
	var registerCredentials models.Register
	json.NewDecoder(r.Body).Decode(&registerCredentials)

	isValid, msg := utils.ValidateUserRegister(registerCredentials)

	if isValid != true{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		response := models.APIResponse{
			Success: false,
			Message: "wrong credentials.",
			Data: msg,
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	// check if the email person exist or not
	userExist := h.Register.UserExistence(registerCredentials.Email)
	if userExist == true{
		// user exist so register can be done
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		response := models.APIResponse{
			Success: false,
			Message: "User already exists.",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	// add the user to the table in the db
	userId, err := h.Register.RegisterUser(registerCredentials)
	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: "can't insert.",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	// if the isRegistered is true then do a thing like send with the jwt token in it
	tokenString, err := h.Token.GenerateTokens(int64(userId), registerCredentials.Email)
	if err != nil{
		fmt.Println("[REGISTER ENDPOINT]: Error in generating the token.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := models.APIResponse{
			Success: false,
			Message: "can't generate",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	response := models.APIResponse{
		Success: true,
		Message: "Registered Successfully.",
		Data: models.AuthData{
			Token: tokenString,
		},
	}
	json.NewEncoder(w).Encode(&response)

}