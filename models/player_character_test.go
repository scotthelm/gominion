package models

import (
	"fmt"
	"testing"
)

func TestCharacterName(t *testing.T) {
	pc := PlayerCharacter{"Bob", 9, 9, 9, 9, 9, 9, 1, 1, 1, 10, 12}
	if pc.Name != "Bob" {
		t.Errorf("Name was not set. Expecting 'Bob', got: '%s'", pc.Name)
	}
}

func TestCharacterClass(t *testing.T) {
	pc := PlayerCharacter{"Jeff", 9, 9, 9, 9, 9, 9, 1, 1, 1, 10, 12}
	if pc.Race().Name != "Dwarf" {
		t.Errorf("Type of class wrong. Expecting Dwarf, got: '%s'", pc.Race().Name)
	}
}

func TestCharacterAttributes(t *testing.T) {
	pc := PlayerCharacter{"Jimmy", 9, 9, 9, 9, 9, 9, 1, 1, 1, 10, 12}
	if pc.Str() != 10 {
		t.Errorf("Str should be 10, got %d", pc.Str())
	}
	if pc.Int() != 9 {
		t.Errorf("Int should be 9, got %d", pc.Int())
	}
	if pc.Wis() != 10 {
		t.Errorf("Wis should be 10, got %d", pc.Wis())
	}
	if pc.Dex() != 9 {
		t.Errorf("Dex should be 9, got %d", pc.Dex())
	}
	if pc.Con() != 10 {
		t.Errorf("Con should be 10, got %d", pc.Con())
	}
	if pc.Cha() != 9 {
		t.Errorf("Cha should be 9, got %d", pc.Cha())
	}
}

func TestPCRoll(t *testing.T) {
	pc := RollCharacter()
	if pc.RolledStr < 3 || pc.RolledStr > 18 {
		t.Errorf("Str should be between 3 and 18, got %d", pc.RolledStr)
	}
}

func TestCreateSql(t *testing.T) {
	pc := RollCharacter()
	insert_sql := CreateSql(pc)
	fmt.Println(insert_sql)
}
