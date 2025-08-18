package router

import (
	"backend-fullstack/lib"
	"backend-fullstack/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)


type JobRequest struct {
	Uuid          string   `json:"uuid" db:"uuid"`
	PublisherId   string   `json:"publisher_id" db:"publisher_id"`
	Position      string   `json:"position" db:"position"`
	Qualification string   `json:"qualification" db:"qualification"`
	Description   string   `json:"description" db:"description"`
	ClosingDate  *time.Time   `json:"closing_date" db:"closing_date"`
	Salary      *string  `json:"salary" db:"salary"`
}


func GetJobs(w http.ResponseWriter, r * http.Request){
	var jobs []models.Job
	db, err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	
	err = db.Select(&jobs,"SELECT * FROM jobs where deleted_at IS NULl",)
	if err != nil{
		fmt.Println(err)
		return
	}


	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(jobs)
}

func GetJobById(w http.ResponseWriter, r * http.Request){
	id := chi.URLParam(r,"id")
	db,err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	var job models.Job
	err = db.Get(&job,"SELECT * FROM jobs WHERE id = ?",id)
	if err != nil{
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(job)
}

func PostJob(w http.ResponseWriter, r *http.Request){
	var job JobRequest
	uuid := uuid.New()
	db,err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	byteValue,_ := io.ReadAll(r.Body)
	json.Unmarshal(byteValue,&job)
	job.Uuid = uuid.String()
	

	_,err = db.NamedExec(`INSERT INTO jobs (uuid,name,email,password,role) VALUES (:uuid,:name,:email,:password,:role)`,job)
	if err != nil{
		fmt.Println("Insert error:", err)
		http.Error(w,"Something went error",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status":"Success Added job"})
}
