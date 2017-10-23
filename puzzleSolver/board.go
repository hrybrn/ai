package main

import (
	"container/list"
	"fmt"
	"math"
)

type point struct {
	x, y int
}

func (p point) equals(other point) bool {
	return p.x == other.x && p.y == other.y
}

func (p point) valid() bool {
	return p.x >= 0 && p.y >= 0 && p.x <= 3 && p.y <= 3
}

func (p point) manhattanDistance(other point) float64 {
	return math.Abs(float64(p.x-other.x)) + math.Abs(float64(p.y-other.y))
}

type board struct {
	a, b, c, agent point
	parent         *board
}

func (b board) equals(other board) bool {
	return b.a.equals(other.a) && b.b.equals(other.b) && b.c.equals(other.c) && b.agent.equals(other.agent)
}

func (b board) manhattanDistance(other board) float64 {
	return b.a.manhattanDistance(other.a) + b.b.manhattanDistance(other.b) + b.c.manhattanDistance(other.c) //+ b.agent.manhattanDistance(other.agent)
}

func (b board) moves() *list.List {
	moves := list.New()

	points := make([]point, 4)
	points[0] = point{b.agent.x - 1, b.agent.y}
	points[1] = point{b.agent.x + 1, b.agent.y}
	points[2] = point{b.agent.x, b.agent.y - 1}
	points[3] = point{b.agent.x, b.agent.y + 1}

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

func (b board) printParents() {
	current := &b

	fullOutput := list.New()

	for current != nil {
		output := [4][4]string{
			{" ", " ", " ", " "},
			{" ", " ", " ", " "},
			{" ", " ", " ", " "},
			{" ", " ", " ", " "}}

		output[current.a.x][current.a.y] = "a"
		output[current.b.x][current.b.y] = "b"
		output[current.c.x][current.c.y] = "c"
		output[current.agent.x][current.agent.y] = "d"

		stringOutput := ""

		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				stringOutput += output[j][3-i] + " "
			}

			stringOutput += "\n"
		}

		fullOutput.PushBack(stringOutput)
		current = current.parent
	}

	for fullOutput.Len() > 0 {
		orderedOut := fullOutput.Remove(fullOutput.Back())
		fmt.Println(orderedOut)
	}

}
