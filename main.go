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
	

	r.Route("/users",func (r chi.Router)  {
		r.Get("/", router.GetUsers)                      
		r.Get("/{id}", router.GetUserById) 
	})

	r.Route("/jobs", func(r chi.Router) {
		r.Get("/", router.GetJobs)                      
		r.Get("/{id}", router.GetJobById)   
		r.Post("/",router.PostJob)
		r.Put("/{id}", router.UpdateJob)
	})
	r.Route("/companies", func(r chi.Router) {
		r.Get("/", router.GetCompanies)            
		r.Get("/{id}", router.GetCompanyById)   
		r.Post("/", router.PostCompany)            
	})
	r.Route("/hr", func(r chi.Router) {
		r.Get("/", router.GetHR)            
		r.Get("/{id}", router.GetHRById)   
		r.Post("/", router.PostHR)            
	})

	PORT := 4770
	addr := fmt.Sprintf(":%d",PORT)
	fmt.Printf("server start at %v",PORT)
	http.ListenAndServe(addr,r)
}