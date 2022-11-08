package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Passport map[string]string

func main() {
	passports, err := readPassports("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(passports))
	fmt.Println(partTwo(passports))
}

func partOne(passports []Passport) int {
	validCount := 0
	for _, passport := range passports {
		if passport.containsRequiredFields() {
			validCount++
		}
	}
	return validCount
}

func (passport Passport) containsRequiredFields() bool {
	_, byrOK := passport["byr"]
	_, iyrOK := passport["iyr"]
	_, eyrOK := passport["eyr"]
	_, hgtOK := passport["hgt"]
	_, hclOK := passport["hcl"]
	_, eclOK := passport["ecl"]
	_, pidOK := passport["pid"]
	return byrOK && iyrOK && eyrOK && hgtOK && hclOK && eclOK && pidOK
}

func partTwo(passports []Passport) int {
	validCount := 0
	for _, passport := range passports {
		if passport.isValid() {
			validCount++
		}
	}
	return validCount
}

func (passport Passport) isValid() bool {
	if !passport.containsRequiredFields() {
		return false
	}
	return isBYRValid(passport["byr"]) && isIYRValid(passport["iyr"]) &&
		isEYRValid(passport["eyr"]) && isHGTValid(passport["hgt"]) &&
		isHCLValid(passport["hcl"]) && isECLValid(passport["ecl"]) &&
		isPIDValid(passport["pid"])
}

func isBYRValid(byr string) bool {
	if len(byr) != 4 {
		return false
	}
	year, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}
	return year >= 1920 && year <= 2002
}

func isIYRValid(iyr string) bool {
	if len(iyr) != 4 {
		return false
	}
	year, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}
	return year >= 2010 && year <= 2020
}

func isEYRValid(eyr string) bool {
	if len(eyr) != 4 {
		return false
	}
	year, err := strconv.Atoi(eyr)
	if err != nil {
		return false
	}
	return year >= 2020 && year <= 2030
}

func isHGTValid(hgt string) bool {
	if strings.HasSuffix(hgt, "cm") {
		hgt = strings.TrimSuffix(hgt, "cm")
		numericValue, err := strconv.Atoi(hgt)
		if err != nil || numericValue < 150 || numericValue > 193 {
			return false
		}
		return true
	} else if strings.HasSuffix(hgt, "in") {
		hgt = strings.TrimSuffix(hgt, "in")
		numericValue, err := strconv.Atoi(hgt)
		if err != nil || numericValue < 59 || numericValue > 76 {
			return false
		}
		return true
	}
	return false
}

func isHCLValid(hcl string) bool {
	if len(hcl) != 7 || hcl[0] != '#' {
		return false
	}
	for i := 1; i < 7; i++ {
		if hcl[i] < 48 || (hcl[i] > 57 && hcl[i] < 97) || hcl[i] > 102 {
			return false
		}
	}
	return true
}

func isECLValid(ecl string) bool {
	return ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" ||
		ecl == "grn" || ecl == "hzl" || ecl == "oth"
}

func isPIDValid(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	_, err := strconv.Atoi(pid)
	return err == nil
}

func readPassports(fileName string) ([]Passport, error) {
	passports := []Passport{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentPassport := Passport{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			passports = append(passports, currentPassport)
			currentPassport = Passport{}
		} else {
			fields := strings.Split(scanner.Text(), " ")
			for _, field := range fields {
				keyValuePair := strings.Split(field, ":")
				if len(keyValuePair) < 2 {
					return []Passport{}, errors.New("field does not contain key value pair")
				}

				currentPassport[keyValuePair[0]] = keyValuePair[1]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(currentPassport) != 0 {
		passports = append(passports, currentPassport)
	}
	return passports, nil
}
