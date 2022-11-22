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
	expressions, err := readExpressions("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(expressions))
	fmt.Println(partTwo(expressions))
}

func partOne(expressions []string) int {
	return sumAllExpressions(expressions, false)
}

func partTwo(expressions []string) int {
	return sumAllExpressions(expressions, true)
}

func sumAllExpressions(expressions []string, plusPrecedence bool) int {
	sum := 0
	for _, expression := range expressions {
		sum += evaluateExpression(expression, plusPrecedence)
	}
	return sum
}

func evaluateExpression(expression string, plusPrecedence bool) int {
	for {
		firstBracketedExpression, bracketsExist := findFirstBracketedExpression(expression)
		if !bracketsExist {
			break
		}
		expression = strings.Replace(expression, firstBracketedExpression, strconv.Itoa(evaluateSimpleExpression(firstBracketedExpression, plusPrecedence)), 1)
	}
	return evaluateSimpleExpression(expression, plusPrecedence)
}

func findFirstBracketedExpression(expression string) (string, bool) {
	lastOpenBracketPosition := 0
	for index, character := range expression {
		if character == '(' {
			lastOpenBracketPosition = index
		} else if character == ')' {
			return expression[lastOpenBracketPosition : index+1], true
		}
	}
	return "", false
}

func evaluateSimpleExpression(expression string, plusPrecedence bool) int {
	if plusPrecedence {
		return evaluateSimpleExpressionPlusPrecedence(expression)
	} else {
		return evaluateSimpleExpressionNoPrecedence(expression)
	}
}

func evaluateSimpleExpressionPlusPrecedence(expression string) int {
	components := strings.Split(strings.Trim(expression, "()"), " ")
	newComponents := []int{}
	for i := 0; i < len(components); i++ {
		if components[i] == "*" {
			value, err := strconv.Atoi(components[i-1])
			if err != nil {
				return -1
			}

			newComponents = append(newComponents, value)
		} else if components[i] == "+" {
			firstValue, err := strconv.Atoi(components[i-1])
			if err != nil {
				return -1
			}

			secondValue, err := strconv.Atoi(components[i+1])
			if err != nil {
				return -1
			}

			components[i+1] = strconv.Itoa(firstValue + secondValue)
		} else if i == len(components)-1 {
			value, err := strconv.Atoi(components[i])
			if err != nil {
				return -1
			}

			newComponents = append(newComponents, value)
		}
	}
	total := 1
	for _, value := range newComponents {
		total *= value
	}
	return total
}

func evaluateSimpleExpressionNoPrecedence(expression string) int {
	operator := ""
	total := 0
	components := strings.Split(strings.Trim(expression, "()"), " ")
	for index, component := range components {
		if component == "*" || component == "+" {
			operator = component
		} else {
			value, err := strconv.Atoi(component)
			if err != nil {
				return -1
			}

			if index == 0 {
				total = value
			} else if operator == "+" {
				total += value
			} else if operator == "*" {
				total *= value
			}
		}
	}
	return total
}

func readExpressions(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	expressions := []string{}
	for scanner.Scan() {
		expressions = append(expressions, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return expressions, nil
}
