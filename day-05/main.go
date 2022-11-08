package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
)

type BoardingPass [2]string

func main() {
	boardingPasses, err := readBoardingPasses("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(boardingPasses))
	fmt.Println(partTwo(boardingPasses))
}

func partOne(boardingPasses []BoardingPass) int {
	highestSeatID := 0
	for _, boardingPass := range boardingPasses {
		seatID := decipherSeatID(boardingPass)
		if highestSeatID < seatID {
			highestSeatID = seatID
		}
	}
	return highestSeatID
}

func decipherSeatID(boardingPass BoardingPass) int {
	row := decipherPosition(boardingPass[0], 'B')
	column := decipherPosition(boardingPass[1], 'R')
	return row*8+column
}

func decipherPosition(code string, lastHalfCharacter byte) int {
	lowerBound := 0
	for i := 0; i < len(code); i++ {
		if code[i] == lastHalfCharacter {
			lowerBound += 1 << (len(code) - i - 1)
		}
	}
	return lowerBound
}

func partTwo(boardingPasses []BoardingPass) int {
	seatIDs := []int{}
	for _, boardingPass := range boardingPasses {
		seatIDs = append(seatIDs, decipherSeatID(boardingPass))
	}
	sort.Ints(seatIDs)
	for index, seatID := range seatIDs {
		if seatID - seatIDs[0] != index {
			return seatID - 1
		}
	}
	return -1
}

func readBoardingPasses(fileName string) ([]BoardingPass, error) {
	boardingPasses := []BoardingPass{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) < 10 {
			return []BoardingPass{}, errors.New("not enough characters in boarding pass")
		}
		
		boardingPasses = append(boardingPasses, BoardingPass{scanner.Text()[0:7], scanner.Text()[7:10]})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return boardingPasses, nil
}
