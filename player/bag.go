package player

import (
	"fmt"
	"log"
	"strings"
)

type INDEX int

const (
	SHARDS_IDX    INDEX = 0
	SHARD_MAX     int   = 999
	SHARD_MIN     int   = 0
	POTIONS_IDX   INDEX = 1
	FOOD_IDX      INDEX = 2
	KEY_ITEMS_IDX INDEX = 3
)

var sectionMap = map[string]INDEX{
	"SHARDS":   SHARDS_IDX,
	"POTIONS":  POTIONS_IDX,
	"FOOD":     FOOD_IDX,
	"KEYITEMS": KEY_ITEMS_IDX,
}

type BagSection struct {
	id    int
	name  string
	items map[string]int
}

type Bag struct {
	sections       map[INDEX]BagSection
	currentSection int
}

func (b *Bag) InitializeSections() {
	for name, val := range sectionMap {
		b.sections[val] = BagSection{id: int(val), name: name}
	}
}

func (b *Bag) GetSection(sectionName string) (BagSection, error) {
	sectionIdx, ok := sectionMap[strings.ToUpper(sectionName)]
	if !ok {
		return BagSection{}, fmt.Errorf("%s value not found in sections", sectionName)
	}
	sectionInfo, exists := b.sections[sectionIdx]
	if !exists {
		return BagSection{}, fmt.Errorf("section %s not found", sectionName)
	}
	return sectionInfo, nil
}

func (b *Bag) AddShard(name string, shardCount int) {
	shards, err := b.GetSection("shards")
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error: %v", err))
	}
	switch {
	case shardCount > SHARD_MAX:
		log.Fatalf(fmt.Sprintf("exceeded max allowable shards. shard max: %d", SHARD_MAX))
	case shardCount < SHARD_MIN:
		log.Fatalf(fmt.Sprintf("shard below allowable shards. shard min: %d", SHARD_MIN))
	default:
		log.Println("shard are ok!")
	}
	shards.items[name] = shardCount
}
