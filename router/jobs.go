package router

import (
	"backend-fullstack/lib"
	"backend-fullstack/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)


type JobRequest struct {
	Id *int  `json:"id" db:"id"`
	Uuid          string   `json:"uuid" db:"uuid"`
	PublisherId   int   `json:"publisher_id" db:"publisher_id"`
	Position      string   `json:"position" db:"position"`
	Qualification string   `json:"qualification" db:"qualification"`
	Description   string   `json:"description" db:"description"`
	ClosingDate  *time.Time   `json:"closing_date" db:"closing_date"`
	Salary      *string  `json:"salary" db:"salary"`
}


func GetJobs(w http.ResponseWriter, r *http.Request) {
    var jobs []models.Job
    db, err := lib.Connection()
    if err != nil {
        http.Error(w, "Database connection failed", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    err = db.Select(&jobs, "SELECT * FROM jobs WHERE deleted_at IS NULL")
    if err != nil {
        http.Error(w, "Query failed: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
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
	

	_, err = db.NamedExec(`
    INSERT INTO jobs (uuid, publisher_id, position, qualification, description, closing_date, salary)
    VALUES (:uuid, :publisher_id, :position, :qualification, :description, :closing_date, :salary)
`, job)
	if err != nil{
		fmt.Println("Insert error:", err)
		http.Error(w,"Something went error",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status":"Success Added job"})
}


func UpdateJob(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r,"id")
	db, err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	byteValue,_ := io.ReadAll(r.Body)

	var job JobRequest
	json.Unmarshal(byteValue,&job)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	job.Id = &idInt
	
	_, err = db.NamedExec(`
    UPDATE jobs 
    SET uuid = :uuid,
        publisher_id = :publisher_id,
        position = :position,
        qualification = :qualification,
        description = :description,
        closing_date = :closing_date,
        salary = :salary
    WHERE id = :id
`, job)
	if err != nil{
		fmt.Println("Update error:", err)
		http.Error(w,"Something went error",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status":"success","message":"Success added job"})

}

// Delete job : /jobs/3
func DeleteJob(w http.ResponseWriter, r *http.Request, id string){
	db, err := lib.Connection()
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE jobs set deleted_at = NOW() WHERE id = ?",id)
	if err != nil{
		fmt.Println("Delete error:", err)
		http.Error(w,"Job not found",http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status":"Success deleting job"})
}