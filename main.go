package main

import (
	"fmt"

	"battlemon/battle"
	"battlemon/monster"
)

func main() {
	testAttackStat := 10
	testActionPower := 50
	testMonsterLevel := 10
	testOppenentDefenseStat := 60
	effectiveList := []float32{monster.HighlyEffective, monster.Effective, monster.Neutral, monster.Ineffective, monster.HighlyIneffective}
	effectiveNameMap := map[float32]string{
		monster.HighlyEffective:   "Highly Effective",
		monster.Effective:         "Effective",
		monster.Neutral:           "Neutral",
		monster.HighlyIneffective: "Highly Ineffective",
		monster.Ineffective:       "Ineffective",
	}
	for _, effectiveness := range effectiveList {
		damageCalculation := battle.DetermineDamage(testMonsterLevel, testActionPower, testAttackStat, testOppenentDefenseStat, float32(effectiveness))
		fmt.Println("Modifier: ", effectiveNameMap[effectiveness])
		fmt.Println("Calculated Damage: ", damageCalculation)
	}
}
