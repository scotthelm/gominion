package gominion

import ()

type PlayerCharacter struct {
	Name           string `db:"name"`
	RolledStr      int    `db:"rolled_str"`
	RolledInt      int    `db:"rolled_int"`
	RolledWis      int    `db:"rolled_wis"`
	RolledDex      int    `db:"rolled_dex"`
	RolledCon      int    `db:"rolled_con"`
	RolledCha      int    `db:"rolled_cha"`
	RaceId         int    `db:"race_id"`
	ClassId        int    `db:"class_id"`
	Level          int    `db:"level"`
	MaxHitPoints   int    `db:"max_hit_points"`
	BaseArmorClass int    `db:"base_armor_class"`
}

func (p *PlayerCharacter) Str() int {
	return p.RolledStr + p.Race().StrMod
}

func (p *PlayerCharacter) Int() int {
	return p.RolledInt + p.Race().IntMod
}

func (p *PlayerCharacter) Wis() int {
	return p.RolledWis + p.Race().WisMod
}

func (p *PlayerCharacter) Dex() int {
	return p.RolledDex + p.Race().DexMod
}

func (p *PlayerCharacter) Con() int {
	return p.RolledCon + p.Race().ConMod
}

func (p *PlayerCharacter) Cha() int {
	return p.RolledCha + p.Race().ChaMod
}

func (pc *PlayerCharacter) Create() (*PlayerCharacter, error) {

	return pc, nil
}

func (pc *PlayerCharacter) Update() (*PlayerCharacter, error) {
	return pc, nil
}

func (pc *PlayerCharacter) Delete() (bool, error) {
	return false, nil
}

func (pc *PlayerCharacter) Race() Race {
	return Dwarf()
}
