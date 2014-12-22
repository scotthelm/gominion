package gominion

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
)

func MakeStructure() {
	schema, err := ioutil.ReadFile("./structure.sql")
	if err != nil {
		log.Fatalln(err)
	}
	db, err := sqlx.Connect(
		"postgres",
		"postgres://gominion_login:gominion_login_pass@localhost/gominion")

	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(string(schema))
}
