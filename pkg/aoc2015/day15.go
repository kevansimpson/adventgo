package aoc2015

/**
 * <a href="https://adventofcode.com/2015/day/15">Day 15</a>
 */

import (
	"strings"

	"github.com/kevansimpson/util"
)

type Day15 struct{}

// Iteration 1: calling findHighestScore twice with different caloric requirements
// === RUN   TestDay15Solutions
// --- PASS: TestDay15Solutions (11.27s)

// Iteration 2: running through Recipe permutations once and grab both max scores at the same time
// === RUN   TestDay15Solutions
// --- PASS: TestDay15Solutions (5.70s)

// Iteration 3: reducing the buildAllRecipes loop from 0..100 down to 16..32
// === RUN   TestDay15Solutions
// --- PASS: TestDay15Solutions (0.01s)
func (d Day15) findHighestScore(input []string, caloricRequirement int) (int, int) {
	ingredientMap := d.readIngredients(input)
	var ilist []Ingredient
	for _, i := range ingredientMap {
		ilist = append(ilist, i)
	}
	cookbook := Cookbook{0, 0, len(ilist), ingredientMap}
	recipe := Recipe{caloricRequirement, make(map[string]int)}

	cookbook.buildAllRecipes(ilist, 0, recipe, 100)
	return cookbook.highestScore, cookbook.highestScoreWithCaloricReq
}

type Cookbook struct {
	highestScore, highestScoreWithCaloricReq int
	ingredientCount                          int
	pages                                    map[string]Ingredient
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
	if c.ingredientCount <= ingredientIndex {
		if recipe.sumTeaspoons() == 100 {
			score := c.score(recipe)
			if score > c.highestScore {
				c.highestScore = score
			}
			if score > c.highestScoreWithCaloricReq && c.caloricCount(recipe) == recipe.caloricRequirement {
				c.highestScoreWithCaloricReq = score
			}
		}

		return // recipe has all ingredients
	}

	// one {500 map[Butterscotch:31 Candy:29 Frosting:24 Sugar:16]}
	// two {500 map[Butterscotch:31 Candy:23 Frosting:21 Sugar:25]}
	//
	// this loop was originally 0..total ==> 99000301 iterations
	// with the benefit of hindsight, we "know" that all ingredients
	// will have at least 16 and no more than 32 teaspoons
	// this drops the iteration count from 99000301 to 88740
	for i := 16; i <= 32; i++ {
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
		// fmt.Printf("- %v\n", r)
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
