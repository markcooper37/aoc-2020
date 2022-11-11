package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	bags, err := readBags("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(bags))
	fmt.Println(partTwo(bags))
}

func partOne(bags map[string]map[string]int) int {
	total := 0
	for bag := range bags {
		if containsShinyGold(bag, bags) {
			total++
		}
	}
	return total
}

func containsShinyGold(containingBag string, bags map[string]map[string]int) bool {
	for containedBag := range bags[containingBag] {
		if containedBag == "shiny gold" {
			return true
		}
		if containsShinyGold(containedBag, bags) {
			return true
		}
	}
	return false
}

func partTwo(bags map[string]map[string]int) int {
	return countBagsWithin("shiny gold", bags)
}

func countBagsWithin(outerBag string, bags map[string]map[string]int) int {
	total := 0
	for containedBag, number := range bags[outerBag] {
		total += number * (1 + countBagsWithin(containedBag, bags))
	}
	return total
}

func readBags(fileName string) (map[string]map[string]int, error) {
	bags := map[string]map[string]int{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		containedBags := map[string]int{}
		replacer := strings.NewReplacer(".", "", " bags", "", " bag", "", " contain ", ", ")
		text := replacer.Replace(scanner.Text())
		splitText := strings.Split(text, ", ")
		for i := 1; i < len(splitText); i++ {
			if splitText[i] == "no other" {
				break
			}
			splitBag := strings.SplitN(splitText[i], " ", 2)
			numberOfBags, err := strconv.Atoi(splitBag[0])
			if err != nil {
				return nil, err
			}

			containedBags[splitBag[1]] = numberOfBags
		}
		bags[splitText[0]] = containedBags
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bags, nil
}
