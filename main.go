package main

import (
	"Goblog/controllers"
	"Goblog/middlewares"
	"Goblog/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		utils.Respond(writer, utils.Message(true, "Welcome to GOBLOG API"))
	}).Methods("GET")
	router.HandleFunc("/api/register", controllers.SignUp).Methods("POST")
	router.HandleFunc("/api/login", controllers.SignIn).Methods("POST")
	router.Use(middlewares.JWT)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Print(port)
	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}


