package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type TicketInfo struct {
	Fields        []Field
	YourTicket    []int
	NearbyTickets [][]int
}

type Field struct {
	Name   string
	Ranges [][2]int
}

func main() {
	ticketInfo, err := readTicketInformation("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(ticketInfo))
	fmt.Println(partTwo(ticketInfo))
}

func partOne(ticketInfo TicketInfo) int {
	invalidTotal := 0
	for _, nearbyTicket := range ticketInfo.NearbyTickets {
		_, invalidValue := findInvalidValuesOnTicket(nearbyTicket, ticketInfo.Fields)
		invalidTotal += invalidValue
	}
	return invalidTotal
}

func partTwo(ticketInfo TicketInfo) int {
	validNearbyTickets := findValidNearbyTickets(ticketInfo.NearbyTickets, ticketInfo.Fields)
	validFieldsCounts := countValidFieldsForColumn(validNearbyTickets, ticketInfo.Fields)
	validFieldsForColumn := findValidFieldsForColumns(validNearbyTickets, validFieldsCounts, ticketInfo.Fields)
	correctFieldForColumns := findCorrectFieldForColumns(validFieldsForColumn, ticketInfo.Fields)
	output := 1
	for key, value := range correctFieldForColumns {
		if strings.HasPrefix(ticketInfo.Fields[value].Name, "departure") {
			output *= ticketInfo.YourTicket[key]
		}
	}
	return output
}

func findInvalidValuesOnTicket(ticket []int, fields []Field) (bool, int) {
	isInvalid := false
	invalidTotal := 0
	for _, value := range ticket {
		validFields := findValidFields(value, fields)
		if len(validFields) == 0 {
			isInvalid = true
			invalidTotal += value
		}
	}
	return isInvalid, invalidTotal
}

func findValidFields(value int, fields []Field) []int {
	validFields := []int{}
	for index, field := range fields {
		for _, valueRange := range field.Ranges {
			if value >= valueRange[0] && value <= valueRange[1] {
				validFields = append(validFields, index)
			}
		}
	}
	return validFields
}

func findValidNearbyTickets(nearbyTickets [][]int, fields []Field) [][]int {
	validNearbyTickets := [][]int{}
	for _, nearbyTicket := range nearbyTickets {
		if invalid, _ := findInvalidValuesOnTicket(nearbyTicket, fields); !invalid {
			validNearbyTickets = append(validNearbyTickets, nearbyTicket)
		}
	}
	return validNearbyTickets
}

func countValidFieldsForColumn(tickets [][]int, fields []Field) []map[int]int {
	validFieldsCounts := []map[int]int{}
	for range fields {
		validFieldsCounts = append(validFieldsCounts, map[int]int{})
	}
	for _, ticket := range tickets {
		for index, value := range ticket {
			for _, validField := range findValidFields(value, fields) {
				validFieldsCounts[index][validField]++
			}
		}
	}
	return validFieldsCounts
}

func findValidFieldsForColumns(validNearbyTickets [][]int, validFieldsCounts []map[int]int, fields []Field) []map[int]bool {
	validFields := []map[int]bool{}
	for _, validFieldsMap := range validFieldsCounts {
		fieldIndices := map[int]bool{}
		for fieldIndex, value := range validFieldsMap {
			if value == len(validNearbyTickets) {
				fieldIndices[fieldIndex] = true
			}
		}
		validFields = append(validFields, fieldIndices)
	}
	return validFields
}

func findCorrectFieldForColumns(validFieldsForColumn []map[int]bool, fields []Field) map[int]int {
	correctFieldForColumn := map[int]int{}
	for len(correctFieldForColumn) < len(fields) {
		for mapIndex, fieldsMap := range validFieldsForColumn {
			if len(fieldsMap) == 1 {
				for fieldIndex := range fieldsMap {
					correctFieldForColumn[mapIndex] = fieldIndex
					for _, secondFieldsMap := range validFieldsForColumn {
						delete(secondFieldsMap, fieldIndex)
					}
				}
				break
			}
		}
	}
	return correctFieldForColumn
}

func readTicketInformation(fileName string) (TicketInfo, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return TicketInfo{}, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	ticketInfo := TicketInfo{Fields: []Field{}, YourTicket: []int{}, NearbyTickets: [][]int{}}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		} else {
			fieldInfo := strings.Split(scanner.Text(), ": ")
			field := Field{Name: fieldInfo[0], Ranges: [][2]int{}}
			rangeStrings := strings.Split(fieldInfo[1], " or ")
			for _, rangeString := range rangeStrings {
				rangeValueStrings := strings.Split(rangeString, "-")
				lowerBound, err := strconv.Atoi(rangeValueStrings[0])
				if err != nil {
					return TicketInfo{}, err
				}

				upperBound, err := strconv.Atoi(rangeValueStrings[1])
				if err != nil {
					return TicketInfo{}, err
				}

				field.Ranges = append(field.Ranges, [2]int{lowerBound, upperBound})
			}
			ticketInfo.Fields = append(ticketInfo.Fields, field)
		}
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		} else if scanner.Text() != "your ticket:" {
			valueStrings := strings.Split(scanner.Text(), ",")
			for _, valueString := range valueStrings {
				value, err := strconv.Atoi(valueString)
				if err != nil {
					return TicketInfo{}, err
				}

				ticketInfo.YourTicket = append(ticketInfo.YourTicket, value)
			}
		}
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		} else if scanner.Text() != "nearby tickets:" {
			valueStrings := strings.Split(scanner.Text(), ",")
			values := []int{}
			for _, valueString := range valueStrings {
				value, err := strconv.Atoi(valueString)
				if err != nil {
					return TicketInfo{}, err
				}

				values = append(values, value)
			}
			ticketInfo.NearbyTickets = append(ticketInfo.NearbyTickets, values)
		}
	}

	if err := scanner.Err(); err != nil {
		return TicketInfo{}, err
	}

	return ticketInfo, nil
}
