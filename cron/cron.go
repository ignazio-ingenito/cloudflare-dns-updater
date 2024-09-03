package cron

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

func Setup(f func()) (*cron.Cron, error) {
	// Set the timezone
	loc, err := time.LoadLocation("Europe/Rome")
	if err != nil {
		log.Panicln(err)
	}

	// Set the cron schedule
	schedule := os.Getenv("APP_CRON_SCHEDULE")
	if schedule == "" {
		schedule = "0 */4 * * *"
	}
	fmt.Println("Cron schedule: " + schedule)

	// Create the cronjob
	token := strings.Split(schedule, " ")
	c := &cron.Cron{}
	if len(token) == 6 {
		log.Println("Using cron with seconds")
		c = cron.New(cron.WithSeconds(), cron.WithLocation(loc))
	} else {
		log.Println("Using cron without seconds")
		c = cron.New(cron.WithLocation(loc))
	}
	// Add a job to the cron
	id, err := c.AddFunc(schedule, f)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("Cron id: %d\n", id)

	// Start the cron
	c.Start()
	return c, nil
}
