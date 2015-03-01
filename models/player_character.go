package models

type PlayerCharacter struct {
	Id                 int           `json:"id"`
	Name               string        `json:"name"`
	Str                int           `json:"str"`
	Int                int           `json:"int"`
	Wis                int           `json:"wis"`
	Dex                int           `json:"dex"`
	Con                int           `json:"con"`
	Cha                int           `json:"cha"`
	Level              int           `json:"level"`
	MaxHitPoints       int           `json:"max_hit_points"`
	BaseArmorClass     int           `json:"base_armor_class"`
	Class              Class         `json:"class"`
	ClassId            int           `json:"class_id"`
	Race               Race          `json:"race"`
	RaceId             int           `json:"race_id"`
	ProficiencyBonus   int           `json:"proficiency_bonus"`
	ArmorClass         int           `json:"armor_class"`
	InitiativeBonus    int           `json:"initiative_bonus"`
	SpeedInFeet        int           `json:"speed_in_feet"`
	CurrentHitPoints   int           `json:"current_hit_points"`
	TemporaryHitPoints int           `json:"temporary_hit_points"`
	Skills             []Skill       `gorm:"many2many:player_character_skills";json:"skills"`
	Spells             []Spell       `gorm:"many2many:player_character_spells";json:"spells"`
	Equipment          []Equipment   `gorm:"many2many:player_character_equipment";json:"equipment"`
	Proficiencies      []Proficiency `gorm:"many2many:player_characer_proficiencies";json:"proficiencies"`
}

func (p *PlayerCharacter) abilityCheck(ability Ability, target int, roll int, skill ...Skill) bool {
	modifier := 0
	for _, s := range skill {
		if p.HasSkill(s) {
			modifier = p.ProficiencyBonus
			break
		}
	}
	return ability.Check(target, roll, modifier)
}

func (p *PlayerCharacter) StrCheck(target int, roll int, skill ...Skill) bool {
	return p.abilityCheck(Ability{"Str", p.Str}, target, roll, skill...)
}
func (p *PlayerCharacter) IntCheck(target int, roll int, skill ...Skill) bool {
	return p.abilityCheck(Ability{"Int", p.Int}, target, roll, skill...)
}

func (p *PlayerCharacter) WisCheck(target int, roll int, skill ...Skill) bool {
	return p.abilityCheck(Ability{"Wis", p.Wis}, target, roll, skill...)
}

func (p *PlayerCharacter) DexCheck(target int, roll int, skill ...Skill) bool {
	return p.abilityCheck(Ability{"Dex", p.Dex}, target, roll, skill...)
}

func (p *PlayerCharacter) ConCheck(target int, roll int, skill ...Skill) bool {
	return p.abilityCheck(Ability{"Con", p.Con}, target, roll, skill...)
}

func (p *PlayerCharacter) ChaCheck(target int, roll int, skill ...Skill) bool {
	return p.abilityCheck(Ability{"Cha", p.Cha}, target, roll, skill...)
}

func (p *PlayerCharacter) HasProficiency(prof string) bool {
	out := false
	for _, proficiency := range p.Proficiencies {
		if proficiency.Name == prof {
			out = true
			break
		}
	}
	return out
}

func (p *PlayerCharacter) HasSkill(skill Skill) bool {
	hasSkill := false
	for _, s := range p.Skills {
		if s.Id == skill.Id {
			hasSkill = true
			break
		}
	}
	return hasSkill
}
