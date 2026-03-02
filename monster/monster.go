package monster

import "fmt"

type ActionStats struct {
	baseAttack     int
	baseDefense    int
	baseSpeed      int
	specialAttack  int
	specialDefense int
}

type Stats struct {
	level         int
	hitPoints     int
	iconFrontId   int
	iconBackId    int
	maxExperience int
}

type Action struct {
	magicPoints int
	description string
	effects     []string
	actionType  Element
	baseLevel   int
}

type Monster struct {
	name       string
	info       Stats
	traits     []Element
	actionList []Action
}

func (m *Monster) changeName(newName string) {
	m.name = newName
}

func (m *Monster) printInfo() {
	fmt.Println(m.info)
}
