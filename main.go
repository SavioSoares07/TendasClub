package main

import (
	"fmt"
	"net/http"
	"tendasclub/handlers"
)

func main() {
	//Set up the Http Server
	http.HandleFunc("/register", handlers.SignUpHandler)

	//Start the server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
