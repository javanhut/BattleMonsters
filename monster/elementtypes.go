package monster

import "strings"

type ElementType string

const (
	HighlyEffective   float32     = 4
	Effective         float32     = 2
	Neutral           float32     = 1
	Ineffective       float32     = 0.5
	HighlyIneffective float32     = 0.25
	Fire              ElementType = "FIRE"
	Water             ElementType = "WATER"
	Grass             ElementType = "GRASS"
	Earth             ElementType = "EARTH"
	Light             ElementType = "LIGHT"
	Dark              ElementType = "DARK"
	Divine            ElementType = "DIVINE"
	Magic             ElementType = "MAGIC"
	Wind              ElementType = "WIND"
)

type Element struct {
	name           ElementType
	advantage      []ElementType
	minorAdvantage []ElementType
	minorWeakness  []ElementType
	weakness       []ElementType
}

func containsElement(element ElementType, cmpList []ElementType) bool {
	for _, eleType := range cmpList {
		if strings.EqualFold(string(element), string(eleType)) {
			return true
		}
	}
	return false
}

func (e Element) CalculatedEffectiveness(element ElementType) float32 {
	minorAdv := containsElement(element, e.minorAdvantage)
	minorWeak := containsElement(element, e.minorWeakness)
	advantage := containsElement(element, e.advantage)
	weakness := containsElement(element, e.weakness)

	switch {
	case advantage:
		return HighlyEffective
	case minorAdv:
		return Effective
	case weakness:
		return HighlyIneffective
	case minorWeak:
		return Ineffective
	default:
		return Neutral
	}
}

var FIRE Element = Element{
	name:           Fire,
	advantage:      []ElementType{Grass},
	minorAdvantage: []ElementType{Wind, Dark},
	weakness:       []ElementType{Water, Divine},
	minorWeakness:  []ElementType{Earth, Fire},
}

var WATER Element = Element{
	name:           Water,
	advantage:      []ElementType{Fire},
	minorAdvantage: []ElementType{Earth, Dark},
	weakness:       []ElementType{Grass, Divine},
	minorWeakness:  []ElementType{Wind, Water},
}

var GRASS Element = Element{
	name:           Grass,
	advantage:      []ElementType{Water, Earth},
	minorAdvantage: []ElementType{Light},
	weakness:       []ElementType{Fire, Divine},
	minorWeakness:  []ElementType{Wind, Grass},
}

var EARTH Element = Element{
	name:           Earth,
	advantage:      []ElementType{Wind},
	minorAdvantage: []ElementType{Fire, Light},
	weakness:       []ElementType{Water, Grass},
	minorWeakness:  []ElementType{Magic, Earth},
}

var WIND Element = Element{
	name:           Wind,
	advantage:      []ElementType{Light},
	minorAdvantage: []ElementType{Grass, Water},
	weakness:       []ElementType{Earth},
	minorWeakness:  []ElementType{Fire, Divine, Wind},
}

var DIVINE Element = Element{
	name:           Divine,
	advantage:      []ElementType{Fire, Water, Grass},
	minorAdvantage: []ElementType{Earth, Wind, Light},
	weakness:       []ElementType{Dark, Magic},
	minorWeakness:  []ElementType{Divine},
}

var MAGIC Element = Element{
	name:           Magic,
	advantage:      []ElementType{Divine},
	minorAdvantage: []ElementType{Dark, Earth},
	weakness:       []ElementType{},
	minorWeakness:  []ElementType{Wind, Water, Magic},
}

var LIGHT Element = Element{
	name:           Light,
	advantage:      []ElementType{Dark},
	minorAdvantage: []ElementType{Grass, Magic},
	weakness:       []ElementType{Earth},
	minorWeakness:  []ElementType{Wind, Water, Light},
}

var DARK Element = Element{
	name:           Dark,
	advantage:      []ElementType{Divine, Light},
	minorAdvantage: []ElementType{Grass, Magic},
	weakness:       []ElementType{},
	minorWeakness:  []ElementType{Fire, Earth, Dark},
}
