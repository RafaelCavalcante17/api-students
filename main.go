package main

import (
	"fmt"
	"net/http"

	"github.com/RafaelCavalcante17/api-students/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getStudents)
	e.POST("/students", createStudent)
	e.GET("/students/:id", getStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	//- Post /students - Create student
	//- GET /students/:id - get infos from a specific student
	//- PUT /students/:id - Update student
	//- DELETE /students/:id - Delete student

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students")
}

func createStudent(c echo.Context) error {
	db.AddStudent()
	return c.String(http.StatusOK, "Create student")
}

func getStudent(c echo.Context) error {
	id := c.Param("id")
	getStud := fmt.Sprintf("Get %s students", id)
	return c.String(http.StatusOK, getStud)
}
func updateStudent(c echo.Context) error {
	id := c.Param("id")
	UpdateStud := fmt.Sprintf("Update %s students", id)
	return c.String(http.StatusOK, UpdateStud)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Delte %s students", id)
	return c.String(http.StatusOK, deleteStud)
}
