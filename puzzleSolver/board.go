package main

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
)

type point struct {
	x, y, xLimit, yLimit int
}

func (p point) equals(other point) bool {
	return p.x == other.x && p.y == other.y
}

func (p point) valid() bool {
	return p.x >= 0 && p.y >= 0 && p.x <= p.xLimit && p.y <= p.yLimit
}

func (p point) manhattanDistance(other point) float64 {
	return math.Abs(float64(p.x-other.x)) + math.Abs(float64(p.y-other.y))
}

type board struct {
	a, b, c, agent point
	parent         *board
}

func (b board) equals(other board) bool {
	return b.a.equals(other.a) && b.b.equals(other.b) && b.c.equals(other.c) //&& b.agent.equals(other.agent)
}

func (b board) manhattanDistance(other board) float64 {
	return b.a.manhattanDistance(other.a) + b.b.manhattanDistance(other.b) + b.c.manhattanDistance(other.c) //+ b.agent.manhattanDistance(other.agent)
}

func (b board) moves() *list.List {
	moves := list.New()

	points := make([]point, 4)
	points[0] = point{b.agent.x - 1, b.agent.y, b.agent.xLimit, b.agent.yLimit}
	points[1] = point{b.agent.x + 1, b.agent.y, b.agent.xLimit, b.agent.yLimit}
	points[2] = point{b.agent.x, b.agent.y - 1, b.agent.xLimit, b.agent.yLimit}
	points[3] = point{b.agent.x, b.agent.y + 1, b.agent.xLimit, b.agent.yLimit}

	for i := 0; i < 4; i++ {
		if points[i].valid() {
			pointa := b.a
			pointb := b.b
			pointc := b.c
			agent := points[i]

			if agent.equals(pointa) {
				pointa = b.agent
			} else if agent.equals(pointb) {
				pointb = b.agent
			} else if agent.equals(pointc) {
				pointc = b.agent
			}

			moves.PushFront(board{pointa, pointb, pointc, agent, &b})
		}
	}

	return moves
}

func (b board) print() string {
	output := make([][]string, b.agent.xLimit+1)
	for i := 0; i <= b.agent.xLimit; i++ {
		output[i] = make([]string, b.agent.yLimit+1)
	}
	for i := 0; i <= b.agent.xLimit; i++ {
		for j := 0; j <= b.agent.yLimit; j++ {
			output[i][j] = " "
		}
	}
	output[b.a.x][b.a.y] = "a"
	output[b.b.x][b.b.y] = "b"
	output[b.c.x][b.c.y] = "c"
	output[b.agent.x][b.agent.y] = "d"

	stringOutput := "|--------|\n"

	for i := 0; i <= b.agent.xLimit; i++ {
		stringOutput += "|"
		for j := 0; j <= b.agent.yLimit; j++ {
			stringOutput += output[i][j] + " "
		}

		stringOutput += "|\n"
	}

	return stringOutput + "|--------|\n"
}

func (b board) printParents() {
	current := &b

	fullOutput := list.New()

	count := 0

	for current != nil {
		fullOutput.PushBack(current.print())
		current = current.parent
	}

	for fullOutput.Len() > 0 {
		orderedOut := fullOutput.Remove(fullOutput.Back())
		fmt.Println("move " + strconv.Itoa(count) + "\n")
		count++
		fmt.Println(orderedOut)
	}
}

func (b board) numberParents() float64 {
	current := &b

	count := 0.0

	for current != nil {
		count++
		current = current.parent
	}

	return count
}

func (b board) explored(bs []board) bool {
	for _, element := range bs {
		if b.equals(element) {
			return true
		}
	}

	return false
}
