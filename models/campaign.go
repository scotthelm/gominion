package models

type Campaign struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Campaigns []Campaign

type Monster struct {
	Id              int
	HitDieNumber    int
	HitDieFrequency int
	Name            string
	ArmorClass      int
	HitPointAverage int
	Speed           int
	Str             int
	Int             int
	Wis             int
	Dex             int
	Con             int
	Cha             int
	Skills          []MonsterSkills
	Description     string `sql:"size:0"`
}

type MonsterSkills struct {
	Id        int
	Monster   Monster
	MonsterId int
	Skill     Skill
	SkillId   int
	Value     int
}
