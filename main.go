package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.POST("/courses", postCourse)
	router.GET("/courses/:id", getCourseByID)
	router.GET("/courses/:teacher", getCourseByTeacher)
	router.Run("localhost:8080")
}

// Course represents data about a universiry course.
type Course struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Teacher  string  `json:"teacher"`
	Students int     `json:"students"`
	Rating   float32 `json:"rating"`
}

// course slice to seed course data.
var courses = []Course{
	{ID: "1", Title: "Algo", Teacher: "Thore", Students: 100, Rating: 9.1},
	{ID: "2", Title: "Code", Teacher: "Teach", Students: 200, Rating: 8.9},
	{ID: "3", Title: "Design", Teacher: "Anders", Students: 50, Rating: 8},
}

// getCourses responds with the list of all courses as JSON.
func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

// postCourses adds course from JSON received in the request body.
func postCourse(c *gin.Context) {
	var newCourse Course

	// Call BindJSON to bind the received JSON to
	// newCourse.
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	// Add the new course to the slice.
	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

//update student count from id
func UpdateStudents(c *gin.Context) {
	id := c.Param("id")
	strcount := c.Param("students")
	count, err := strconv.Atoi(strcount)
	if count < 0 || err == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
		return
	}

	for _, a := range courses {
		if a.ID == id {
			a.Students = count
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}

// getCourseByID locates the course whose ID value matches the id
// parameter sent by the client, then returns that course as a response.
func getCourseByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of courses, looking for
	// an album whose ID value matches the parameter.
	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}
func getCourseByTeacher(c *gin.Context) {
	teacher := c.Param("teacher")

	// Loop over the list of courses, looking for
	// a course whose ID value matches the parameter.
	for _, a := range courses {
		if a.Teacher == teacher {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "teacher not found"})
}
