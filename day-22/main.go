package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	decks, err := readDecks("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(decks))
	fmt.Println(partTwo(decks))
}

func partOne(decks [][]int) int {
	copyDecks := [][]int{}
	for _, deck := range decks {
		copyDeck := []int{}
		copy(copyDeck, deck)
		copyDecks = append(copyDecks, deck)
	}
	for len(copyDecks[0]) > 0 && len(copyDecks[1]) > 0 {
		firstValue := copyDecks[0][0]
		secondValue := copyDecks[1][0]
		if firstValue > secondValue {
			copyDecks = determineNewDecks(copyDecks, 0)
		} else {
			copyDecks = determineNewDecks(copyDecks, 1)
		}

	}
	winningDeck := copyDecks[0]
	if len(copyDecks[1]) > 0 {
		winningDeck = copyDecks[1]
	}
	return calculateTotal(winningDeck)
}

func calculateTotal(deck []int) int {
	total := 0
	for i := 1; i <= len(deck); i++ {
		total += i * deck[len(deck)-i]
	}
	return total
}

func determineNewDecks(decks [][]int, winner int) [][]int {
	firstValue := decks[0][0]
	secondValue := decks[1][0]
	decks[0] = decks[0][1:]
	decks[1] = decks[1][1:]
	if winner == 0 {
		decks[0] = append(decks[0], append([]int{firstValue}, secondValue)...)
	} else {
		decks[1] = append(decks[1], append([]int{secondValue}, firstValue)...)
	}
	return decks
}

func partTwo(decks [][]int) int {
	winningDeck, _ := findWinningRecursiveDeck(decks)
	return calculateTotal(winningDeck)
}

func findWinningRecursiveDeck(decks [][]int) ([]int, int) {
	previousDecks := [][][]int{}
	for len(decks[0]) > 0 && len(decks[1]) > 0 {
		for _, previous := range previousDecks {
			if areDeckPairsEqual(previous, decks) {
				return decks[0], 0
			}
		}
		previousOne, previousTwo := []int{}, []int{}
		previousOne = append(previousOne, decks[0]...)
		previousTwo = append(previousTwo, decks[1]...)
		previousDecks = append(previousDecks, append([][]int{previousOne}, previousTwo))
		winner := 0
		if decks[0][0] < len(decks[0]) && decks[1][0] < len(decks[1]) {
			subDeckOne, subDeckTwo := []int{}, []int{}
			subDeckOne = append(subDeckOne, decks[0][1:decks[0][0]+1]...)
			subDeckTwo = append(subDeckTwo, decks[1][1:decks[1][0]+1]...)
			_, winner = findWinningRecursiveDeck([][]int{subDeckOne, subDeckTwo})
		} else {
			if decks[0][0] < decks[1][0] {
				winner = 1
			}
		}
		decks = determineNewDecks(decks, winner)
	}
	if len(decks[0]) > 0 {
		return decks[0], 0
	} else {
		return decks[1], 1
	}
}

func areDeckPairsEqual(deckPairOne, deckPairTwo [][]int) bool {
	if len(deckPairOne[0]) != len(deckPairTwo[0]) || len(deckPairOne[1]) != len(deckPairTwo[1]) {
		return false
	}
	for index := range deckPairOne[0] {
		if deckPairOne[0][index] != deckPairTwo[0][index] {
			return false
		}
	}
	for index := range deckPairOne[1] {
		if deckPairOne[1][index] != deckPairTwo[1][index] {
			return false
		}
	}
	return true
}

func readDecks(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	decks := [][]int{}
	newDeck := []int{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			decks = append(decks, newDeck)
			newDeck = []int{}
		} else if scanner.Text()[0] == 'P' {
			continue
		} else {
			value, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}

			newDeck = append(newDeck, value)
		}
	}
	decks = append(decks, newDeck)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return decks, nil
}
