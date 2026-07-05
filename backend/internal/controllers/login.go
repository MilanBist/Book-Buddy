package controllers

import (
	"database/sql"
	"fmt"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/MilanBist/AI-Powered-Book-Answerer/utils"
)


func Login(credentials *models.Login, db *sql.DB) (){
	// first task is check if the given login data is present in my
	// database

	isValid, errMsg := utils.ValidateUserLogin(credentials)

	if !isValid{
		fmt.Println(errMsg)
		return
	}

	
	
}