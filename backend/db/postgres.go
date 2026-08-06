package db

//  make a connection with the postgres server
import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func PostgresConnection() (*pgxpool.Pool, error){
	// create the connection with the postgres server
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := fmt.Sprintf(
	"postgres://%s:%s@%s:%s/%s?sslmode=disable",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"),
)

	dbpool, dberr := pgxpool.New(context.Background(), url)
	if dberr != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, errors.New("Error in creating the connection to the postgres.")
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
    	log.Fatal("Database is not reachable:", err)
	}

fmt.Println("Successfully connected to PostgreSQL!")

	// get the table information
	rows, err := dbpool.Query(
    	context.Background(),
		`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
		ORDER BY table_name;
		`,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Tables:")

	for rows.Next() {
		var table string
		rows.Scan(&table)
		fmt.Println("-", table)
	}

	return dbpool, nil
}