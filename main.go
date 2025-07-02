package main

import (
	"fmt"
	"net/http"
	"tendasclub/database"
	"tendasclub/handlers"
	"tendasclub/middleware"
)

func main() {
	//Connection Database
	database.ConnectionDB()

	//Set up the Http Server
	//Method Post
	http.HandleFunc("/api/register", handlers.SignUpHandler)
	http.HandleFunc("/api/login", handlers.SignInHandler)
	http.HandleFunc("/api/registerTime", middleware.AuthMiddleware(handlers.RegisterTimeHandler))

	//Method Get
	http.HandleFunc("/api/getallrecord", handlers.GetAllTimeRecords)
	http.HandleFunc("/api/getrecordbyemail", middleware.AuthMiddleware(handlers.GetTimeRecordByEmail))

	//Start the server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
