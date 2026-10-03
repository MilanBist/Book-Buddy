package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

type FakeUserConversation struct {
	Conversations []models.ReturningConversation
	Error         error
}

func (f *FakeUserConversation) GetAllChats(userId, bookId int) ([]models.ReturningConversation, error) {
	return f.Conversations, f.Error
}

func TestHandleConversation(t *testing.T) {
	store := &FakeUserConversation{
		Conversations: []models.ReturningConversation{},
		Error:         nil,
	}

	handler := &UserConversationHandler{
		Conversation: store,
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/conversation?bookId=1&bookName=MyBook",
		nil,
	)

	ctx := context.WithValue(req.Context(), "userId", 1)
	req = req.WithContext(ctx)

	res := httptest.NewRecorder()

	handler.HandleConversation(res, req)

	if res.Code != http.StatusAccepted {
		t.Fatalf("Required 202 and got %d", res.Code)
	}

	var response models.APIResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if !response.Success {
		t.Fatal("Expected success to be true")
	}

	if response.Message != "Conversation Data fetched successfully." {
		t.Fatalf("Unexpected message: %s", response.Message)
	}
}