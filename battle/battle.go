package battle

func DetermineDamage(level int, actionPower int, attackStat int, opponentDefenseStat int, modifierStat float32) float32 {
	basicFormula := (((2*float32(level)/5)+2)*float32(actionPower)*(float32(attackStat)/float32(opponentDefenseStat))/float32(50) + 2) * modifierStat
	return basicFormula
}
