package main

import (
	"container/list"
	"fmt"
	"runtime"
)

func breadthFirst(start, solution board) int {
	fringe := list.New()
	fringe.PushFront(start)

	count := 0

	current := fringe.Remove(fringe.Front()).(board)

	for !current.equals(solution) {
		count++

		moves := current.moves()

		for i := 0; moves.Len() > 0; i++ {
			next := moves.Remove(moves.Front()).(board)

			fringe.PushBack(next)
		}

		current = fringe.Remove(fringe.Front()).(board)

	}

	return count
}
