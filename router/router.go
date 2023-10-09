package router

import (
	"github.com/gin-gonic/gin"
	"go-schedule-email-task/controller"
)

func UrlRouter() {

	router := gin.Default()
	router.POST("empl/", controller.EmpHandler)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
