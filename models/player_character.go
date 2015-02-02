package models

import ()

type PlayerCharacter struct {
	Name               string
	Str                int
	Int                int
	Wis                int
	Dex                int
	Con                int
	Cha                int
	Level              int
	MaxHitPoints       int
	BaseArmorClass     int
	Class              Class
	ClassId            int
	Race               Race
	RaceId             int
	ProficiencyBonus   int
	ArmorClass         int
	InitiativeBonus    int
	SpeedInFeet        int
	CurrentHitPoints   int
	TemporaryHitPoints int
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
