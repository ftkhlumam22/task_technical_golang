package main

import (
	"fmt"
	"net/http"
	"os"
	"test_golang/handler"
)

func main() {
	// Default Route
	http.HandleFunc("/", handler.DefaultRoute)

	//Route for Fetching All Data Student on DB
	http.HandleFunc("/students", handler.GetAllStudent)

	// Route for register, Get Student By Id, Update Student, And Delete Student
	http.HandleFunc("/student", handler.StudentHandler)

	// Initialisasi Server Running
	fmt.Println("Server Running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
