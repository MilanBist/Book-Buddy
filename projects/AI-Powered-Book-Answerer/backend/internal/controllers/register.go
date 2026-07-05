package controllers

import (
	"database/sql"
	"fmt"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/MilanBist/AI-Powered-Book-Answerer/utils"
)


func Register(credentials *models.Register, db *sql.DB) {
	// validate all of the credential data

	isValid, errMsg := utils.ValidateUserRegister(credentials)

	if !isValid{
		fmt.Println(errMsg)
		return
	}

	fmt.Println(isValid)
}