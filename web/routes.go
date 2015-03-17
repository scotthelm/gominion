package web

import (
	"net/http"

	m "github.com/scotthelm/gominion/models"
	// "reflect"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"CampaignIndex", "GET", "/api/campaigns", IndexHandler(m.Campaign{})},
	Route{"CampaignShow", "GET", "/api/campaigns/{id}", ShowHandler(m.Campaign{})},
	Route{"CampaignCreate", "POST", "/api/campaigns", CreateHandler(m.Campaign{})},
	Route{"CampaignDelete", "DELETE", "/api/campaigns/{id}", DeleteHandler(m.Campaign{})},
	Route{"CampaignUpdate", "PUT", "/api/campaigns/{id}", UpdateHandler(m.Campaign{})},

	Route{"RaceIndex", "GET", "/api/races", IndexHandler(m.Race{})},
	Route{"RaceShow", "GET", "/api/races/{id}", ShowHandler(m.Race{})},
	Route{"RaceCreate", "POST", "/api/races", CreateHandler(m.Race{})},
	Route{"RaceDelete", "DELETE", "/api/races/{id}", DeleteHandler(m.Race{})},
	Route{"RaceUpdate", "PUT", "/api/races/{id}", UpdateHandler(m.Race{})},

	Route{"SpellTypeIndex", "GET", "/api/spell_types", IndexHandler(m.SpellType{})},
	Route{"SpellTypeShow", "GET", "/api/spell_types/{id}", ShowHandler(m.SpellType{})},
	Route{"SpellTypeCreate", "POST", "/api/spell_types", CreateHandler(m.SpellType{})},
	Route{"SpellTypeDelete", "DELETE", "/api/spell_types/{id}", DeleteHandler(m.SpellType{})},
	Route{"SpellTypeUpdate", "PUT", "/api/spell_types/{id}", UpdateHandler(m.SpellType{})},

	Route{"ProficiencyTypeIndex", "GET", "/api/proficiency_types", IndexHandler(m.ProficiencyType{})},
	Route{"ProficiencyTypeShow", "GET", "/api/proficiency_types/{id}", ShowHandler(m.ProficiencyType{})},
	Route{"ProficiencyTypeCreate", "POST", "/api/proficiency_types", CreateHandler(m.ProficiencyType{})},
	Route{"ProficiencyTypeDelete", "DELETE", "/api/proficiency_types/{id}", DeleteHandler(m.ProficiencyType{})},
	Route{"ProficiencyTypeUpdate", "PUT", "/api/proficiency_types/{id}", UpdateHandler(m.ProficiencyType{})},

	Route{"ProficiencyIndex", "GET", "/api/proficiencies", IndexHandler(m.Proficiency{}, "ProficiencyType")},
	Route{"ProficiencyShow", "GET", "/api/proficiencies/{id}", ShowHandler(m.Proficiency{})},
	Route{"ProficiencyCreate", "POST", "/api/proficiencies", CreateHandler(m.Proficiency{})},
	Route{"ProficiencyDelete", "DELETE", "/api/proficiencies/{id}", DeleteHandler(m.Proficiency{})},
	Route{"ProficiencyUpdate", "PUT", "/api/proficiencies/{id}", UpdateHandler(m.Proficiency{})},

	Route{"SkillIndex", "GET", "/api/skills", IndexHandler(m.Skill{})},
	Route{"SkillShow", "GET", "/api/skills/{id}", ShowHandler(m.Skill{})},
	Route{"SkillCreate", "POST", "/api/skills", CreateHandler(m.Skill{})},
	Route{"SkillDelete", "DELETE", "/api/skills/{id}", DeleteHandler(m.Skill{})},
	Route{"SkillUpdate", "PUT", "/api/skills/{id}", UpdateHandler(m.Skill{})},

	Route{"ClassIndex", "GET", "/api/classes", IndexHandler(m.Class{}, "Proficiencies", "Skills")},
	Route{"ClassShow", "GET", "/api/classes/{id}", ShowHandler(m.Class{}, "Proficiencies", "Skills")},
	Route{"ClassCreate", "POST", "/api/classes", CreateHandler(m.Class{})},
	Route{"ClassDelete", "DELETE", "/api/classes/{id}", DeleteHandler(m.Class{})},
	Route{"ClassUpdate", "PUT", "/api/classes/{id}", UpdateHandler(m.Class{})},
}
