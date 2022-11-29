package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value    int
	Next     *Node
	Children []*Node
}

func main() {
	cups, err := readCups("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(cups))
	fmt.Println(partTwo(cups))
}

func partOne(inputCups []int) string {
	oneNode := performMoves(inputCups, 9, 100)
	currentNode := oneNode.Next
	output := ""
	for i := 1; i < len(inputCups); i++ {
		output += strconv.Itoa(currentNode.Value)
		currentNode = currentNode.Next
	}
	return output
}

func partTwo(inputCups []int) int {
	oneNode := performMoves(inputCups, 1000000, 10000000)
	return oneNode.Next.Value * oneNode.Next.Next.Value
}

func performMoves(inputCups []int, totalCups, moves int) *Node {
	currentNode, oneNode := initialiseNodes(inputCups, totalCups)
	for move := 1; move <= moves; move++ {
		performMove(currentNode, totalCups)
		currentNode = currentNode.Next
	}
	return oneNode
}

func performMove(currentNode *Node, totalCups int) {
	for k := 0; k < 4; k++ {
		value := currentNode.Value - k - 1
		if value <= 0 {
			value += totalCups
		}
		if currentNode.Next.Value != value && currentNode.Next.Next.Value != value && currentNode.Next.Next.Next.Value != value {
			moveThreeNodes(currentNode, currentNode.Children[k])
			break
		}
	}
}

func initialiseNodes(inputCups []int, totalCups int) (*Node, *Node) {
	newNode := Node{Value: inputCups[0]}
	nodeMap := map[int]*Node{}
	nodeMap[inputCups[0]] = &newNode
	currentNode := &newNode
	var oneNode *Node
	for i := 1; i < len(inputCups); i++ {
		node := Node{Value: inputCups[i]}
		nodeMap[inputCups[i]] = &node
		if node.Value == 1 {
			oneNode = &node
		}
		currentNode.Next = &node
		currentNode = &node
	}
	for j := len(inputCups); j < totalCups; j++ {
		node := Node{Value: j + 1}
		nodeMap[j+1] = &node
		currentNode.Next = &node
		currentNode = &node
	}
	currentNode.Next = &newNode
	currentNode = &newNode
	for i := 0; i < totalCups; i++ {
		for j := 1; j <= 4; j++ {
			value := currentNode.Value - j
			if value <= 0 {
				value += totalCups
			}
			currentNode.Children = append(currentNode.Children, nodeMap[value])
		}
		currentNode = currentNode.Next
	}
	return currentNode, oneNode
}

func moveThreeNodes(oldParent *Node, newParent *Node) {
	newParentOldNext := newParent.Next
	newParent.Next = oldParent.Next
	oldParent.Next = oldParent.Next.Next.Next.Next
	newParent.Next.Next.Next.Next = newParentOldNext
}

func readCups(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	cups := []int{}
	cupStringValues := strings.Split(scanner.Text(), "")
	for _, cupStringValue := range cupStringValues {
		cupValue, err := strconv.Atoi(cupStringValue)
		if err != nil {
			return nil, err
		}

		cups = append(cups, cupValue)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cups, nil
}
