package controllers

import (
	"errors"
	"fmt"
	"net/http"

	dbqueries "github.com/MilanBist/AI-Powered-Book-Answerer/db/dbQueries"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/MilanBist/AI-Powered-Book-Answerer/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)


func Register(credentials models.Register, db *pgxpool.Pool) (int,int, error){
	// validate all of the credential data
	isValid, errMsg := utils.ValidateUserRegister(credentials)

	if isValid == false{
		// There is some error print that and return 
		fmt.Println(errMsg)
		return http.StatusBadRequest,-1, errors.New(errMsg)
	}
	fmt.Println("[REGISTER VALIDATION] For the validation: ", isValid)

	// check if the email person exist or not
	userExist := dbqueries.UserExistence(credentials.Email, db)
	if userExist == true{
		// user exist so register can be done
		return http.StatusBadRequest, -1, errors.New("User exists.") 
	}

	// add the user to the table in the db
	id, err := dbqueries.RegisterUser(credentials, db)

	if err != nil{
	// show the user that there is error
		fmt.Println("[REGISTER]: Error in executing the command. Acutal error: ", err)
		return http.StatusBadRequest, -1, err
	}
	
	
	fmt.Println("[REGISTER]: Successfully registered user.")

	return http.StatusAccepted,id, nil	
}