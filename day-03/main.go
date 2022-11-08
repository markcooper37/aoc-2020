package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	grid, err := readGrid("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(grid))
	fmt.Println(partTwo(grid))
}

func partOne(grid []string) int {
	return checkSlope(grid, 3, 1)
}

func partTwo(grid []string) int {
	treeCounts := []int{checkSlope(grid, 1, 1), checkSlope(grid, 3, 1), checkSlope(grid, 5, 1), checkSlope(grid, 7, 1), checkSlope(grid, 1, 2)}
	product := 1
	for _, count := range treeCounts {
		product *= count
	}
	return product
}

func checkSlope(grid []string, right, down int) int {
	column, treeCount := 0, 0
	for row := down; row < len(grid); row += down {
		column = (column + right) % len(grid[row])
		if grid[row][column] == '#' {
			treeCount++
		}
	}
	return treeCount
}

func readGrid(fileName string) ([]string, error) {
	rows := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		rows = append(rows, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}
