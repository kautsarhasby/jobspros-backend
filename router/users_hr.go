package router

import (
	"backend-fullstack/lib"
	"backend-fullstack/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HrRequest struct {
	Id *int `json:"id" db:"id"`
	HrId   int   `json:"hr_id" db:"hr_id"`
    CompanyId     int `json:"company_id" db:"company_id"`
}


// Get all users : /users
func GetHR(w http.ResponseWriter, r *http.Request){
	var hrs []models.HR

	db,err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Select(&hrs,"SELECT users_hr.id,users.uuid as hr_uuid,users.name as hr_name, users.email as hr_email, companies.name as company_name, companies.industry as company_industry, companies.uuid as comp_uuid FROM users_hr JOIN users ON users.id = users_hr.hr_id JOIN companies ON companies.id = users_hr.company_id WHERE users.deleted_at is null")
	if err != nil{
		fmt.Println(err)
		return
	}
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(hrs)
}

func GetHRById(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	db,err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()
	var hr models.HR
	err = db.Get(&hr,"SELECT users_hr.id,users.uuid as hr_uuid,users.name as hr_name, users.email as hr_email, companies.name as company_name, companies.industry as company_industry, companies.uuid as comp_uuid FROM users_hr JOIN users ON users.id = users_hr.hr_id JOIN companies ON companies.id = users_hr.company_id WHERE users_hr.id = $1",id)
	if err != nil{
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(hr)
}

func PostHR(w http.ResponseWriter, r *http.Request){
	var hr HrRequest
	
	db,err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	byteValue,_ := io.ReadAll(r.Body)
	json.Unmarshal(byteValue,&hr)

	_,err = db.NamedExec(`INSERT INTO users_hr (hr_id,company_id) VALUES (:hr_id,:company_id)`,hr)
	if err != nil{
		fmt.Println("Insert error:", err)
		http.Error(w,"Something went error",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status":"Success Added HR"})
}