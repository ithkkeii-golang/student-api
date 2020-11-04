package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Teacher *Teacher `json:"teacher"`
}

type Teacher struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params

	// Loop through student and findById
	for _, student := range students {
		if student.ID == params["id"] {
			json.NewEncoder(w).Encode(student)
			return
		}
	}

	json.NewEncoder(w).Encode(&Student{})
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)

	student.ID = strconv.Itoa(rand.Intn(10000000))

	students = append(students, student)

	json.NewEncoder(w).Encode(student)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params

	// Loop through student and findById
	for _, student := range students {
		if student.ID == params["id"] {
			json.NewEncoder(w).Encode(student)
			return
		}
	}

	json.NewEncoder(w).Encode(&Student{})
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params

	// Loop through student and findById
	for index, student := range students {
		if student.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			json.NewEncoder(w).Encode(students)
		}
	}
}

func getTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func getTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// Init students var as a slice Student struct
var students []Student

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock data
	students = append(students, Student{ID: "1", Name: "kei", Teacher: &Teacher{ID: "1", Name: "kenny"}})
	students = append(students, Student{ID: "2", Name: "bob", Teacher: &Teacher{ID: "1", Name: "kenny"}})
	students = append(students, Student{ID: "3", Name: "bop", Teacher: &Teacher{ID: "1", Name: "n"}})

	// Route Handlers
	router.HandleFunc("/api/students", getStudents)
	router.HandleFunc("/api/students/{id}", getStudent)
	router.HandleFunc("/api/student", createStudent).Methods("POST")
	router.HandleFunc("/api/students/{id}", updateStudent).Methods("PUT")
	router.HandleFunc("/api/students/{id}", deleteStudent).Methods("DELETE")

	router.HandleFunc("/api/teachers", getTeachers)
	router.HandleFunc("/api/teachers/{id}", getTeacher)
	// router.HandleFunc("/api/teacher", createTeacher).Method("POST")
	// router.HandleFunc("/api/teachers/{id}", updateTeacher).Method("PUT")
	// router.HandleFunc("/api/teachers/{id}", deleteTeacher).Method("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

	fmt.Println("hello world")
}
