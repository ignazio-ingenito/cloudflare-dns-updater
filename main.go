package main

import (
	"dnsupdater/cron"
	"dnsupdater/db"
	"dnsupdater/models"
	"dnsupdater/web"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// conncect to db
	sql, err := db.Connect()
	if err != nil {
		log.Panicln(err)
	}
	// get connection to defer the close
	conn, err := sql.DB()
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	// Setup the db if needed
	err = models.Setup(sql)
	if err != nil {
		log.Panicln(err)
	}

	// Start the cronjob
	c, err := cron.Setup(func() {
		db.PublicIpLogCreate(sql)
	})
	if err != nil {
		log.Panicln(err)
	}
	defer c.Stop()

	// Setup the server routes
	web.SetupRoutes(sql)

	// Start the server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "80"
	}
	log.Printf("Listening on port %s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panicln(err)
	}
}
