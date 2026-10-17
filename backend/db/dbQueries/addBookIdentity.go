package dbqueries

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct{
	Db *pgxpool.Pool
}

// function to add the job of the book
func(p *PostgresStore) AddBookIdentity(bookTitle string, userId int, status string) (int, error){
	ctx := context.Background()
	var id int = -1
	query := `INSERT INTO "books"("bookName", "userId", "status")
			 VALUES ($1, $2, $3)
			 RETURNING "id"
			 `

	err := p.Db.QueryRow(ctx, query, bookTitle, userId, status).Scan(&id)
	if err != nil{
		fmt.Println("[DATABASE CONNECTION]: Error in querying: ", err)
		return -1, errors.New("Error in doing")
	}
	return id, nil
}

// function to add the job of the book
func (p *PostgresStore) AddBookJob(bookId int, status, errMsg string, startedTime, completedTime *time.Time) (int, error){
	ctx := context.Background()

	var id int = -1

	query := `INSERT INTO "bookJob"("bookId", "status", "errorMessage", "startedAt", "completedAt")
			 VALUES ($1, $2, $3, $4, $5)
			 RETURNING "id"
			 `

	err := p.Db.QueryRow(ctx, query, bookId, status, errMsg, startedTime, completedTime).Scan(&id)
	if err != nil{
		fmt.Println("[DATABASE CONNECTION]: Error in querying: ", err)
		return -1, err
	}

	return id, nil
}



func (p *PostgresStore) UpdateBookStatus(bookId int, status, errMsg string) error{
	ctx := context.Background()

	if status == ""{
		status = "completed"
	}

	if status != "pending" && status != "completed"{
		return errors.New("Status: pending or completed.")
	}


	query := `
		UPDATE "books"
		SET "status" = $1
		WHERE "id" = $2
	`

	_, err := p.Db.Exec(ctx, query, status, bookId)
	if err != nil {
		return fmt.Errorf("failed to update book status: %w", err)
	}
	return nil
}

func (p *PostgresStore) UpdateBookJobStatus(jobId int, status, errMsg string) error{
	ctx := context.Background()

	if status ==""{
		status = "completed"
	}

	if status != "pending" && status != "completed"{
		return errors.New("Status: pending or completed.")
	}

	query := `
		UPDATE "bookJob"
		SET "status" = $1
		WHERE "id" = $2
	`

	_, err := p.Db.Exec(ctx, query, status, jobId)
	if err != nil {
		return fmt.Errorf("failed to update book status: %w", err)
	}

	return nil

}