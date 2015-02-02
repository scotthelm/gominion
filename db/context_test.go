package db

import (
	"testing"
)

func TestContextCreation(t *testing.T) {
	c := NewContext("sqlite3", "/tmp/gominion_test.db")
	db := c.DB
	dbaddr := &db
	if dbaddr == nil {
		t.Fatal("unable to create db context")
	}
}

func TestSchemaCreation(t *testing.T) {
	c := NewContext("sqlite3", "/tmp/gominion_test.db")
	c.Migrate()
}
