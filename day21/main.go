// Advent of Code 2020, Day 21
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Day 21: Allergen Assessment
// Part 1 answer: 2389
// Part 2 answer: fsr,skrxt,lqbcg,mgbv,dvjrrkv,ndnlm,xcljh,zbhp
func main() {
	fmt.Println("Advent of Code 2020, Day 21")
	const filename = "input.txt"
	fmt.Printf("Reading file %s\n", filename)
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	foods := readFoods(input)
	part1, part2 := part1And2(foods)
	fmt.Printf("Part 1. Answer = %d\n", part1)
	fmt.Printf("Part 2. Answer = %s\n", part2)
}

type food struct {
	ingredients []string
	allergens   []string
}

func (f food) containsIngredient(ingredient string) bool {
	for _, i := range f.ingredients {
		if i == ingredient {
			return true
		}
	}
	return false
}

func (f food) containsAllergen(allergen string) bool {
	for _, a := range f.allergens {
		if a == allergen {
			return true
		}
	}
	return false
}

type ingredientSet map[string]bool

// group 1 is space separated list of ingredients, group 2 is comma separated list of allergens
var foodRegexp = regexp.MustCompile(`^(.*) \(contains (.*)\)$`)

func readFoods(r io.Reader) []food {
	var foods []food
	input := bufio.NewScanner(r)
	for input.Scan() {
		line := input.Text()
		m := foodRegexp.FindStringSubmatch(line)
		foods = append(foods, food{strings.Split(m[1], " "), strings.Split(m[2], ", ")})
	}
	return foods
}

func part1And2(foods []food) (int, string) {
	// Per allergen, gather all the ingredients (a set) that might have it
	allergenToPotentialIngredients := make(map[string]ingredientSet)
	for _, f := range foods {
		// For each allergen...
		for _, a := range f.allergens {
			is := allergenToPotentialIngredients[a]
			if is == nil {
				is = make(ingredientSet)
				allergenToPotentialIngredients[a] = is
			}
			// ... mark each ingredient as a possibility
			for _, ingred := range f.ingredients {
				is[ingred] = true
			}
		}
	}
	s, _ := searchAllergens(foods, allergenToPotentialIngredients, newState())
	// Count the number of ingredients that cannot contain allergens
	var count int
	for _, f := range foods {
		for _, ingred := range f.ingredients {
			if _, ok := s.ingredientToAllergen[ingred]; !ok {
				count++
			}
		}
	}
	return count, strings.Join(ingredientsInAllergenOrder(s.allergenToIngredient), ",")
}

func ingredientsInAllergenOrder(allergenToIngredient map[string]string) []string {
	allergens := make([]string, 0, len(allergenToIngredient))
	for a := range allergenToIngredient {
		allergens = append(allergens, a)
	}
	sort.Strings(allergens)
	ingredients := make([]string, 0, len(allergens))
	for _, a := range allergens {
		ingredients = append(ingredients, allergenToIngredient[a])
	}
	return ingredients
}

type state struct {
	allergenToIngredient map[string]string
	ingredientToAllergen map[string]string
}

func newState() state {
	var s state
	s.allergenToIngredient = make(map[string]string)
	s.ingredientToAllergen = make(map[string]string)
	return s
}

func copyState(other state) state {
	var s state
	s.allergenToIngredient = make(map[string]string)
	for k, v := range other.allergenToIngredient {
		s.allergenToIngredient[k] = v
	}
	s.ingredientToAllergen = make(map[string]string)
	for k, v := range other.ingredientToAllergen {
		s.ingredientToAllergen[k] = v
	}
	return s
}

// If returns true, then we found a state that fits the criteria
func searchAllergens(foods []food, allergenMap map[string]ingredientSet, s state) (state, bool) {
	for a, is := range allergenMap {
		// Allergen is only in one ingredient
		if _, ok := s.allergenToIngredient[a]; ok {
			continue
		}
		// Okay we haven't assigned this allergen yet
		for ingred := range is {
			// Ingredient can only have at most one allergen
			if _, ok := s.ingredientToAllergen[ingred]; ok {
				continue
			}
			// Do a quick check to see if this invalidates any foods
			if !checkFoods(foods, a, ingred) {
				continue
			}
			s2 := copyState(s)
			s2.allergenToIngredient[a] = ingred
			s2.ingredientToAllergen[ingred] = a
			if s3, ok := searchAllergens(foods, allergenMap, s2); ok {
				return s3, ok
			}
		}
	}
	// Now check all the foods to see if those rules work
	for _, f := range foods {
		if !f.isValid(s) {
			return s, false
		}
	}
	return s, true
}

func checkFoods(foods []food, allergen, ingredient string) bool {
	for _, f := range foods {
		if f.containsAllergen(allergen) && !f.containsIngredient(ingredient) {
			return false
		}
	}
	return true
}

func (f food) isValid(s state) bool {
AllergenLoop:
	for _, a := range f.allergens {
		for _, ingred := range f.ingredients {
			if s.ingredientToAllergen[ingred] == a {
				continue AllergenLoop
			}
		}
		// didn't find it
		return false
	}
	return true
}
