package database

import (
	"fmt"
	"go-schedule-email-task/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	Database, err = gorm.Open(sqlite.Open("Employee.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

func Migrate() {
	Connect()
	err := Database.AutoMigrate(&model.Employee{})
	if err != nil {
		return
	}
}
