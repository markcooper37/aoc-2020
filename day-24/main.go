package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	instructions, err := readInstructions("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	blackTiles, blackTileCount := partOne(instructions)
	fmt.Println(blackTileCount)
	fmt.Println(partTwo(blackTiles))
}

func partOne(instructions []string) (map[[2]int]bool, int) {
	blackTiles := map[[2]int]bool{}
	for _, instruction := range instructions {
		x, y := 0, 0
		for i := 0; i < len(instruction); i++ {
			character := instruction[i]
			if character == 'e' {
				x += 2
			} else if character == 'w' {
				x -= 2
			} else if character == 'n' {
				y += 1
				if instruction[i+1] == 'e' {
					x += 1
				} else if instruction[i+1] == 'w' {
					x -= 1
				}
				i++
			} else if character == 's' {
				y -= 1
				if instruction[i+1] == 'e' {
					x += 1
				} else if instruction[i+1] == 'w' {
					x -= 1
				}
				i++
			}
		}
		if !blackTiles[[2]int{x, y}] {
			blackTiles[[2]int{x, y}] = true
		} else {
			delete(blackTiles, [2]int{x, y})
		}
	}
	return blackTiles, len(blackTiles)
}

func partTwo(blackTiles map[[2]int]bool) int {
	for i := 0; i < 100; i++ {
		blackTiles = findNewBlackTiles(blackTiles)
	}
	return len(blackTiles)
}

func getAdjacentTiles(tilePosition [2]int) [][2]int {
	x, y := tilePosition[0], tilePosition[1]
	return [][2]int{{x + 2, y}, {x - 2, y}, {x + 1, y + 1}, {x + 1, y - 1}, {x - 1, y - 1}, {x - 1, y + 1}}
}

func findNewBlackTiles(blackTiles map[[2]int]bool) map[[2]int]bool {
	newBlackTiles := map[[2]int]bool{}
	checked := map[[2]int]bool{}
	for blackTile := range blackTiles {
		adjacentBlackCount := 0
		adjacentTiles := getAdjacentTiles(blackTile)
		for _, adjacentTile := range adjacentTiles {
			if blackTiles[adjacentTile] {
				adjacentBlackCount++
			}
			if !blackTiles[adjacentTile] && !checked[adjacentTile] {
				countBlackAdjacentToAdjacent := 0
				for _, adjacentToAdjacent := range getAdjacentTiles(adjacentTile) {
					if blackTiles[adjacentToAdjacent] == true {
						countBlackAdjacentToAdjacent++
					}
				}
				if countBlackAdjacentToAdjacent == 2 {
					newBlackTiles[adjacentTile] = true
				}
				checked[adjacentTile] = true
			}
		}
		if adjacentBlackCount == 1 || adjacentBlackCount == 2 {
			newBlackTiles[blackTile] = true
		}
	}
	return newBlackTiles
}

func readInstructions(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	instructions := []string{}
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return instructions, nil
}
