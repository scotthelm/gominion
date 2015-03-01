package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/GeertJohan/go.rice"
	"github.com/scotthelm/gominion/db"
	"github.com/scotthelm/gominion/web"
)

func main() {

	f, err := os.OpenFile("./gominion.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Could not open logfile")
	}

	defer f.Close()

	log.SetOutput(io.MultiWriter(os.Stdout, f))

	router := web.NewRouter()
	router.PathPrefix("/public/").Handler(assetsHandler())
	router.PathPrefix("/templates/").Handler(templatesHandler())

	web.Ctx = db.NewContext("sqlite3", "./gominion.db")
	web.Ctx.Migrate()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func assetsHandler() http.Handler {
	assets := rice.MustFindBox("public")
	return http.StripPrefix("/public/", http.FileServer(assets.HTTPBox()))
}

func templatesHandler() http.Handler {
	assets := rice.MustFindBox("templates")
	return http.StripPrefix("/templates/", http.FileServer(assets.HTTPBox()))
}
