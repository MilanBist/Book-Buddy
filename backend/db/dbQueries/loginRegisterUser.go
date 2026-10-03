package dbqueries

import (
	"context"
	"errors"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)


func(p *PostgresStore) RegisterUser(credentials models.Register) (int, error){
	ctx := context.Background()

	// NOW ADD THIS TO THE DATABASe using the pgx
	query := `INSERT INTO "users"("firstName", "lastName", "address", "email", "password")
			 VALUES ($1, $2, $3, $4, $5)
			 RETURNING "id"
			 `
	var id int

	// add this query to the database
	// hash the password before using it
	err := p.Db.QueryRow(ctx, query, credentials.FirstName, credentials.LastName, credentials.Address, credentials.Email, credentials.Password).Scan(&id)
	if err != nil{
		// show the user that there is error
		return -1, err
	}

	return id, nil
}

func(p *PostgresStore) UserExistence(email string) (bool){
	// if the user exist return true and if not return false
	ctx := context.Background()
	query := `
			SELECT "id" FROM "users"
			WHERE "email" = $1
		`

	var id int
	id = -1

	// now get the id
	p.Db.QueryRow(ctx, query, email).Scan(&id)
	if id == -1{
		return false
	}

	// user exist
	return true
}

func(p *PostgresStore) CheckUser(credentials models.Login) (int, error){
	ctx := context.Background()
	query := `
				SELECT "id" FROM "users" 
				WHERE "email" = $1
				AND "password" = $2
			 `
	var id int
	id = -1
	// hash the password before using it
	err := p.Db.QueryRow(ctx, query, credentials.Email, credentials.Password).Scan(&id)
	if err != nil{
		// show the user that there is error
		return -1, err
	}
	if id == -1{
		// return false
		return -1, errors.New("No user found")
	}

	return id, nil
}