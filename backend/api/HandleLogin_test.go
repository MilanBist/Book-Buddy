
package api

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "github.com/MilanBist/AI-Powered-Book-Answerer/internal/models"
)

type FakeLoginStore struct {
    UserId int
    Error  error
}

func (f *FakeLoginStore) CheckUser(loginCredentials models.Login) (int, error) {
    return f.UserId, f.Error
}

func TestHandleLogin(t *testing.T) {
	// fake token generator from the register handler
    token := &FakeTokenGenerator{
        Token: "milan-123",
        Error: nil,
    }

    store := &FakeLoginStore{
        UserId: 1,
        Error:  nil,
    }

    fakeLoginSrv := &LoginHandler{
        Check: store,
        Token: token,
    }

    body := `{
        "userEmail": "bistmilan46@gmail.com",
        "userPassword": "Milbis123@#"
    }`

    req := httptest.NewRequest(
        http.MethodPost,
        "/api/login",
        strings.NewReader(body),
    )

    res := httptest.NewRecorder()

    fakeLoginSrv.HandleLogin(res, req)

    if res.Code != http.StatusOK {
        t.Fatalf("Required 200 and got %d", res.Code)
    }

	// almost similar format as register response
    var loginRes ResgisterResponse

    if err := json.NewDecoder(res.Body).Decode(&loginRes); err != nil {
        t.Fatalf("Failed to decode response: %v", err)
    }

    if !loginRes.Success {
        t.Fatal("Expected login success to be true")
    }

    if loginRes.Data.Token != "milan-123" {
        t.Fatalf(
            "Required milan-123 and got: %s",
            loginRes.Data.Token,
        )
    }
}