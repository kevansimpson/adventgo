package aoc2015

import (
	"fmt"
	"math"
)

/**
 * <a href="https://adventofcode.com/2015/day/21">Day 21</a>
 */

// import "strings"

type Day21 struct{}

type Boss struct { // also used in Day22
	hitPoints, damage, armor int
}

func (d Day21) fightBosses(boss Boss) (int, int) {
	var empty []newItem
	bestWorst := bestWorstOutfits{
		outfit{empty, newItem{"best", math.MaxInt, 0, 0}}, outfit{empty, newItem{"worst", math.MinInt, 0, 0}}}
	rings := itemShop["rings"]

	for _, weapon := range itemShop["weapons"] {
		for _, armor := range itemShop["armor"] {
			// no rings
			d.evaluateOutfit(d.buyOutfit(weapon, armor), boss, &bestWorst)

			// one ring
			for _, leftHand := range rings {
				d.evaluateOutfit(d.buyOutfit(weapon, armor, leftHand), boss, &bestWorst)

				// two rings
				for _, rightHand := range rings {
					if leftHand.cost != rightHand.cost {
						d.evaluateOutfit(d.buyOutfit(weapon, armor, leftHand, rightHand), boss, &bestWorst)
					}
				}
			}
		}
	}

	return bestWorst.best.stats.cost, bestWorst.worst.stats.cost
}

func (d Day21) evaluateOutfit(of outfit, boss Boss, bw *bestWorstOutfits) {
	if d.isWinningOutfit(of, boss) {
		if of.stats.cost < bw.best.stats.cost {
			bw.best = of
		}
	} else {
		if of.stats.cost > bw.worst.stats.cost {
			bw.worst = of
		}
	}
}

func (d Day21) isWinningOutfit(of outfit, boss Boss) bool {
	myHP, bossHP := 100, boss.hitPoints
	myDmg, bossDmg := of.stats.damage-boss.armor, boss.damage-of.stats.armor

	for myHP > 0 {
		bossHP -= myDmg
		if bossHP <= 0 {
			return true
		}
		myHP -= bossDmg
	}

	return false
}

func (d Day21) buyOutfit(items ...newItem) outfit {
	cost, damage, armor := 0, 0, 0
	name := fmt.Sprintf("%v", items)
	for _, i := range items {
		cost += i.cost
		damage += i.damage
		armor += i.armor
	}
	return outfit{items, newItem{name, cost, damage, armor}}
}

type bestWorstOutfits struct {
	best, worst outfit
}

type outfit struct {
	items []newItem
	stats newItem
}

type newItem struct {
	name                string
	cost, damage, armor int
}

var itemShop = map[string][]newItem{
	"weapons": {
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0}},
	"armor": {
		{"Naked", 0, 0, 0},
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5}},
	"rings": {
		{"Damage +1", 25, 1, 0},
		{"Damage +2", 50, 2, 0},
		{"Damage +3", 100, 3, 0},
		{"Defense +1", 20, 0, 1},
		{"Defense +2", 40, 0, 2},
		{"Defense +3", 80, 0, 3}}}
