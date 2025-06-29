package main

import (
	"fmt"
	"net/http"
	"tendasclub/database"
	"tendasclub/handlers"
)

func main() {
	//Connection Database
	database.ConnectionDB()

	//Set up the Http Server
	http.HandleFunc("/register", handlers.SignUpHandler)
	http.HandleFunc("/login", handlers.SignInHandler)

	//Start the server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
