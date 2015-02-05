package main

import (
	"github.com/scotthelm/gominion/db"
	"github.com/scotthelm/gominion/web"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	f, err := os.OpenFile("./gominion.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Could not open logfile")
	}
	defer f.Close()

	log.SetOutput(io.MultiWriter(os.Stdout, f))
	router := web.NewRouter()
	web.Ctx = db.NewContext("sqlite3", "./gominion.db")
	web.Ctx.Migrate()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func setLogger() {
}
