package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RafaelCavalcante17/api-students/db"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudent(0) // Pass 0 or another appropriate integer argument
	if err != nil {
		return c.String(http.StatusNotFound, "Falid to get students")
	}
	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}

	return c.String(http.StatusOK, "Create student")
}

func (api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ID")
	}

	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")

	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ")
	}

	return c.JSON(http.StatusOK, student)
}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	UpdateStud := fmt.Sprintf("Update %s students", id)
	return c.String(http.StatusOK, UpdateStud)
}

func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Delte %s students", id)
	return c.String(http.StatusOK, deleteStud)
}
