package main

import (
	"github.com/scotthelm/gominion/db"
	"github.com/scotthelm/gominion/web"
	"log"
	"net/http"
)

func main() {

	router := web.NewRouter()
	web.Ctx = db.NewContext("sqlite3", "./gominion.db")
	web.Ctx.Migrate()
	log.Fatal(http.ListenAndServe(":8080", router))
}
