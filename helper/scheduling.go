package helper

import (
	"fmt"
	"github.com/claudiu/gocron"
)

func Cronjob() {
	gocron.Clear()
	fmt.Print("Scheduling Function Starts")

	gocron.Start()
	s := gocron.NewScheduler()
	gocron.Every(5).Seconds().Do(func() {
		fmt.Print("\nCalling Email Sending Function\n")
		SendingEmailOperation()
	})
	<-s.Start()

}
