package dbqueries

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct{
	Db *pgxpool.Pool
}

// iniitally just add the book title and later on also add the userId as well
func(p *PostgresStore) AddBookIdentity(bookTitle string, userId int) (int,int, error){
	// add all of this data to the conn pool
	ctx := context.Background()

	fmt.Println("[DATABASE CONNECTION]: User id: ", userId)

	var id int = -1


	// NOW ADD THIS TO THE DATABASe using the pgx
	query := `INSERT INTO "books"("bookName", "user_id")
			 VALUES ($1, $2)
			 RETURNING "id"
			 `

	err := p.Db.QueryRow(ctx, query, bookTitle, userId).Scan(&id)
	if err != nil{
		// show the user that there is error
		fmt.Println("[DATABASE CONNECTION]: Error in querying")
		fmt.Println("Actual Error: ", err)
		return http.StatusInternalServerError, -1, errors.New("Error in doing")
	}

	return http.StatusAccepted,id, nil
}