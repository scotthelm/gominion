package gominion

import "time"

type Race struct {
	Name        string `db:"name"`
	Subrace     string `db:"subrace"`
	StrMod      int    `db:"str_mod"`
	IntMod      int    `db:"int_mod"`
	WisMod      int    `db:"wis_mod"`
	DexMod      int    `db:"dex_mod"`
	ConMod      int    `db:"con_mod"`
	ChaMod      int    `db:"cha_mod"`
	HitPointMod int    `db:"hit_point_mod"`
	Entity
}

func (r *Race) Create() (*Race, error) {
	return r, nil
}

func (r *Race) Update() (*Race, error) {
	return r, nil
}

func (r *Race) Delete() (bool, error) {
	return false, nil
}

func DwarfWithSub(subrace string) Race {
	switch subrace {
	case "Hill":
		return Race{"Dwarf", "Hill", 0, 0, 1, 0, 2, 0, 1, Entity{0, time.Now()}}
	case "Mountain":
		return Race{"Dwarf", "Mountain", 2, 0, 0, 0, 2, 0, 1, Entity{0, time.Now()}}
	}
	return Race{"Dwarf", "Plains", 1, 0, 1, 0, 1, 0, 1, Entity{0, time.Now()}}
}
func Dwarf() Race {
	return DwarfWithSub("Plains")
}
