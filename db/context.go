package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/scotthelm/gominion/models"
)

type Context struct {
	Db              gorm.DB
	DriverType      string
	ConnctionString string
}

func NewContext(driverType string, connectionString string) Context {
	db, _ := gorm.Open(driverType, connectionString)
	db.LogMode(true)
	return Context{db, driverType, connectionString}
}
func (c *Context) DBSync() {
	c.DriverType = "sqlite3"
}

func (c *Context) Migrate() {
	c.Db.AutoMigrate(
		&models.Skill{},
		&models.Class{},
		&models.Proficiency{},
		&models.ProficiencyType{},
		&models.PlayerCharacterSkill{},
		&models.Equipment{},
		&models.DamageType{},
		&models.Spell{},
		&models.SpellType{},
		&models.Campaign{},
		&models.Monster{},
		&models.MonsterSkills{},
		&models.PlayerCharacter{},
	)
}
