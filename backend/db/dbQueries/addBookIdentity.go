package dbqueries

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

// iniitally just add the book title and later on also add the userId as well
func AddBookIdentity(bookTitle string, bookId int, db *pgxpool.Pool) (int, error){
	// add all of this data to the conn pool
	ctx := context.Background()


	// // check before adding to the database if the book's name already exists
	// // check for the name of the book added

	// query := `SELECT "id" FROM "books" 
	// 		WHERE "bookName" = $1
	// 		`

	var id int = -1
	// err := db.QueryRow(ctx, query, bookTitle).Scan(&id)
	// if err != nil{
	// 	if errors.Is(err, pgx.ErrNoRows){
	// 	// 1. Handle the "no data found" case safely
	// 	fmt.Println("[DATABASE CONNECTION] no data found.")
	// 	} else{
	// 	// show the user that there is error
	// 	fmt.Println("[DATABASE CONNECTION]: Error in querying")
	// 	return -1, err
	// 	}		
	// }


	// // if the id is not null
	// if id != -1{
	// 	return id, nil
	// }

	// NOW ADD THIS TO THE DATABASe using the pgx
	query := `INSERT INTO "books"("bookName", "userId"")
			 VALUES ($1, $2)
			 RETURNING "id"
			 `

	err := db.QueryRow(ctx, query).Scan(&id)
	if err != nil{
		// show the user that there is error
		fmt.Println("[DATABASE CONNECTION]: Error in querying")
		return http.StatusInternalServerError, errors.New("Error in doing")
	}

	return http.StatusAccepted, nil
}