package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Food struct {
	Ingredients []string
	Allergens   []string
}

func main() {
	foods, err := readIngredients("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	validAllergenIngredients, count := partOne(foods)
	fmt.Println(count)
	fmt.Println(partTwo(foods, validAllergenIngredients))
}

func partOne(foods []Food) (map[string]map[string]bool, int) {
	allergenCounts := map[string]int{}
	allergenIngredientCounts := map[string]map[string]int{}
	for _, food := range foods {
		for _, allergen := range food.Allergens {
			if _, ok := allergenIngredientCounts[allergen]; !ok {
				allergenIngredientCounts[allergen] = map[string]int{}
			}
			allergenCounts[allergen]++
			for _, ingredient := range food.Ingredients {
				allergenIngredientCounts[allergen][ingredient]++
			}
		}
	}
	count := 0
	validAllergenIngredients := map[string]map[string]bool{}
	for _, food := range foods {
	CheckIngredients:
		for _, ingredient := range food.Ingredients {
			for allergen, allergenCount := range allergenCounts {
				if allergenIngredientCounts[allergen][ingredient] == allergenCount {
					if _, ok := validAllergenIngredients[allergen]; !ok {
						validAllergenIngredients[allergen] = map[string]bool{}
					}
					validAllergenIngredients[allergen][ingredient] = true
					continue CheckIngredients
				}
			}
			count++
		}
	}
	return validAllergenIngredients, count
}

func partTwo(foods []Food, validAllergenIngredients map[string]map[string]bool) string {
	correctAllergenForIngredient := map[string]string{}
	ingredients := []string{}
	for len(validAllergenIngredients) > 0 {
		allergenToDelete, ingredientToDelete := "", ""
		for allergen, ingredientMap := range validAllergenIngredients {
			if len(ingredientMap) == 1 {
				for ingredient := range ingredientMap {
					correctAllergenForIngredient[ingredient] = allergen
					ingredientToDelete = ingredient
				}
				allergenToDelete = allergen
				break
			}
		}
		ingredients = append(ingredients, ingredientToDelete)
		for _, ingredientMap := range validAllergenIngredients {
			delete(ingredientMap, ingredientToDelete)
		}
		delete(validAllergenIngredients, allergenToDelete)
	}
	sort.Slice(ingredients, func(i, j int) bool {
		return correctAllergenForIngredient[ingredients[i]] < correctAllergenForIngredient[ingredients[j]]
	})
	return strings.Join(ingredients, ",")
}

func readIngredients(fileName string) ([]Food, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	foods := []Food{}
	for scanner.Scan() {
		splitFood := strings.Split(scanner.Text(), " (contains ")
		foods = append(foods, Food{Ingredients: strings.Split(splitFood[0], " "), Allergens: strings.Split(strings.TrimSuffix(splitFood[1], ")"), ", ")})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return foods, nil
}
