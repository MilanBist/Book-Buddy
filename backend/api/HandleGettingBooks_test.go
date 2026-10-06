package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

type FakeBookInformation struct {
	Books []models.BookData
	Error error
}

func (f *FakeBookInformation) GetAllBooksBasedOnUserId(userId int) ([]models.BookData, error) {
	return f.Books, f.Error
}

func TestHandleGettingBooks(t *testing.T) {
	store := &FakeBookInformation{
		Books: []models.BookData{},
		Error: nil,
	}

	handler := &BookGettingHandler{
		Book: store,
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/getBooks",
		nil,
	)

	ctx := context.WithValue(req.Context(), "userId", 1)
	req = req.WithContext(ctx)

	res := httptest.NewRecorder()

	handler.HandleGettingBooks(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Required 202 and got %d", res.Code)
	}

	var response models.APIResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if !response.Success {
		t.Fatal("Expected success to be true")
	}

	if response.Message != "Data successfully fetched." {
		t.Fatalf("Unexpected message: %s", response.Message)
	}
}