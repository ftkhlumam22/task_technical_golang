package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test_golang/data"
)

// Initialisasi local db for save the data
var (
	localdb = make(map[int]data.Student)
)

// Making Response Functions
func setResponse(w http.ResponseWriter, message []byte, httpCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(message)
}

// Handler Default
func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	message := []byte(`{"message":"Rest API Student"}`)
	setResponse(w, message, http.StatusOK)
}

// Handler To Get Student
func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		message := []byte(`{"message" : "Method Not Allowed"}`)
		setResponse(w, message, http.StatusMethodNotAllowed)
		return
	}

	var students []data.Student

	for _, student := range localdb {
		students = append(students, student)
	}

	studentResp, err := json.Marshal(&students)
	if err != nil {
		message := []byte(`{"message" : "Error Parshing Data"}`)
		setResponse(w, message, http.StatusInternalServerError)
		return
	}

	setResponse(w, studentResp, http.StatusOK)
}

// Handler To Get Student By Id, Update Student, Delete Student
func StudentHandler(w http.ResponseWriter, r *http.Request) {
	//Register Method
	if r.Method == "POST" {
		var student data.Student

		payload := r.Body

		defer r.Body.Close()

		err := json.NewDecoder(payload).Decode(&student)

		if err != nil {
			message := []byte(`{"message" : "Error When Parsing Data"}`)
			setResponse(w, message, http.StatusMethodNotAllowed)
		}

		localdb[student.ID] = student

		message := []byte(`{"message" : "Register Successfully"}`)
		setResponse(w, message, http.StatusOK)
	}

	// Get Student By Id Method
	if r.Method == "GET" {
		if _, validationId := r.URL.Query()["id"]; !validationId {
			message := []byte(`{"message" : "Required Student ID"}`)
			setResponse(w, message, http.StatusBadRequest)
			return
		}

		req, _ := strconv.Atoi(r.URL.Query().Get("id"))
		student, validationStudent := localdb[req]
		if !validationStudent {
			message := []byte(`{"message" : "Student Not Found"}`)
			setResponse(w, message, http.StatusOK)
		}

		studentJSON, err := json.Marshal(student)
		if err != nil {
			message := []byte(`{"message" : "Failed to marshal student JSON"}`)
			setResponse(w, message, http.StatusInternalServerError)
			return
		}

		setResponse(w, studentJSON, http.StatusOK)
	}

	// Delete Student Method
	if r.Method == "DELETE" {
		if _, validationId := r.URL.Query()["id"]; !validationId {
			message := []byte(`{"message" : "Required Student ID"}`)
			setResponse(w, message, http.StatusBadRequest)
			return
		}

		req, _ := strconv.Atoi(r.URL.Query().Get("id"))
		_, validationStudent := localdb[req]

		if !validationStudent {
			message := []byte(`{"message" : "Student Not Found"}`)
			setResponse(w, message, http.StatusOK)
		}

		delete(localdb, req)

		message := []byte(`{"message" : "Delete Student Data Successfully"}`)
		setResponse(w, message, http.StatusOK)
	}

	// Update Student Data
	if r.Method == "PATCH" {
		if _, validationId := r.URL.Query()["id"]; !validationId {
			message := []byte(`{"message" : "Required Student ID"}`)
			setResponse(w, message, http.StatusBadRequest)
			return
		}

		req, _ := strconv.Atoi(r.URL.Query().Get("id"))
		student, validationStudent := localdb[req]
		if !validationStudent {
			message := []byte(`{"message" : "Student Not Found"}`)
			setResponse(w, message, http.StatusOK)
		}

		var newStudent data.Student

		payload := r.Body

		defer r.Body.Close()

		err := json.NewDecoder(payload).Decode(&newStudent)
		if err != nil {
			message := []byte(`{"message" : "Error When Parsing Data"}`)
			setResponse(w, message, http.StatusMethodNotAllowed)
		}

		if newStudent.Name != "" {
			student.Name = newStudent.Name
		}

		if newStudent.Age != 0 {
			student.Age = newStudent.Age
		}

		localdb[student.ID] = student

		message := []byte(`{"message" : "Update Student Data Successfully"}`)
		setResponse(w, message, http.StatusOK)
	}
}
