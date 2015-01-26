package models

import (
	"testing"
	"time"
)

func TestEntityCreate(t *testing.T) {
	r := Race{"TestRace", "TestSubRace", 1, 0, 0, 0, 1, 0, 1, Entity{0, time.Now()}}
	ent, err := r.Create()
	if err != nil {
		t.Errorf("%s#Create produced error: '%s'", ent, err)
	}
}

func TestFieldNumbers(t *testing.T) {
	r := Race{"TestRace", "TestSubRace", 1, 0, 0, 0, 1, 0, 1, Entity{0, time.Now()}}
	insert_sql := CreateSql(r)
	if insert_sql != "insert into race (name, subrace, str_mod, int_mod, wis_mod, dex_mod, con_mod, cha_mod, hit_point_mod) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)" {
		t.Errorf("Error in CreateSql: %s", insert_sql)
	}
}
