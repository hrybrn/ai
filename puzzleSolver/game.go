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

	for i := 13; i >= 0; i-- {
		start := loadFromFile("difficulty/" + strconv.Itoa(i) + ".txt")

		nodesExpanded, queueAtLargest := breadthFirst(start, solution)

		fmt.Println(strconv.Itoa(13-i) + "," + strconv.Itoa(nodesExpanded) + "," + strconv.Itoa(queueAtLargest))
	}

	fmt.Println()

	for i := 13; i >= 0; i-- {
		start := loadFromFile("difficulty/" + strconv.Itoa(i) + ".txt")

		nodesExpanded, queueAtLargest := breadthFirstGraph(start, solution)

		fmt.Println(strconv.Itoa(13-i) + "," + strconv.Itoa(nodesExpanded) + "," + strconv.Itoa(queueAtLargest))
	}

	fmt.Println()

	for i := 13; i >= 0; i-- {
		start := loadFromFile("difficulty/" + strconv.Itoa(i) + ".txt")

		nodesExpanded, queueAtLargest := astar(start, solution)

		fmt.Println(strconv.Itoa(13-i) + "," + strconv.Itoa(nodesExpanded) + "," + strconv.Itoa(queueAtLargest))
	}

	fmt.Println()

	for i := 13; i >= 0; i-- {
		start := loadFromFile("difficulty/" + strconv.Itoa(i) + ".txt")

		nodesExpanded, queueAtLargest := astarGraph(start, solution)

		fmt.Println(strconv.Itoa(13-i) + "," + strconv.Itoa(nodesExpanded) + "," + strconv.Itoa(queueAtLargest))
	}

	//start := loadFromFile("difficulty/0.txt")
	//_, _ = breadthFirst(start, solution)
	//_, _ = depthFirst(start, solution)
	//_, _ = astar(start, solution)
	//_, _ = iterativeDeepening(start, solution)

	//nodesExpanded, queueAtLargest := breadthFirstGraph(start, solution)
	//fmt.Println(strconv.Itoa(nodesExpanded) + "," + strconv.Itoa(queueAtLargest))
	//nodesExpanded, queueAtLargest = astarGraph(start, solution)
	//fmt.Println(strconv.Itoa(nodesExpanded) + "," + strconv.Itoa(queueAtLargest))
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
