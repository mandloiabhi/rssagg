package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	fmt.Println("hello world")
	router:=chi.NewRouter()
	// CONFIGURATION 
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","DELETE","PUT","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	srv:= &http.Server{
		Handler: router,
		Addr: ":8080",
	}
	v1Router:=chi.NewRouter()
	v1Router.Get("/abhi",handlerReadiness)
	router.Mount("/v1",v1Router)
	err:=srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
	
	


}
