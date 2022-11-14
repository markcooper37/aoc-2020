package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Name  string
	Value int
}

func main() {
	instructions, err := readInstructions("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(instructions))
	fmt.Println(partTwo(instructions))
}

func partOne(instructions []Instruction) int {
	accumulator, _ := calculateAccumulator(instructions)
	return accumulator
}

func calculateAccumulator(instructions []Instruction) (int, bool) {
	currentIndex, accumulator := 0, 0
	reachedEnd := false
	visitedIndices := map[int]bool{}
	for !visitedIndices[currentIndex] {
		if currentIndex >= len(instructions) {
			reachedEnd = true
			break
		}
		visitedIndices[currentIndex] = true
		if instructions[currentIndex].Name == "acc" {
			accumulator += instructions[currentIndex].Value
			currentIndex++
		} else if instructions[currentIndex].Name == "jmp" {
			currentIndex += instructions[currentIndex].Value
		} else {
			currentIndex++
		}
	}
	return accumulator, reachedEnd
}

func partTwo(instructions []Instruction) int {
	nameSwapMap := map[string]string{"jmp": "nop", "nop": "jmp"}
	accumulator := 0
	for index, instruction := range instructions {
		if instruction.Name == "acc" {
			continue
		} else {
			instructions[index].Name = nameSwapMap[instruction.Name]
		}
		accumulatorValue, reachedEnd := calculateAccumulator(instructions)
		if reachedEnd {
			accumulator = accumulatorValue
			break
		} else {
			instructions[index].Name = nameSwapMap[instructions[index].Name]
		}
	}
	return accumulator
}

func readInstructions(fileName string) ([]Instruction, error) {
	instructions := []Instruction{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitText := strings.Split(scanner.Text(), " ")
		value, err := strconv.Atoi(splitText[1])
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, Instruction{Name: splitText[0], Value: value})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return instructions, nil
}
