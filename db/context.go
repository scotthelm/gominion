package db

import (
	"github.com/jinzhu/gorm"
	// importing the sql driver for gorm
	_ "github.com/mattn/go-sqlite3"
	"github.com/scotthelm/gominion/models"
)

// Context the struct is where database connection information is held during
// program execution
type Context struct {
	Db              gorm.DB
	DriverType      string
	ConnctionString string
}

// NewContext creates a new context with the given settings
func NewContext(driverType string, connectionString string) Context {
	db, _ := gorm.Open(driverType, connectionString)
	db.LogMode(true)
	return Context{db, driverType, connectionString}
}

// DBSync sets the driver type to sqlite3
func (c *Context) DBSync() {
	c.DriverType = "sqlite3"
}

// Migrate uses gorm to migrate the database based on the struct models provided
func (c *Context) Migrate() {
	c.Db.AutoMigrate(
		&models.Skill{},
		&models.Class{},
		&models.Proficiency{},
		&models.ProficiencyType{},
		&models.Equipment{},
		&models.DamageType{},
		&models.Spell{},
		&models.SpellType{},
		&models.Campaign{},
		&models.Monster{},
		&models.PlayerCharacter{},
		&models.Race{},
	)
}
