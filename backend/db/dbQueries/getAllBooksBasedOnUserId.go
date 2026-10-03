package dbqueries

import (
	"context"
	"fmt"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

func(p *PostgresStore) GetAllBooksBasedOnUserId(userId int) ([]models.BookData, error){
	// get the books based on the id of the user
	query := `SELECT "id", "bookName" FROM "books" WHERE "user_id" = $1 ORDER BY "uploadedDate" DESC`

	rows, err := p.Db.Query(context.Background(), query, userId)

	if err != nil{
		// just return the error
		fmt.Println("[DATABASE ERROR]: Error in querying the data.")
		return []models.BookData{}, err
	}

	defer rows.Close()
	var Books []models.BookData

	for rows.Next(){
		var book models.BookData

		// make a struct of the book
		err := rows.Scan(
			&book.BookId,
			&book.BookName,
		)

		if err != nil{
			fmt.Println("[DATABASE BOOKS FETCH]: Error in fetching the book")
			return []models.BookData{}, err
		}
		Books = append(Books, book)
	}
	return Books, nil
}