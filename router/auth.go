package router

import (
	"backend-fullstack/lib"
	"backend-fullstack/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AuthRequest struct {
	Email    string `json:"email" db:"email"`
    Password string `json:"password" db:"password"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
		case http.MethodPost:
			var authUser AuthRequest

			db, err := lib.Connection();
			if err != nil{
				fmt.Println(err)
				return
			}

			defer db.Close()

			byteValue,_ := io.ReadAll(r.Body)
			json.Unmarshal(byteValue,&authUser)

			var result models.User
			
			err = db.Get(&result,"SELECT * FROM users WHERE email = ?",authUser.Email)
			if err != nil{
				http.Error(w,"Email and Password are incorrect",http.StatusUnauthorized)
				return
			}

			credentials := lib.CheckPasswordHash(authUser.Password,result.Password)
			if (!credentials){
				http.Error(w,"Password are incorrect",http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"status":"success","message":"Login Success"})
		default:
			http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
	}
	
}

