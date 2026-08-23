package dbqueries

import (
	"context"
	"fmt"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetTop5Chats(userId, bookId int, db *pgxpool.Pool)(error, []models.Messages){
	query := `SELECT "userQuestion", "chatResponse" FROM "conversation" WHERE
				"userId" = $1 AND "bookId" = $2
				ORDER BY "conversationTime" LIMIT 5;
			`
	rows, err := db.Query(context.Background(), query, userId, bookId)
	if err != nil{
		fmt.Println("[DATABASE QUERY ERROR]: Error in inserting.")
		return err, nil
	}

	defer rows.Close()

	var conversation []models.Messages

	for rows.Next(){
		var conv models.Messages

		err := rows.Scan(&conv.UserQuestion, &conv.LLMResponse)
		if err != nil{
			fmt.Println("[DATABASE QUERY ERROR]: Error in reading the data.", err)
			return err, nil
		}

		conversation = append(conversation, conv)
	}

	return  nil, conversation

}


func GetAllChats(userId, bookId int, db *pgxpool.Pool)(error, []models.ReturningConversation){
	query := `SELECT "userQuestion", "chatResponse" FROM "conversation" WHERE
				"userId" = $1 AND "bookId" = $2
				ORDER BY "conversationTime" ASC;
			`
	rows, err := db.Query(context.Background(), query, userId, bookId)
	if err != nil{
		fmt.Println("[DATABASE QUERY ERROR]: Error in inserting.")
		return err, nil
	}

	defer rows.Close()

	var conversation []models.ReturningConversation

	for rows.Next(){
		var user models.ReturningConversation
		var chat models.ReturningConversation

		user.Role = "user"
		chat.Role = "assistant"

		err := rows.Scan(&user.Message, &chat.Message)
		if err != nil{
			fmt.Println("[DATABASE QUERY ERROR]: Error in reading the data.", err)
			return err, nil
		}

		conversation = append(conversation, user)
		conversation = append(conversation, chat)
	}
	return  nil, conversation
}