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

func GetResumes(w http.ResponseWriter, r * http.Request){
	var resumes []models.Resume
	db, err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()
	
	err = db.Select(&resumes,"SELECT * FROM resumes where deleted_at IS NULl",)
	if err != nil{
		fmt.Println(err)
		return
	}


	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(resumes)
}

func GetResumeById(w http.ResponseWriter, r * http.Request){
	id := chi.URLParam(r,"id")
	db,err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	var resume models.Resume
	err = db.Get(&resume,"SELECT * FROM resumes WHERE id = ?",id)
	if err != nil{
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(resume)
}

func PostResume(w http.ResponseWriter, r *http.Request){
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
	

	_,err = db.NamedExec(`INSERT INTO resumes (uuid,job_id,applicant_id,resume_url) VALUES (:uuid,:job_id,:applicant_id,:resume_url)`,job)
	if err != nil{
		fmt.Println("Insert error:", err)
		http.Error(w,"Something went error",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status":"Success Added Resumes"})
}
