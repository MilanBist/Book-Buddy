package dbqueries

import (
	"context"
	"fmt"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

func(p *PostgresStore) GetTop5Chats(userId, bookId int)([]models.Messages, error){
	query := `SELECT "userQuestion", "chatResponse" FROM "conversation" WHERE
				"userId" = $1 AND "bookId" = $2
				ORDER BY "conversationTime" LIMIT 5;
			`
	rows, err := p.Db.Query(context.Background(), query, userId, bookId)
	if err != nil{
		fmt.Println("[DATABASE QUERY ERROR]: Error in inserting.")
		return nil,err
	}

	defer rows.Close()

	var conversation []models.Messages

	for rows.Next(){
		var conv models.Messages

		err := rows.Scan(&conv.UserQuestion, &conv.LLMResponse)
		if err != nil{
			fmt.Println("[DATABASE QUERY ERROR]: Error in reading the data.", err)
			return nil, err
		}

		conversation = append(conversation, conv)
	}
	return  conversation, nil
}


func(p *PostgresStore) GetAllChats(userId, bookId int)([]models.ReturningConversation, error){
	query := `SELECT "userQuestion", "chatResponse" FROM "conversation" WHERE
				"userId" = $1 AND "bookId" = $2
				ORDER BY "conversationTime" ASC;
			`
	rows, err := p.Db.Query(context.Background(), query, userId, bookId)
	if err != nil{
		fmt.Println("[DATABASE QUERY ERROR]: Error in inserting.")
		return nil, err
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
			return nil, err
		}

		conversation = append(conversation, user)
		conversation = append(conversation, chat)
	}
	return  conversation, nil
}