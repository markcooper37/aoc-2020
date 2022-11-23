package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Rule struct {
	Number     int
	SubRules   [][]int
	RuleLetter byte
}

func main() {
	rulesPartOne, messagesPartOne, err := readMessages("input_1.txt")
	if err != nil {
		log.Fatal(err)
	}

	rulesPartTwo, messagesPartTwo, err := readMessages("input_2.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(countValidMessages(rulesPartOne, messagesPartOne))
	fmt.Println(countValidMessages(rulesPartTwo, messagesPartTwo))
}

func countValidMessages(rules []Rule, messages []string) int {
	initialRule := rules[0]
	validCount := 0
	for _, message := range messages {
		remainders := findValidMessageEnds([]string{message}, initialRule, rules)
		for _, remainder := range remainders {
			if len(remainder) == 0 {
				validCount++
				break
			}
		}
	}
	return validCount
}

func findValidMessageEnds(messages []string, rule Rule, rules []Rule) []string {
	if rule.RuleLetter == 'a' || rule.RuleLetter == 'b' {
		validMessageRemainders := []string{}
		for _, message := range messages {
			if len(message) > 0 && message[0] == rule.RuleLetter {
				validMessageRemainders = append(validMessageRemainders, message[1:])
			}
		}
		return validMessageRemainders
	} else {
		allRemainders := []string{}
	CheckSubRule:
		for _, subRule := range rule.SubRules {
			remainders := messages
			for _, number := range subRule {
				newRule := findRule(rules, number)
				remainders = findValidMessageEnds(remainders, newRule, rules)
				if len(remainders) == 0 {
					continue CheckSubRule
				}
			}
			allRemainders = append(allRemainders, remainders...)
		}
		return allRemainders
	}
}

func findRule(rules []Rule, number int) Rule {
	for _, rule := range rules {
		if rule.Number == number {
			return rule
		}
	}
	return Rule{}
}

func readMessages(fileName string) ([]Rule, []string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	rules := []Rule{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		splitRule := strings.Split(scanner.Text(), ": ")
		ruleName, err := strconv.Atoi(splitRule[0])
		if err != nil {
			return nil, nil, err
		}

		if splitRule[1] == "\"a\"" || splitRule[1] == "\"b\"" {
			rules = append(rules, Rule{Number: ruleName, RuleLetter: strings.Trim(splitRule[1], "\"")[0]})
		} else {

			subRulesStrings := strings.Split(splitRule[1], " | ")
			subRules := [][]int{}
			for _, subRulesString := range subRulesStrings {
				ruleNumbersStrings := strings.Split(subRulesString, " ")
				ruleNumbers := []int{}
				for _, ruleNumbersString := range ruleNumbersStrings {
					number, err := strconv.Atoi(ruleNumbersString)
					if err != nil {
						return nil, nil, err
					}
					ruleNumbers = append(ruleNumbers, number)
				}
				subRules = append(subRules, ruleNumbers)
			}

			rules = append(rules, Rule{Number: ruleName, SubRules: subRules})
		}
	}
	sort.Slice(rules, func(i, j int) bool {
		return rules[i].Number < rules[j].Number
	})
	messages := []string{}
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rules, messages, nil
}
