package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	seats, err := readSeats("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(seats))
	fmt.Println(partTwo(seats))
}

func partOne(seats [][]rune) int {
	seatChanged := true
	for seatChanged {
		seats, seatChanged = determineNewOccupationLayout(seats, true, 4)
	}
	return countAllOccupiedSeats(seats)
}

func determineNewOccupationLayout(seats [][]rune, surrounding bool, leaveSeatThreshold int) ([][]rune, bool) {
	seatChanged := false
	newSeats := [][]rune{}
	for rowIndex, row := range seats {
		newRow := []rune{}
		for columnIndex, seat := range row {
			if seat != '.' {
				occupiedSeatsCount := countNearbyOccupiedSeats(rowIndex, columnIndex, seats, surrounding)
				if  seat == '#' && occupiedSeatsCount >= leaveSeatThreshold {
					newRow = append(newRow, 'L')
					seatChanged = true
				} else if  seat == 'L' && occupiedSeatsCount ==0 {
					newRow = append(newRow, '#')
					seatChanged = true
				} else {
					newRow = append(newRow, seat)
				}
			} else {
				newRow = append(newRow, '.')
			}
		}
		newSeats = append(newSeats, newRow)
	}
	return newSeats, seatChanged
}

func countNearbyOccupiedSeats(rowIndex, columnIndex int, seats [][]rune, surrounding bool) int {
	if len(seats) == 0 {
		return 0
	}
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			currentRow, currentColumn := rowIndex, columnIndex
			for {
				currentRow += i
				currentColumn += j
				if currentRow < 0 || currentRow >= len(seats) || currentColumn < 0 || currentColumn >= len(seats[0]) {
					break
				}
				if seats[currentRow][currentColumn] == 'L' {
					break
				} else if seats[currentRow][currentColumn] == '#' {
					count++
					break
				}
				if surrounding {
					break
				}
			}
		}
	}
	return count
}

func countAllOccupiedSeats(seats [][]rune) int {
	occupiedSeatsCount := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == '#' {
				occupiedSeatsCount++
			}
		}
	}
	return occupiedSeatsCount
}

func partTwo(seats [][]rune) int {
	seatChanged := true
	for seatChanged {
		seats, seatChanged = determineNewOccupationLayout(seats, false, 5)
	}
	return countAllOccupiedSeats(seats)
}

func readSeats(fileName string) ([][]rune, error) {
	seats := [][]rune{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seats = append(seats, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}
