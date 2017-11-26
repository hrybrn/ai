package main

import (
	"math/rand"
)

func depthFirst(start, solution board) (int, int) {
	count := 0

	current := start

	for !current.equals(solution) {
		count++

		moves := current.moves()

		index := rand.Intn(moves.Len())

		for index > 0 {
			moves.Remove(moves.Front())
			index--
		}

		current = moves.Front().Value.(board)
	}

	return count, 1
}
