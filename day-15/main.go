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
	numbers, err := readNumbers("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(numbers))
	fmt.Println(partTwo(numbers))
}

func partOne(numbers []int) int {
	return findNthNumberSpoken(numbers, 2020)
}

func partTwo(numbers []int) int {
	return findNthNumberSpoken(numbers, 30000000)
}

func findNthNumberSpoken(numbers []int, turns int) int {
	if len(numbers) == 0 {
		return 0
	}
	numberLastSpoken := map[int]int{}
	lastNumberSpoken := numbers[0]
	for i := 2; i <= turns; i++ {
		if i <= len(numbers) {
			numberLastSpoken[lastNumberSpoken] = i-1
			lastNumberSpoken = numbers[i-1]
		} else {
			if _, ok := numberLastSpoken[lastNumberSpoken]; !ok {
				numberLastSpoken[lastNumberSpoken] = i-1
				lastNumberSpoken = 0
			} else {
				nextNumber := i-1-numberLastSpoken[lastNumberSpoken]
				numberLastSpoken[lastNumberSpoken] = i-1
				lastNumberSpoken = nextNumber
			}
		}
	}
	return lastNumberSpoken
}

func readNumbers(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []int{}, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numberStrings := strings.Split(scanner.Text(), ",")
	numbers := []int{}
	for _, numberString := range numberStrings {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		return []int{}, err
	}

	return numbers, nil
}
