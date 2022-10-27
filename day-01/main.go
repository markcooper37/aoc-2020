package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	entries, err := readEntries("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(entries))
	fmt.Println(partTwo(entries))
}

func partOne(entries []int) int {
	for i, entry := range entries {
		for j := i + 1; j < len(entries); j++ {
			if entry+entries[j] == 2020 {
				return entry * entries[j]
			}
		}
	}
	return -1
}

func partTwo(entries []int) int {
	for i, entry := range entries {
		for j := i + 1; j < len(entries); j++ {
			for k := j + 1; k < len(entries); k++ {
				if entry+entries[j]+entries[k] == 2020 {
					return entry * entries[j] * entries[k]
				}
			}
		}
	}
	return -1
}

func readEntries(fileName string) ([]int, error) {
	entries := []int{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		entries = append(entries, i)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
