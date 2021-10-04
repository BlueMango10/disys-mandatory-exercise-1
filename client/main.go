package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Course struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Workload     int     `json:"workload"`
	Satisfaction float32 `json:"satisfaction"`
	Teacher      int     `json:"teacher"`
	Students     []int   `json:"students"`
}

const (
	host = "http://localhost:8080"
	basePath = "/itucm"
	path = host + basePath
	contentType = "application/json"
)

func main() {
	fmt.Println("Welcome to itucm client. Choose an option:")
	
	for {
		// Instructions
		fmt.Println("=== Choose action ===")
		fmt.Println("0: List all courses")
		fmt.Println("1: Find specific course")
		fmt.Println("2: Create new course")
		fmt.Println("3: Update course information")
		fmt.Println("4: Remove course")
		
		// Get input
		fmt.Print("[0-4]> ")
		var input int
		fmt.Scanln(&input)

		// Choose action
		switch input {
		case 0:
			listCourses()
		case 1:
			showCourse()
		case 2:
			createCourse()
		case 3:
			updateCourse()
		case 4:
			removeCourse()
		default:
			fmt.Print("Input not recognized")
		}

		// Wait for input
		fmt.Print("Press [Enter] to continue> ")
		fmt.Scanln()
	}
}

func listCourses() {
	resp, err := http.Get(path + "/course")
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func showCourse() {
	// Get input
	fmt.Print("Course Id:> ")
	var input int
	fmt.Scanln(&input)
	
	resp, err := http.Get(path + "/course/" + strconv.Itoa(input))
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func createCourse() {
	scanner := bufio.NewScanner(os.Stdin)
	// Create new course struct
	var newCourse Course
	fmt.Print("Id (must not be in use):> ")
	fmt.Scanln(&newCourse.Id)
	fmt.Print("Name:> ")
	if scanner.Scan() {
		newCourse.Name = scanner.Text()
	}
	fmt.Print("Workload:> ")
	fmt.Scanln(&newCourse.Workload)
	fmt.Print("Satisfaction:> ")
	fmt.Scanln(&newCourse.Satisfaction)
	fmt.Print("Teacher Id:> ")
	fmt.Scanln(&newCourse.Teacher)
	fmt.Print("Students (seperate with ','):> ")
	var str string
	fmt.Scanln(&str)
	studentStrs := strings.Split(str, ",")
	for _, studentStr := range studentStrs {
		student, _ := strconv.Atoi(studentStr)
		newCourse.Students = append(newCourse.Students, student)
	}
	asJson, _ := json.Marshal(newCourse)
	http.Post(path + "/course", contentType, bytes.NewReader(asJson))
}

func updateCourse() {
	// Get id
	fmt.Print("Course Id:> ")
	var id int
	fmt.Scanln(&id)
	
	scanner := bufio.NewScanner(os.Stdin)
	// Create new course struct
	var newCourse Course
	newCourse.Id = id
	fmt.Print("Name:> ")
	if scanner.Scan() {
		newCourse.Name = scanner.Text()
	}
	fmt.Print("Workload:> ")
	fmt.Scanln(&newCourse.Workload)
	fmt.Print("Satisfaction:> ")
	fmt.Scanln(&newCourse.Satisfaction)
	fmt.Print("Teacher Id:> ")
	fmt.Scanln(&newCourse.Teacher)
	fmt.Print("Students (seperate with ','):> ")
	var str string
	fmt.Scanln(&str)
	studentStrs := strings.Split(str, ",")
	for _, studentStr := range studentStrs {
		student, _ := strconv.Atoi(studentStr)
		newCourse.Students = append(newCourse.Students, student)
	}
	asJson, _ := json.Marshal(newCourse)
	putRequest(path + "/course/" + strconv.Itoa(id), bytes.NewReader(asJson))
}

func removeCourse() {
	// Get id
	fmt.Print("Course Id:> ")
	var id int
	fmt.Scanln(&id)

	deleteRequest(path + "/course/" + strconv.Itoa(id), nil)
}


func putRequest(url string, data io.Reader) {
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodPut, url, data)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func deleteRequest(url string, data io.Reader) {
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodDelete, url, data)
	if err != nil {
		fmt.Println("ERR: " + err.Error())
	}
	_, err = client.Do(req)
	if err != nil {
		fmt.Println("ERR: " + err.Error())
	}
}
