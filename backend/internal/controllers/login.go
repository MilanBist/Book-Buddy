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


func Login(credentials *models.Login, db *pgxpool.Pool) (int, int, error){
	// first task is check if the given login data is present in my
	// database

	isValid, errMsg := utils.ValidateUserLogin(credentials)

	if !isValid{
		fmt.Println(errMsg)
		return http.StatusBadRequest, -1, errors.New(errMsg)
	}


	// check for the user and get the id
	id, err := dbqueries.CheckUser(*credentials, db)
	if err != nil{
		fmt.Println("[LOGIN]: Error: ", err)
		return http.StatusBadRequest, -1, err
	}

	fmt.Println("[LOGIN] User found with simiar credentials.")

	// return the id if there is 
	return http.StatusAccepted, id, nil
}