package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/15">Day 15</a>
 */

import (
	"fmt"
	"strings"

	"github.com/kevansimpson/util"
)

type Day15 struct{}

// === RUN   TestDay15Solutions
// --- PASS: TestDay15Solutions (11.27s)
func (d Day15) findHighestScore(input []string, caloricRequirement int) int {
	fmt.Println()
	ingredientMap := d.readIngredients(input)
	cookbook := Cookbook{0, ingredientMap}
	var ilist []Ingredient
	for _, i := range ingredientMap {
		ilist = append(ilist, i)
	}
	recipe := Recipe{caloricRequirement, make(map[string]int)}

	cookbook.buildAllRecipes(ilist, 0, recipe, 100)
	return cookbook.highestScore
}

type Cookbook struct {
	highestScore int
	pages        map[string]Ingredient
}

type Ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

type Recipe struct {
	caloricRequirement int
	ingredients        map[string]int
}

func (c *Cookbook) buildAllRecipes(ingredientList []Ingredient, ingredientIndex int, recipe Recipe, total int) {
	if len(ingredientList) <= ingredientIndex {
		if recipe.sumTeaspoons() == 100 {
			if recipe.caloricRequirement <= 0 || c.caloricCount(recipe) == recipe.caloricRequirement {
				score := c.score(recipe)
				if score > c.highestScore {
					c.highestScore = score
					fmt.Print(".")
				}
			}
		}

		return // recipe has all ingredients
	}

	for i := 0; i <= total; i++ {
		recipe.setTeaspoons(ingredientList[ingredientIndex].name, i)
		c.buildAllRecipes(ingredientList, ingredientIndex+1, recipe, total-1)
	}
}

func (c Cookbook) caloricCount(r Recipe) int {
	count := 0
	for name, ingredient := range c.pages {
		ts, _ := r.ingredients[name]
		count += ts * ingredient.calories
	}
	return count
}

func (c Cookbook) score(r Recipe) int {
	cap, d, f, t := 0, 0, 0, 0
	for name, ingredient := range c.pages {
		ts, _ := r.ingredients[name]
		cap += ts * ingredient.capacity
		d += ts * ingredient.durability
		f += ts * ingredient.flavor
		t += ts * ingredient.texture
	}

	if cap <= 0 || d <= 0 || f <= 0 || t <= 0 {
		return 0
	}
	score := cap * d * f * t
	return score
}

func (r *Recipe) setTeaspoons(ingredient string, teaspoons int) {
	r.ingredients[ingredient] = teaspoons
}

func (r Recipe) sumTeaspoons() int {
	ts := 0
	for _, v := range r.ingredients {
		ts += v
	}
	return ts
}

func (d Day15) readIngredients(input []string) map[string]Ingredient {
	ingredientMap := make(map[string]Ingredient)
	for _, str := range input {
		name := str[:strings.Index(str, ":")]
		nums := util.ExtractInts(str)
		ingredientMap[name] = Ingredient{name, nums[0], nums[1], nums[2], nums[3], nums[4]}
	}

	return ingredientMap
}
