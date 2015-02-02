package models

import (
	"math/rand"
	"sort"
	"time"
)

func RollCharacter() PlayerCharacter {
	str := makeAbilityRoll()
	intel := makeAbilityRoll()
	wis := makeAbilityRoll()
	dex := makeAbilityRoll()
	con := makeAbilityRoll()
	cha := makeAbilityRoll()
	return PlayerCharacter{
		selectName(),
		str,
		intel,
		wis,
		dex,
		con,
		cha,
		1,
		1,
		1,
		10,
		12,
	}
}

func makeAbilityRoll() int {
	roll1 := random(1, 6)
	roll2 := random(1, 6)
	roll3 := random(1, 6)
	roll4 := random(1, 6)
	return dropLowest([]int{roll1, roll2, roll3, roll4})
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) + 1
}

func dropLowest(rolls []int) int {
	sum := 0
	sort.Ints(rolls)
	for _, roll := range rolls[1:] {
		sum += roll
	}
	return sum
}

func selectName() string {
	names := []string{"Never", "Angus", "Fergus", "Seamus", "Morgan", "James", "Landon"}
	return names[random(1, 6)]
}
