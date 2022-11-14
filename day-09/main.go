package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers, err := readNumbers("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(25, numbers))
	fmt.Println(partTwo(25, numbers))
}

func partOne(preambleLength int, numbers []int) int {
	for i := preambleLength; i < len(numbers); i++ {
		if !valueHasSum(numbers[i], numbers[i-preambleLength:i]) {
			return numbers[i]
		}
	}
	return -1
}

func valueHasSum(value int, numbers []int) bool {
	for i := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == value {
				return true
			}
		}
	}
	return false
}

func partTwo(preambleLength int, numbers []int) int {
	requiredValue := partOne(preambleLength, numbers)
	sumValues := findSumValues(requiredValue, numbers)
	if len(sumValues) == 0 {
		return -1
	}
	minValue, maxValue := sumValues[0], sumValues[0]
	for _, value := range sumValues {
		if value < minValue {
			minValue = value
		} else if value > maxValue {
			maxValue = value
		}
	}
	return minValue + maxValue
}

func findSumValues(value int, numbers []int) []int {
	for i, number := range numbers {
		sum := number
		for j := i + 1; j < len(numbers); j++ {
			sum += numbers[j]
			if sum == value {
				return numbers[i : j+1]
			} else if sum > value {
				continue
			}
		}
	}
	return []int{}
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

	return numbers, nil
}
