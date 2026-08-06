package middlewares

import (
	"context"
	"net/http"
	"strings"
	"github.com/MilanBist/AI-Powered-Book-Answerer/utils"
)


func AuthMiddleware(next http.Handler) http.Handler{
	// return the new handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		stringedToken := r.Header.Get("Authorization")
		if stringedToken == ""{
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
    		return
		}
		splitedToken := strings.Split(stringedToken, " ")


		if len(splitedToken) != 2 || splitedToken[0] != "Bearer" {
    		http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
    		return
		}

		token := splitedToken[len(splitedToken)-1]

		// now validate the token 
		userId, err := utils.ValidateToken(token)
		if err != nil{
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// create a context -> travelling request
		ctx := context.WithValue(r.Context(), "userId", userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w,r)
	})
}