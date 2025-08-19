package router

import (
	"backend-fullstack/lib"
	"backend-fullstack/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)


type UserRequest struct {
	Id *int `json:"id" db:"id"`
	Uuid      string  `json:"uuid" db:"uuid"`
    Name     string `json:"name" db:"name"`
    Email    string `json:"email" db:"email"`
    Password string `json:"password" db:"password"`
    Role     string `json:"role" db:"role"`
}


// Get all users : /users
func GetUsers(w http.ResponseWriter, r *http.Request){
	db,err := lib.Connection()
	var users []models.User

	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()


	err = db.Select(&users,"SELECT * FROM users WHERE deleted_at is null")
	if err != nil{
		fmt.Println(err)
		return
	}
	

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(users)
	
}

// Get user by id : /users/2
func GetUserById(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	db,err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()
	var user models.User
	err = db.Get(&user,"SELECT * FROM users WHERE id = ?",id)
	if err != nil{
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(user)
}

// Adding new user : /users
func PostUser(w http.ResponseWriter, r *http.Request){
	var user UserRequest
	uuid := uuid.New()
	db,err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	byteValue,_ := io.ReadAll(r.Body)
	json.Unmarshal(byteValue,&user)
	user.Uuid = uuid.String()
	user.Password = lib.HashedPassword(user.Password)

	_,err = db.NamedExec(`INSERT INTO users (uuid,name,email,password,role) VALUES (:uuid,:name,:email,:password,:role)`,user)
	if err != nil{
		fmt.Println("Insert error:", err)
		http.Error(w,"Something went error",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status":"Success Added User"})
}

// Update user : /users/3
func UpdateUser(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	db, err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	byteValue,_ := io.ReadAll(r.Body)

	var user UserRequest
	json.Unmarshal(byteValue,&user)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	user.Id = &idInt

	_,err = db.NamedExec(`UPDATE users SET name = :name, email = :email, password = :password, role = :role WHERE id = :id`,user)
	if err != nil{
		fmt.Println("Update error:", err)
		http.Error(w,"Something went error",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status":"success","message":"Success added user"})

}

// Delete user : /users/3
func DeleteUser(w http.ResponseWriter, r *http.Request, id string){
	db, err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE users set deleted_at = NOW() WHERE id = ?",id)
	if err != nil{
		fmt.Println("Delete error:", err)
		http.Error(w,"User not found",http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status":"Success Delete User"})
}