package lib

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Middleware(next http.Handler) http.Handler{
	err := godotenv.Load()
    if err != nil {
		log.Println("No .env file found, reading from environment variables")
    }
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" || apiKey != os.Getenv("API_KEY"){
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w,r)
	})
}