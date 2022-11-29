package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	keyOne, keyTwo, err := readKeys("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(keyOne, keyTwo))
}

func partOne(keyOne, keyTwo int) int {
	value, loops := 1, 0
	for {
		loops++
		value *= 7
		value = value % 20201227
		if value == keyOne {
			break
		}
	}
	value = 1
	for i := 0; i < loops; i++ {
		value *= keyTwo
		value = value % 20201227
	}
	return value
}

func readKeys(fileName string) (int, int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, 0, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	keyOne, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, 0, err
	}

	scanner.Scan()
	keyTwo, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, 0, err
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	return keyOne, keyTwo, nil
}
