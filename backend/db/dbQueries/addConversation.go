package dbqueries


import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

// iniitally just add the book title and later on also add the userId as well
func AddConversation(userQuery, response string, userId, bookId int, db *pgxpool.Pool) (int, error){
	// add all of this data to the conn pool
	ctx := context.Background()


	var id int = -1

	// NOW ADD THIS TO THE DATABASe using the pgx
	query := `INSERT INTO "conversation"("userQuestion", "chatResponse", "userId", "bookId")
			 VALUES ($1, $2, $3, $4)
			 RETURNING "id"
			 `

	err := db.QueryRow(ctx, query, userQuery, response, userId, bookId).Scan(&id)
	if err != nil{
		// show the user that there is error
		fmt.Println("[DATABASE CONNECTION]: Error in querying")
		return http.StatusInternalServerError, errors.New("Error in doing")
	}

	return http.StatusAccepted, nil
}