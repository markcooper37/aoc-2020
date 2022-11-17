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
	Action string
	Value  int
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
	shipX, shipY := 0, 0
	directionsOrdered := []string{"E", "S", "W", "N"}
	currentDirectionIndex := 0
	for _, instruction := range instructions {
		if instruction.Action == "L" {
			currentDirectionIndex = (currentDirectionIndex + 3*instruction.Value/90) % 4
		} else if instruction.Action == "R" {
			currentDirectionIndex = (currentDirectionIndex + instruction.Value/90) % 4
		} else if instruction.Action == "F" {
			shipX, shipY = findNewCoordinates(shipX, shipY, instruction.Value, directionsOrdered[currentDirectionIndex])
		} else {
			shipX, shipY = findNewCoordinates(shipX, shipY, instruction.Value, instruction.Action)
		}
	}
	return absoluteValue(shipX) + absoluteValue(shipY)
}

func findNewCoordinates(currentX, currentY, value int, direction string) (int, int) {
	if direction == "E" {
		currentX += value
	} else if direction == "W" {
		currentX -= value
	} else if direction == "N" {
		currentY += value
	} else if direction == "S" {
		currentY -= value
	}
	return currentX, currentY
}

func absoluteValue(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func partTwo(instructions []Instruction) int {
	shipX, shipY, waypointX, waypointY := 0, 0, 10, 1
	for _, instruction := range instructions {
		if instruction.Action == "L" {
			for i := 0; i < instruction.Value/90; i++ {
				waypointX, waypointY = -waypointY, waypointX
			}
		} else if instruction.Action == "R" {
			for i := 0; i < instruction.Value/90; i++ {
				waypointX, waypointY = waypointY, -waypointX
			}
		} else if instruction.Action == "F" {
			shipX += instruction.Value * waypointX
			shipY += instruction.Value * waypointY
		} else {
			waypointX, waypointY = findNewCoordinates(waypointX, waypointY, instruction.Value, instruction.Action)
		}
	}
	return absoluteValue(shipX) + absoluteValue(shipY)
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
		splitInstructions := strings.SplitAfterN(scanner.Text(), "", 2)
		value, err := strconv.Atoi(splitInstructions[1])
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, Instruction{Action: splitInstructions[0], Value: value})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return instructions, nil
}
