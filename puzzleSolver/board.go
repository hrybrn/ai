package main

import "container/list"

type point struct {
	x, y int
}

func (p point) equals(other point) bool {
	return p.x == other.x && p.y == other.y
}

func (p point) valid() bool {
	return p.x >= 0 && p.y >= 0 && p.x <= 3 && p.y <= 3
}

type board struct {
	a, b, c, agent point
}

func (b board) equals(other board) bool {
	return b.a.equals(other.a) && b.b.equals(other.b) && b.c.equals(other.c) && b.agent.equals(other.agent)
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

			moves.PushFront(board{pointa, pointb, pointc, agent})
		}
	}

	return moves
}
