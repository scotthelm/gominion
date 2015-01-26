package db

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Entity struct {
	Id        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
}

type SqlEntity interface {
	Create() (interface{}, error)
	Update() (interface{}, error)
	Delete() (bool, error)
}

func CreateSql(e interface{}) string {
	var buffer bytes.Buffer
	var values bytes.Buffer
	t := reflect.TypeOf(e)
	counter := 1
	buffer.WriteString(fmt.Sprintf("insert into %s (", strings.ToLower(t.Name())))
	for fieldPos := 0; fieldPos < t.NumField(); fieldPos++ {
		tag := t.Field(fieldPos).Tag.Get("db")
		if tag != "" {
			buffer.WriteString(fmt.Sprintf("%s, ", t.Field(fieldPos).Tag.Get("db")))
			values.WriteString(fmt.Sprint("?, "))
			counter++
		}
	}
	buffer.Truncate(buffer.Len() - 2)
	values.Truncate(values.Len() - 2)
	buffer.WriteString(") values (")
	buffer.WriteString(values.String())
	buffer.WriteString(")")
	return buffer.String()
}
