package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func handleHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}

func main(){
	
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/",handleHome)
	server := &http.Server{
		Addr: ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if(err!=nil){
		fmt.Println("An error occured")
	}

}