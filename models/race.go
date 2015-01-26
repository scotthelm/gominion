package models

import (
	"github.com/scotthelm/gominion/db"
)

type Race struct {
	Name        string `db:"name"`
	Subrace     string `db:"subrace"`
	StrMod      int    `db:"str_mod"`
	IntMod      int    `db:"int_mod"`
	WisMod      int    `db:"wis_mod"`
	DexMod      int    `db:"dex_mod"`
	ConMod      int    `db:"con_mod"`
	ChaMod      int    `db:"cha_mod"`
	HitPointMod int    `db:"hit_point_mod"`
	db.Entity
}

func (r *Race) Create() (*Race, error) {
	return r, nil
}

func (r *Race) Update() (*Race, error) {
	return r, nil
}

func (r *Race) Delete() (bool, error) {
	return false, nil
}
