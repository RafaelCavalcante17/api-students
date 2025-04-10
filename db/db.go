package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	CPF    int
	Email  string
	Age    int
	Active bool
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func AddStudent() {
	db := Init()

	student := Student{
		Name:   "Rafael",
		CPF:    12345,
		Email:  "rafael@gmail.com",
		Age:    39,
		Active: true,
	}

	if result := db.Create(&student); result.Error != nil {
		fmt.Println("Erro to create student")
	}

	fmt.Println("Create student!")
}
