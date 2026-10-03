package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

type FakeRegisterStore struct{
	UserExist bool
	UserId int
	Error error
}

type FakeTokenGenerator struct{
	Token 		string
	Error 		error
}


func (f *FakeRegisterStore) UserExistence(email string) bool{
	return f.UserExist
}

func (f *FakeRegisterStore) RegisterUser(registerCredentials models.Register) (int, error){
	return f.UserId, f.Error
}
func (f *FakeTokenGenerator) GenerateTokens(userId int64, email string) (string, error){
	return f.Token, f.Error
}

type ResgisterResponse struct{
	Success bool        	`json:"success"`
    Message string      	`json:"message"`
    Data   models.AuthData `json:"data"`
}


func TestHandleRegister( t *testing.T){
	token := &FakeTokenGenerator{
		Token: "milan-123",
		Error: nil,
	}
	store := &FakeRegisterStore{
		UserExist: false,
		UserId: 1,
		Error: nil,
	}

	fakeRegSrv := &RegisterHandler{
		Register: store,
		Token: token,
	}

	body := `{
		"userFirstName": "Milan",
		"userLastName": "Bist",
		"userAddress":"Mahakali-07, Darchula",
		"userEmail": "bistmilan46@gmail.com",
		"userPassword": "Milbis123@#"
	}`


	req := httptest.NewRequest(http.MethodPost, "/api/register", strings.NewReader(body))
	res := httptest.NewRecorder()
	fakeRegSrv.HandleRegister(res, req)


	if res.Code != 200 {
		t.Fatalf("Required 200 and got %d", res.Code)
	}

	var registerRes ResgisterResponse
	json.NewDecoder(res.Body).Decode(&registerRes)

	if registerRes.Data.Token != "milan-123"{
		t.Fatalf("Required milan-123 and got: %s", registerRes.Data.Token)
	}
}