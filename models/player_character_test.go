package models

import (
	"testing"
)

func TestAbilityCheck(t *testing.T) {
	pc := testCharacter("Jeff", map[string]int{"str": 12, "int": 10, "wis": 15, "dex": 10, "con": 10, "cha": 10})
	if pc.StrCheck(12, 11) == false {
		t.Errorf("StrengthCheck should be true, got false")
	}
	if pc.IntCheck(12, 11) == true {
		t.Errorf("IntCheck should be false, got true")
	}
	if pc.WisCheck(12, 11) == false {
		t.Errorf("WisCheck should be true, got false")
	}
	if pc.DexCheck(12, 11) == true {
		t.Errorf("DexCheck should be false, got true")
	}
	if pc.ConCheck(12, 11) == true {
		t.Errorf("ConCheck should be false, got true")
	}
	if pc.ChaCheck(12, 11) == true {
		t.Errorf("ChaCheck should be false, got true")
	}
}

func TestProficiency(t *testing.T) {
	pc := testCharacter("Jeff", map[string]int{"str": 12, "int": 10, "wis": 15, "dex": 10, "con": 10, "cha": 10})
	if pc.HasProficiency("All Weapons") == false {
		t.Errorf("Expected All Weapons HasProficiency to be true, got false")
	}
}

func testCharacter(name string, abilities map[string]int) PlayerCharacter {
	return PlayerCharacter{
		1,
		name,
		abilities["str"],
		abilities["int"],
		abilities["wis"],
		abilities["dex"],
		abilities["con"],
		abilities["cha"],
		1,
		10,
		12,
		Class{},
		1,
		Race{},
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		nil,
		nil,
		nil,
		[]Proficiency{Proficiency{1, "All Weapons", ProficiencyType{1, "Weapon"}, 1}}}
}
