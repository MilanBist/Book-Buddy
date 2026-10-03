package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

type FakeRawQuestionStore struct {
	Chats         []models.Messages
	GetChatsError error
	AddConvError  error
}

func (f *FakeRawQuestionStore) GetTop5Chats(userId int, bookId int) ([]models.Messages, error) {
	return f.Chats, f.GetChatsError
}

func (f *FakeRawQuestionStore) AddConversation(userQuery, response string, userId, bookId int) (int, error) {
	return 1, f.AddConvError
}

type FakeRawQuestionEmbeddings struct {
	Embedding     []float32
	GenerateError error
	Documents     []string
	FindError     error
}

func (f *FakeRawQuestionEmbeddings) GenerateEmebedding(splittedDocx string) ([]float32, error) {
	return f.Embedding, f.GenerateError
}

func (f *FakeRawQuestionEmbeddings) FindBestEmbeddings(point []float32, bookName string) ([]string, error) {
	return f.Documents, f.FindError
}

type FakeQuestionResponse struct {
	Response string
	Error    error
}

func (f *FakeQuestionResponse) AskLLM(docx []string, preChats []models.Messages, query string, language string, w http.ResponseWriter, r *http.Request, userId int, bookId int) (string, error) {
	return f.Response, f.Error
}

func TestHandleRawQuestion(t *testing.T) {
	store := &FakeRawQuestionStore{
		Chats:         []models.Messages{},
		GetChatsError: nil,
		AddConvError:  nil,
	}

	embeddings := &FakeRawQuestionEmbeddings{
		Embedding:     []float32{0.1, 0.2, 0.3},
		GenerateError: nil,
		Documents: []string{
			"This is the first relevant document.",
			"This is the second relevant document.",
		},
		FindError: nil,
	}

	response := &FakeQuestionResponse{
		Response: "This is the answer.",
		Error:    nil,
	}

	handler := &QuestionHanlder{
		StoreQuestion:  store,
		AnswerQuestion: embeddings,
		Response:       response,
	}

	body := `{
		"query": "What is this book about?",
		"bookName": "My Book",
		"bookId": 1,
		"language": "English"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/extractAnswer", strings.NewReader(body))

	ctx := context.WithValue(req.Context(), "userId", 1)
	req = req.WithContext(ctx)

	res := httptest.NewRecorder()

	handler.HandleRawQuestion(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("Required 200 and got %d", res.Code)
	}
}