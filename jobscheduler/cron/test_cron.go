package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	// schedule a task to run every minute
	c.AddFunc("* * * * *", func() {
		fmt.Println("This task runs every minute")
	})

	// schedule a task to run every 5 seconds
	c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("This task runs every 5 seconds")
	})

	// start the cron job
	c.Start()

	// wait for 10 seconds
	time.Sleep(10 * time.Second)

	// stop the cron job
	c.Stop()

}
