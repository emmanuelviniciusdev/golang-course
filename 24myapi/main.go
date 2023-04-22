package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

var courses []Course = []Course{}

func (course *Course) IsEmpty() bool {
	return course.Name == ""
}

func (course *Course) IsNotEmpty() bool {
	return !course.IsEmpty()
}

func main()  {
	fmt.Println("API - LearnCodeOnline.in")

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteOneCourse")
	
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for idx, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:idx], courses[idx + 1:]...)
		}
	}

	json.NewEncoder(w).Encode(nil)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateOneCourse")
	
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	var updatedCourse Course

	for idx, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:idx], courses[idx + 1:]...)

			json.NewDecoder(r.Body).Decode(&updatedCourse)

			// Rewriting Id (idk why)
			updatedCourse.Id = params["id"]

			courses = append(courses, updatedCourse)
		}
	}

	json.NewEncoder(w).Encode(updatedCourse)
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createOneCourse")
	
	w.Header().Set("content-type", "application/json")

	if r.Body == nil {
		fmt.Println("Body is null")
		json.NewEncoder(w).Encode(nil)
		return
	}

	var newCourse Course

	json.NewDecoder(r.Body).Decode(&newCourse)

	if newCourse.IsEmpty() {
		fmt.Println("Course is empty")
		json.NewEncoder(w).Encode(nil)
		return
	}

	newCourseId := len(courses) + 1

	newCourse.Id = fmt.Sprint(newCourseId)

	courses = append(courses, newCourse)

	json.NewEncoder(w).Encode(newCourse)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getOneCourse")
	
	w.Header().Set("content-type", "application/json")
	
	params := mux.Vars(r)

	var courseFound Course

	for _, course := range courses {
		if course.Id == params["id"] {
			courseFound = course
		}
	}

	if courseFound.IsNotEmpty() {
		json.NewEncoder(w).Encode(courseFound)
	} else {
		json.NewEncoder(w).Encode(nil)
	}
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getAllCourses")

	w.Header().Set("content-type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline.in</h1>"))
}
