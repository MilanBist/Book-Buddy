package utils

import "testing"

func TestGenerateJwtToken(t *testing.T){
	token := &Token{
		SecretKey: []byte("d9702d28601fe4aa671e65e7c28d8f87e69f8807cc1c6b52c11376390aed383d"),
	}

	email := "bistmialn46@gmail.com"
	userId := int64(1)

	tokenString, err := token.GenerateTokens(userId, email)
	if err != nil{
		t.Fatalf("expected no error, got %v", err)
	}

	if tokenString == ""{
		t.Fatal("Expected to get token string to be some value.")
	}
}