package models

import (
	"database/sql"
)

type Campaign struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Class struct {
	Id              int           `json:"id"`
	Name            string        `json:"name"`
	HitDieNumber    int           `json:"hit_die_number"`
	HitDieFrequency int           `json:"hit_die_frequency"`
	HitDieAverage   int           `json:"hit_die_average"`
	Proficiencies   []Proficiency `gorm:"many2many:class_proficiencies;";json:"proficiencies"`
	Skills          []Skill       `gorm:"many2many:class_skills;";json:"skills"`
}

type DamageType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
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

type Monster struct {
	Id              int     `json:"id;"`
	HitDieNumber    int     `json:"hit_die_number;"`
	HitDieFrequency int     `json:"hit_die_frequency;"`
	Name            string  `json:"name;"`
	ArmorClass      int     `json:"armor_class;"`
	HitPointAverage int     `json:"hit_point_average;"`
	Speed           int     `json:"speed;"`
	Str             int     `json:"str;"`
	Int             int     `json:"int;"`
	Wis             int     `json:"wis;"`
	Dex             int     `json:"dex;"`
	Con             int     `json:"con;"`
	Cha             int     `json:"cha;"`
	Skills          []Skill `gorm:"many2many:monster_skills;"`
	Description     string  `sql:"size:0"`
}

type MonsterSkills struct {
	Id        int
	Monster   Monster
	MonsterId int
	Skill     Skill
	SkillId   int
	Value     int
}

type PlayerCharacterSkill struct {
	Id                int
	PlayerCharacter   PlayerCharacter
	PlayerCharacterId int
	Skill             Skill
	SkillId           int
	Value             int
}

type Proficiency struct {
	Id                int             `json:"id"`
	Name              string          `json:"name"`
	ProficiencyType   ProficiencyType `json:"proficiency_type"`
	ProficiencyTypeId int             `json:"proficiency_type_id"`
}

type ProficiencyType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Race struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	StrMod      int    `json:"str_mod"`
	IntMod      int    `json:"int_mod"`
	WisMod      int    `json:"wis_mod"`
	DexMod      int    `json:"dex_mod"`
	ConMod      int    `json:"con_mod"`
	ChaMod      int    `json:"cha_mod"`
	HitPointMod int    `json:"hit_point_mod"`
}

type Skill struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func Difficulties() map[string]int {
	return map[string]int{
		"Very Easy":         5,
		"Easy":              10,
		"Medium":            15,
		"Hard":              20,
		"Very Hard":         25,
		"Nearly Impossible": 30}
}

type Spell struct {
	Id          int
	Name        string
	SpellType   SpellType
	SpellTypeId int
	Range       int
	V           bool
	S           bool
	M           bool
	Description string
}

type SpellType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SpellSlot struct {
	Id                int `json:"id"`
	PlayerCharacterId int `json:"player_character_id"`
	SpellId           int `json:"spell_id"`
	Level             int `json:"level"`
}
