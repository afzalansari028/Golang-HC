package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId   string `json:"courseid"`
	CourseName string `json:"coursename"`
	// CoursePrice int     `json:"-"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - Learning in Golang")
	router := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "Hitesh Chaudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 399, Author: &Author{FullName: "Hitesh Chaudhary", Website: "go.dev"}})

	//routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOnecourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", router))

}

//controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the API by golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	//loop through courses, find matching id and return reponse
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
	// w.WriteHeader(http.StatusNotFound)
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// //what about - {}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// //TODO: check only if title is duplicate
	// //loop, title matches with course.courseName
	for _, courseN := range courses {
		if courseN.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("This course is already exist, please try with other")
			return
		}
	}

	// //generate unique id, string
	// //append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from request
	params := mux.Vars(r)

	//loop , id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//send response when id is not found
	//json.NewEncoder(w).Encode("Please send the existing id")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var removed Course
	//loop, id, remove (index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			removed = course
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(removed)
			break
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

//----------------code for update-----------
// func UpdateOneEmployee(c *gin.Context) {
// 	fmt.Println("Update function called")
// 	var newEmployee Employee
// 	jsonBytes, _ := ioutil.ReadAll(c.Request.Body)
// 	json.Unmarshal(jsonBytes, &newEmployee)

// 	fmt.Println("req emp", newEmployee)
// 	for i := range employees {
// 		if employees[i].Id == newEmployee.Id {
// 			fmt.Println("id matched")
// 			employees[i].Name = newEmployee.Name
// 			employees[i].Age = newEmployee.Age
// 		}
// 	}
// 	c.JSON(200, "Updated successfully")
// }
