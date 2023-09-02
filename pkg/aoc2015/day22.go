package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/22">Day 22</a>
 */

import (
	"container/heap"
	"fmt"
	"sync"

	"github.com/kevansimpson/util"
)

type Day22 struct{}

func (d Day22) wizardsFightBoss(boss Boss) (int, int) {
	var wg sync.WaitGroup
	manaChannel := make(chan int)
	hardChannel := make(chan int)
	wg.Add(2)

	go d.wizardVsBoss(boss, 0, manaChannel, &wg)
	go d.wizardVsBoss(boss, 1, hardChannel, &wg)

	leastMana0, leastMana1 := <-manaChannel, <-hardChannel

	close(manaChannel)
	close(hardChannel)
	wg.Wait()

	return leastMana0, leastMana1
}

func (d Day22) wizardVsBoss(boss Boss, extraDamage int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	lowestManaSpent := 2000
	// https://pkg.go.dev/container/heap#:~:text=Example%20(-,PriorityQueue,-)%20%C2%B6
	// implementation refactored to be generic - resides in util/queue.go
	// anecdotal evidence: takes about 4 seconds less than manually sorting or
	//   								finding the highest priority in a slice
	pq := make(util.PriorityQueue[magicalDuel], 1)
	pq[0] = &util.PQItem[magicalDuel]{
		Value:    magicalDuel{wizard{50, 500, 0}, 0, make(map[string]int), boss.hitPoints},
		Priority: 0,
		Index:    0}

	for pq.Len() > 0 {
		pqItem := heap.Pop(&pq).(*util.PQItem[magicalDuel])
		duel := pqItem.Value

		duel.wiz.hitPoints -= extraDamage
		if duel.wiz.hitPoints <= 0 || duel.totalManaSpent > lowestManaSpent {
			continue
		}

		duel.fightRound()

		for _, spell := range spellBook {
			if duel.canCast(spell) {
				copy := d.battleVariation(&duel)
				copy.cast(spell)
				copy.fightRound()

				if copy.bossHP <= 0 {
					if copy.totalManaSpent < lowestManaSpent {
						lowestManaSpent = copy.totalManaSpent
					}
				} else {
					wound := boss.damage - copy.wiz.armor
					if wound < 1 {
						copy.wiz.hitPoints -= 1
					} else {
						copy.wiz.hitPoints -= wound
					}

					if copy.wiz.hitPoints > 0 && copy.wiz.mana > 0 && copy.totalManaSpent < lowestManaSpent {
						variant := &util.PQItem[magicalDuel]{
							Value:    copy,
							Priority: 1,
						}
						heap.Push(&pq, variant)
					}
				}
			}
		}
	}

	ch <- lowestManaSpent
}

func (d Day22) battleVariation(duel *magicalDuel) magicalDuel {
	doppleganger := &duel.wiz
	variation := make(map[string]int, len(duel.activeEffects))
	for k, v := range duel.activeEffects {
		variation[k] = v
	}
	return magicalDuel{*doppleganger, duel.totalManaSpent, variation, duel.bossHP}
}

func (duel *magicalDuel) fightRound() {
	shield, hasShield := duel.activeEffects["Shield"]
	if !hasShield || shield == 0 {
		duel.wiz.armor = 0
	}

	for spell, dur := range duel.activeEffects {
		if dur > 0 {
			duel.activeEffects[spell] = dur - 1
			switch spell {
			case "Shield":
				duel.wiz.armor = 7
			case "Poison":
				duel.bossHP -= 3
			case "Recharge":
				duel.wiz.mana += 101
			}
		}
	}
}

func (duel *magicalDuel) cast(spell magicSpell) {
	duel.wiz.mana -= spell.manaCost
	duel.totalManaSpent += spell.manaCost

	switch spell.name {
	case "MagicMissile":
		duel.bossHP -= 4
	case "Drain":
		duel.bossHP -= 2
		duel.wiz.hitPoints += 2
	default:
		duel.activeEffects[spell.name] = spell.effectLasts
	}
}

func (duel magicalDuel) canCast(spell magicSpell) bool {
	return duel.wiz.mana >= spell.manaCost &&
		(spell.name == "MagicMissile" || spell.name == "Drain" || duel.getDuration(spell.name) == 0)
}

func (duel magicalDuel) getDuration(spell string) int {
	dur, has := duel.activeEffects[spell]
	if has {
		return dur
	} else {
		return 0
	}
}

var spellBook = []magicSpell{
	{"MagicMissile", 53, 0}, // 4 dmg
	{"Drain", 73, 0},        // 2 dmg + 2 HP
	{"Shield", 113, 6},      // +7 armor / turn
	{"Poison", 173, 6},      // 3 dmg / turn
	{"Recharge", 229, 5}}    // +101 mana / turn

type magicSpell struct {
	name                  string
	manaCost, effectLasts int
}

type magicalDuel struct {
	wiz            wizard
	totalManaSpent int
	activeEffects  map[string]int
	bossHP         int
}

type wizard struct {
	hitPoints, mana, armor int
}

func (w wizard) String() string {
	return fmt.Sprintf("Wizard(hp=%d,mana=%d,armor=%d)", w.hitPoints, w.mana, w.armor)
}
