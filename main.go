package main

import (
	"go-schedule-email-task/database"
	"go-schedule-email-task/helper"
	"go-schedule-email-task/router"
)

func main() {
	helper.LoadEnv()
	database.Migrate()
	go helper.Cronjob()
	router.UrlRouter()

}
