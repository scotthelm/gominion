package models

import (
	"database/sql"
)

type Class struct {
	Id              int
	Name            string
	HitDieNumber    int
	HitDieFrequency int
	HitDieAverage   int
	Proficiencies   []Proficiency
	Skills          []Skill
}

type Proficiency struct {
	Id   int
	Name string
	Type ProficiencyType
}

type ProficiencyType struct {
	Id   int
	Name string
}

type Skill struct {
	Id   int
	Name string
}

type PlayerCharacterSkill struct {
	Id                int
	PlayerCharacter   PlayerCharacter
	PlayerCharacterId int
	Skill             Skill
	SkillId           int
	Value             int
}

type Equipment struct {
	Id           int
	Name         string
	Description  string
	WeightInGold int
	AttackBonus  int
	DamageBonus  int
	DamageType   DamageType
	DamageTypeId sql.NullInt64
}

type DamageType struct {
	Id   int
	Name string
}

type Spell struct {
	Id          int
	Name        string
	SpellType   SpellType
	Range       int
	V           bool
	S           bool
	M           bool
	Description string
}

type SpellType struct {
	Id   int
	Name string
}
