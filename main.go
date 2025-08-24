package main

import (
	"backend-fullstack/lib"
	"backend-fullstack/router"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)


func main(){
	r := chi.NewRouter()
	
	r.Use(lib.Middleware)

	frontendOrigin := os.Getenv("FRONTEND_ORIGIN")
    if frontendOrigin == "" {
        frontendOrigin = "http://localhost:5173"
    }


	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{frontendOrigin}, 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token","X-API-KEY"},
		AllowCredentials: true,
	}))

	r.Route("/auth", func(r chi.Router){
		r.Post("/", router.AuthHandler)
	})

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