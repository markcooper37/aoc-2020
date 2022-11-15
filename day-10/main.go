package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	numbers, err := readNumbers("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	numbers = append([]int{0}, append(numbers, numbers[len(numbers)-1]+3)...)
	fmt.Println(partOne(numbers))
	fmt.Println(partTwo(numbers))
}

func partOne(numbers []int) int {
	differenceOneCount, differenceThreeCount := 0, 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i+1]-numbers[i] == 1 {
			differenceOneCount++
		} else if numbers[i+1]-numbers[i] == 3 {
			differenceThreeCount++
		}
	}
	return differenceOneCount * differenceThreeCount
}

func partTwo(numbers []int) int {
	counts := make([]int, len(numbers))
	for index, number := range numbers {
		// Initialise the first count to be 1
		if index == 0 {
			counts[index] = 1
		}
		for j := 1; j <= 3; j++ {
			if index-j >= 0 && number-numbers[index-j] <= 3 {
				counts[index] += counts[index-j]
			}
		}
	}
	return counts[len(counts)-1]
}

func readNumbers(fileName string) ([]int, error) {
	numbers := []int{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	sort.Ints(numbers)
	return numbers, nil
}
