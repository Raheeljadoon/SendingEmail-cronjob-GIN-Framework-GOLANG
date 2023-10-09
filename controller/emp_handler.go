package controller

import (
	"github.com/gin-gonic/gin"
	"go-schedule-email-task/database"
	"go-schedule-email-task/model"
	"net/http"
)

func EmpHandler(c *gin.Context) {
	var input model.Employee

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Database.Create(&input)
	c.JSON(http.StatusCreated, input)
}
