package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Student struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Branch string  `json:"branch"`
	Year   int     `json:"year"`
	CGPA   float64 `json:"cgpa"`
}

var students = []Student{
	{ID: 1, Name: "Abhi", Branch: "Computer Science", Year: 2, CGPA: 8.5},
	{ID: 2, Name: "Vijay", Branch: "IT", Year: 3, CGPA: 7.8},
	{ID: 3, Name: "Asher", Branch: "Electronics", Year: 1, CGPA: 9.1},
}

func main() {
	router := gin.Default()

	router.GET("/", homeHandler)
	router.GET("/api/students", getAllStudents)
	router.GET("/api/students/:id", getStudentByID)
	router.POST("/api/students", addStudent)
	router.DELETE("/api/students/:id", deleteStudent)

	router.Run(":8080")
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Student API in Go 🚀",
		"version": "1.0",
	})
}

func getAllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

func getStudentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, student := range students {
		if student.ID == id {
			c.JSON(http.StatusOK, student)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
}

func addStudent(c *gin.Context) {
	var newStudent Student
	c.BindJSON(&newStudent)
	students = append(students, newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

func deleteStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Student deleted!"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
}
