package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Teacher struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Popularity float32 `json:"popularity"`
}

type Course struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Workload     int     `json:"workload"`
	Satisfaction float32 `json:"satisfaction"`
	Teacher      int     `json:"teacher"`
	Students     []int   `json:"student"`
}

var students = []Student{
	{Id: 0, Name: "Byron Crowther"},
	{Id: 1, Name: "Ellesha Rutledge"},
	{Id: 2, Name: "Clayton Cantu"},
	{Id: 3, Name: "Shirley Atkins"},
	{Id: 4, Name: "Brodie Griffiths"},
}

var teachers = []Teacher{
	{Id: 0, Name: "Mert Sandoval", Popularity: 0.9},
	{Id: 1, Name: "Aiysha Pratt", Popularity: 0.7},
}

var courses = []Course{
	{
		Id:           0,
		Name:         "Basic basics",
		Workload:     15,
		Satisfaction: 0.75,
		Teacher:      0,
		Students:     []int{0, 1, 2},
	},
	{
		Id:           1,
		Name:         "Rudimentary roots 101",
		Workload:     7,
		Satisfaction: 0.95,
		Teacher:      1,
		Students:     []int{2, 3, 4},
	},
}

const basePath = "/itucm"

// GET /course
func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

// POST /course
func postCourse(c *gin.Context) {
	var newCourse Course

	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	courses = append(courses, newCourse)
	c.Writer.WriteHeader(http.StatusCreated)
}

// GET /course/:id
func getCourseById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if (err != nil) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, course := range courses {
		if course.Id == id {
			c.IndentedJSON(http.StatusOK, course)
			return
		}
	}

	c.Writer.WriteHeader(http.StatusNotFound)
}

// PUT /course/:id
func updateCourse(c *gin.Context) {

}

// DELETE /course/:id
func deleteCourse(c *gin.Context) {

}

func main(){
	router := gin.Default()
	router.GET(basePath + "/course", getCourses)
	router.POST(basePath + "/course", postCourse)
	router.GET(basePath + "/course/:id", getCourseById)
	router.PUT(basePath + "/course:id", updateCourse)
	router.DELETE(basePath + "/course:id", deleteCourse)

	router.Run("localhost:8080")
}
