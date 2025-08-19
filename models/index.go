package models

import "time"

type User struct {
	Id        int     `json:"id" db:"id"`
	Uuid      string  `json:"uuid" db:"uuid"`
	Name      string  `json:"name" db:"name"`
	Email     string  `json:"email" db:"email"`
	Password  string  `json:"password" db:"password"`
	Role      string  `json:"role" db:"role"`
	IsValidated      bool  `json:"is_validated" db:"is_validated"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type Job struct {
	Id            int      `json:"id" db:"id"`
	Uuid          string   `json:"uuid" db:"uuid"`
	PublisherId   string   `json:"publisher_id" db:"publisher_id"`
	Position      string   `json:"position" db:"position"`
	Qualification string   `json:"qualification" db:"qualification"`
	Description   string   `json:"description" db:"description"`
	ClosingDate   *time.Time   `json:"closing_date" db:"closing_date"`
	Salary      *string  `json:"salary" db:"salary"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time  `json:"deleted_at" db:"deleted_at"`
}

type Company struct {
	Id        int     `json:"id" db:"id"`
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
	IsValidated      bool  `json:"is_validated" db:"is_validated"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}


type Resume struct {
	Id        int     `json:"id" db:"id"`
	Uuid      string  `json:"uuid" db:"uuid"`
	JobId     string  `json:"job_id" db:"job_id"`
	ApplicantId     string  `json:"applicant_id" db:"applicant_id"`
	ResumeUrl    string  `json:"resume_url" db:"resume_url"`
	SubmittedAt time.Time  `json:"submitted_at" db:"submitted_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type HR struct {
	Id              int    `json:"id" db:"id"`
	HRUuid         string `json:"hr_uuid" db:"hr_uuid"`
	HRName          string `json:"hr_name" db:"hr_name"`
	HREmail         string `json:"hr_email" db:"hr_email"`
	CompanyName     string `json:"company_name" db:"company_name"`
	CompanyIndustry string `json:"company_industry" db:"company_industry"`
	CompanyUUID     string `json:"comp_uuid" db:"comp_uuid"`
}