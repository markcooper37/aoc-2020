package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PasswordDatum struct {
	Letter     byte
	LowerValue int
	UpperValue int
	Password   string
}

func main() {
	passwordData, err := readPasswordData("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(passwordData))
	fmt.Println(partTwo(passwordData))
}

func partOne(passwordData []PasswordDatum) int {
	correctPasswordsCount := 0
	for _, passwordDatum := range passwordData {
		letterCount := 0
		for _, letter := range passwordDatum.Password {
			if letter == rune(passwordDatum.Letter) {
				letterCount++
			}
		}
		if letterCount >= passwordDatum.LowerValue && letterCount <= passwordDatum.UpperValue {
			correctPasswordsCount++
		}
	}
	return correctPasswordsCount
}

func partTwo(passwordData []PasswordDatum) int {
	correctPasswordsCount := 0
	for _, passwordDatum := range passwordData {
		if (passwordDatum.Password[passwordDatum.LowerValue-1] == passwordDatum.Letter && passwordDatum.Password[passwordDatum.UpperValue-1] != passwordDatum.Letter) ||
			(passwordDatum.Password[passwordDatum.LowerValue-1] != passwordDatum.Letter && passwordDatum.Password[passwordDatum.UpperValue-1] == passwordDatum.Letter) {
			correctPasswordsCount++
		}
	}
	return correctPasswordsCount
}

func readPasswordData(fileName string) ([]PasswordDatum, error) {
	passwordData := []PasswordDatum{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		passwordDatum := PasswordDatum{Letter: text[1][0], Password: text[2]}
		bounds := strings.Split(text[0], "-")
		lowerValue, err := strconv.Atoi(bounds[0])
		if err != nil {
			return nil, err
		}

		passwordDatum.LowerValue = lowerValue
		upperValue, err := strconv.Atoi(bounds[1])
		if err != nil {
			return nil, err
		}

		passwordDatum.UpperValue = upperValue
		passwordData = append(passwordData, passwordDatum)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return passwordData, nil
}
