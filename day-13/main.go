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
	earliestTime, buses, err := readBusData("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(earliestTime, buses))
	fmt.Println(partTwo(buses))
}

func partOne(earliestTime int, buses []string) int {
	minWaitTime, minWaitBus := -1, -1
	for _, bus := range buses {
		if bus != "x" {
			busID, err := strconv.Atoi(bus)
			if err != nil {
				continue
			}
			waitTime := (busID - (earliestTime % busID)) % busID
			if minWaitTime == -1 || waitTime < minWaitTime {
				minWaitTime = waitTime
				minWaitBus = busID
			}
		}
	}
	return minWaitTime * minWaitBus
}

func partTwo(buses []string) int {
	earliestTime, lcm := 0, 1
	for index, bus := range buses {
		if bus != "x" {
			busID, err := strconv.Atoi(bus)
			if err != nil {
				continue
			}
			for {
				if (((earliestTime + index) % busID) + busID) % busID == 0 {
					lcm *= busID
					break
				} else {
					earliestTime += lcm
				}
			}
		}
	}
	return earliestTime
}

func readBusData(fileName string) (int, []string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	earliestTime, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, nil, err
	}

	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")
	if err := scanner.Err(); err != nil {
		return 0, nil, err
	}

	return earliestTime, buses, nil
}
