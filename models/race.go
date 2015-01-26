package models

import ()

type Race struct {
	Id          int
	Name        string
	StrMod      int
	IntMod      int
	WisMod      int
	DexMod      int
	ConMod      int
	ChaMod      int
	HitPointMod int
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
