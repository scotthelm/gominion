package models

import ()

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
	Id        int
	Name      string
	Attribute int
}
