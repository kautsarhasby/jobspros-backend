package main

import (
	"backend-fullstack/router"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func main(){
	r := chi.NewRouter()

	http.HandleFunc("/auth",router.AuthHandler)
	http.HandleFunc("/users/",router.UsersHandler)

	r.Route("/jobs", func(r chi.Router) {
		r.Get("/", router.GetJobs)                      
		r.Get("/{id}", router.GetJobById)   
	})
	r.Route("/companies", func(r chi.Router) {
		r.Get("/", router.GetCompanies)            
		r.Post("/", router.PostCompany)            
		r.Get("/{id}", router.GetCompanyById)   
	})

	PORT := 4770
	addr := fmt.Sprintf(":%d",PORT)
	fmt.Printf("server start at %v",PORT)
	http.ListenAndServe(addr,r)
}