package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solution := loadFromFile("difficulty/solution.txt")

	for i := 15; i >= 0; i-- {
		start := loadFromFile("difficulty/" + strconv.Itoa(i) + ".txt")

		nodesExpanded := breadthFirst(start, solution)

		fmt.Println(strconv.Itoa(16-i) + "," + strconv.Itoa(nodesExpanded))
	}
	
	fmt.Println()
	fmt.Println()
	
	for i := 15; i >= 0; i-- {
		start := loadFromFile("difficulty/" + strconv.Itoa(i) + ".txt")

		nodesExpanded := depthFirst(start, solution)

		fmt.Println(strconv.Itoa(16-i) + "," + strconv.Itoa(nodesExpanded))
	}
	
	fmt.Println()
	fmt.Println()
	
	for i := 15; i >= 0; i-- {
		start := loadFromFile("difficulty/" + strconv.Itoa(i) + ".txt")

		nodesExpanded := astar(start, solution)

		fmt.Println(strconv.Itoa(16-i) + "," + strconv.Itoa(nodesExpanded))
	}
	
	fmt.Println()
	fmt.Println()
	
	for i := 15; i >= 0; i-- {
		start := loadFromFile("difficulty/" + strconv.Itoa(i) + ".txt")

		nodesExpanded := iterativeDeepening(start, solution)

		fmt.Println(strconv.Itoa(16-i) + "," + strconv.Itoa(nodesExpanded))
	}

	//breadthFirst(start, solution)
	//depthFirst(start, solution)
	//astar(start, solution)
	//iterativeDeepening(start, solution)
}

func loadFromFile(fileName string) board {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	row := 0

	scanner.Scan()
	current := scanner.Text()
	split := strings.Split(current, ",")

	xLimit, _ := strconv.Atoi(split[0])
	yLimit, _ := strconv.Atoi(split[1])

	a, b, c, agent := point{0, 0, 0, 0}, point{0, 0, 0, 0}, point{0, 0, 0, 0}, point{0, 0, 0, 0}

	for scanner.Scan() {
		current = scanner.Text()
		split := strings.Split(current, ",")

		for i := 0; i < len(split); i++ {
			if split[i] == "a" {
				a = point{row, i, yLimit, xLimit}
			} else if split[i] == "b" {
				b = point{row, i, yLimit, xLimit}
			} else if split[i] == "c" {
				c = point{row, i, yLimit, xLimit}
			} else if split[i] == "d" {
				agent = point{row, i, yLimit, xLimit}
			}
		}
		row++
	}
	file.Close()
	return board{a, b, c, agent, nil}
}
