package middlewares

import (
	"fmt"
	"net/http"
	"time"
)


func LoggingMiddleware(next http.Handler) http.Handler{
	// return the new handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		// Get the header and other thing
		fmt.Println("====================================")
		fmt.Println("Request Method: ", r.Method)
		fmt.Println("Url Path: ", r.URL.Path)
		fmt.Println("IP address: ", r.RemoteAddr)

		// serve the next request
		next.ServeHTTP(w,r)

		fmt.Println("Duration time: ", time.Since(start))
		fmt.Println("====================================")

	})
}