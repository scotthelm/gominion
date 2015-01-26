package models

import (
	"testing"
)

func TestCharacterName(t *testing.T) {
	pc := PlayerCharacter{"Bob", 9, 9, 9, 9, 9, 9, 1, 1, 1, 10, 12}
	if pc.Name != "Bob" {
		t.Errorf("Name was not set. Expecting 'Bob', got: '%s'", pc.Name)
	}
}

func TestCharacterAttributes(t *testing.T) {
	pc := PlayerCharacter{"Jimmy", 9, 9, 9, 9, 9, 9, 1, 1, 1, 10, 12}
	if pc.Str != 9 {
		t.Errorf("Str should be 9, got %d", pc.Str)
	}
	if pc.Int != 9 {
		t.Errorf("Int should be 9, got %d", pc.Int)
	}
	if pc.Wis != 9 {
		t.Errorf("Wis should be 9, got %d", pc.Wis)
	}
	if pc.Dex != 9 {
		t.Errorf("Dex should be 9, got %d", pc.Dex)
	}
	if pc.Con != 9 {
		t.Errorf("Con should be 9, got %d", pc.Con)
	}
	if pc.Cha != 9 {
		t.Errorf("Cha should be 9, got %d", pc.Cha)
	}
}

func TestPCRoll(t *testing.T) {
	pc := RollCharacter()
	if pc.Str < 3 || pc.Str > 18 {
		t.Errorf("Str should be between 3 and 18, got %d", pc.Str)
	}
}
