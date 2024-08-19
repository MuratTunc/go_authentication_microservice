package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPORT = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	log.Println("Authentication service is started...")

	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
