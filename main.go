package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nithinchandra/authservice/app"
	"github.com/nithinchandra/authservice/controllers"
	"net/http"
	"os"
)

func main(){
	router := mux.NewRouter()

	router.Use(app.JwtAutherntication) //attach JWT auth middleware

	router.HandleFunc("/api/user/new",controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	port := os.Getenv("PORT")// Get port from .env file

	if port == ""{
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port,router)

	if err != nil{
		fmt.Println(err)
	}

}
