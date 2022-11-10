package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	forms, err := readForms("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(forms))
	fmt.Println(partTwo(forms))
}

func partOne(forms [][]string) int {
	overallTotal := 0
	for _, groupForms := range forms {
		yesAnswers := countYesAnswers(groupForms)
		overallTotal += len(yesAnswers)
	}
	return overallTotal
}

func partTwo(forms [][]string) int {
	overallTotal := 0
	for _, groupForms := range forms {
		yesAnswers := countYesAnswers(groupForms)
		for _, value := range yesAnswers {
			if value == len(groupForms) {
				overallTotal++
			}
		}
	}
	return overallTotal
}

func countYesAnswers(forms []string) map[rune]int {
	yesAnswers := map[rune]int{}
	for _, form := range forms {
		for _, yesAnswer := range form {
			yesAnswers[yesAnswer] += 1
		}
	}
	return yesAnswers
}

func readForms(fileName string) ([][]string, error) {
	forms := [][]string{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentForms := []string{}
	for scanner.Scan() {
		if scanner.Text() != "" {
			currentForms = append(currentForms, scanner.Text())
		} else {
			forms = append(forms, currentForms)
			currentForms = []string{}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(currentForms) != 0 {
		forms = append(forms, currentForms)
	}
	return forms, nil
}
