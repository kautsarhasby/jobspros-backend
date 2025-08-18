package router

import (
	"backend-fullstack/lib"
	"backend-fullstack/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)


type CompanyRequest struct {
	Uuid      string  `json:"uuid" db:"uuid"`
	Name      string  `json:"name" db:"name"`
	Email     string  `json:"email" db:"email"`
	LogoUrl  *string  `json:"logo_url" db:"logo_url"`
	Address      string  `json:"address" db:"address"`
	City      string  `json:"city" db:"city"`
	Country      string  `json:"country" db:"country"`
	Description      *string  `json:"description" db:"description"`
	Phone      *string  `json:"phone" db:"phone"`
	Industry      string  `json:"industry" db:"industry"`
}





func GetCompanies(w http.ResponseWriter, r * http.Request){
	var companies []models.Company
	db, err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Select(&companies,"SELECT * FROM companies WHERE deleted_at IS NULL",)
	if err != nil{
		fmt.Println(err)
		return
	}


	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(companies)
}

func GetCompanyById(w http.ResponseWriter, r * http.Request){
	id := chi.URLParam(r,"id")
	var company models.Company
	db,err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Get(&company,"SELECT * FROM companies WHERE id = ?",id)
	if err != nil{
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(company)
}


// Adding new company : /companies
func PostCompany(w http.ResponseWriter, r *http.Request){
	var company CompanyRequest
	uuid := uuid.New()
	db,err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	fmt.Println("company : 	",company)
	byteValue,_ := io.ReadAll(r.Body)
	json.Unmarshal(byteValue,&company)
	company.Uuid = uuid.String()


	_,err = db.NamedExec(`INSERT INTO companies (uuid, name, email, logo_url, address, city, country, description, phone, industry) VALUES (:uuid, :name, :email, :logo_url, :address, :city, :country, :description, :phone, :industry)`,company)
	if err != nil{
		fmt.Println("Insert error:", err)
		http.Error(w,"Something went error",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status":"Success Added Company"})
}
