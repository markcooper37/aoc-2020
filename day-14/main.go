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
	instructions, err := readInstructions("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(instructions))
	fmt.Println(partTwo(instructions))
}

func partOne(instructions [][2]string) uint64 {
	mask := ""
	addressValues := map[uint64]uint64{}
	for _, instruction := range instructions {
		if instruction[0] == "mask" {
			mask = instruction[1]
		} else {
			valueInt, err := strconv.Atoi(instruction[1])
			if err != nil {
				continue
			}

			value := uint64(valueInt)
			addressInt, err := strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(instruction[0], "]"), "mem["))
			if err != nil {
				continue
			}

			address := uint64(addressInt)
			for index, character := range mask {
				if character == '0' {
					value &^= (1 << (35 - index))
				} else if character == '1' {
					value |= (1 << (35 - index))
				}
			}
			addressValues[address] = value
		}
	}
	total := uint64(0)
	for _, value := range addressValues {
		total += value
	}
	return total
}

func partTwo(instructions [][2]string) uint64 {
	mask := ""
	addressValues := map[uint64]uint64{}
	for _, instruction := range instructions {
		if instruction[0] == "mask" {
			mask = instruction[1]
		} else {
			valueInt, err := strconv.Atoi(instruction[1])
			if err != nil {
				continue
			}

			value := uint64(valueInt)
			addressInt, err := strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(instruction[0], "]"), "mem["))
			if err != nil {
				continue
			}

			address := uint64(addressInt)
			for index, character := range mask {
				if character == '1' {
					address |= (1 << (35 - index))
				}
			}
			addresses := []uint64{address}
			for index, character := range mask {
				if character == 'X' {
					newAddresses := []uint64{}
					for _, addressValue := range addresses {
						newAddresses = append(newAddresses, addressValue &^ (1 << (35 - index)))
						newAddresses = append(newAddresses, addressValue | (1 << (35 - index)))
					}
					addresses = newAddresses
				}
			}
			for _, addressValue := range addresses {
				addressValues[addressValue] = value
			}
		}
	}
	total := uint64(0)
	for _, value := range addressValues {
		total += value
	}
	return total
}

func readInstructions(fileName string) ([][2]string, error) {
	instructions := [][2]string{}
	file, err := os.Open(fileName)
	if err != nil {
		return [][2]string{}, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " = ")
		instructions = append(instructions, [2]string{instruction[0], instruction[1]})
	}
	if err := scanner.Err(); err != nil {
		return [][2]string{}, err
	}

	return instructions, nil
}
