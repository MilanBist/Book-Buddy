package dbqueries


import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// iniitally just add the book title and later on also add the userId as well
func(p *PostgresStore) AddConversation(userQuery, response string, userId, bookId int) (int, error){
	// add all of this data to the conn pool
	ctx := context.Background()


	var id int = -1

	// NOW ADD THIS TO THE DATABASe using the pgx
	query := `INSERT INTO "conversation"("userQuestion", "chatResponse", "userId", "bookId")
			 VALUES ($1, $2, $3, $4)
			 RETURNING "id"
			 `

	err := p.Db.QueryRow(ctx, query, userQuery, response, userId, bookId).Scan(&id)
	if err != nil{
		// show the user that there is error
		fmt.Println("[DATABASE CONNECTION]: Error in querying")
		fmt.Println(err)
		return http.StatusInternalServerError, errors.New("Error in doing")
	}

	return http.StatusAccepted, nil
}