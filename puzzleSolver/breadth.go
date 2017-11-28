package main

import (
	"container/list"
)

func breadthFirst(start, solution board) (int, int) {
	fringe := list.New()
	fringe.PushFront(start)

	largest := 1
	count := 0

	current := fringe.Remove(fringe.Front()).(board)

	for !current.equals(solution) {
		count++

		moves := current.moves()

		for i := 0; moves.Len() > 0; i++ {
			next := moves.Remove(moves.Front()).(board)

			fringe.PushBack(next)
		}

		if largest < fringe.Len() {
			largest = fringe.Len()
		}

		current = fringe.Remove(fringe.Front()).(board)

	}

	return count, largest
}

func breadthFirstGraph(start, solution board) (int, int) {
	explored := make([]board, 1)

	fringe := list.New()
	fringe.PushFront(start)

	largest := 1
	count := 0

	current := start

	for !current.equals(solution) {
		count++

		explored = append(explored, current)

		moves := current.moves()

		for i := 0; moves.Len() > 0; i++ {
			next := moves.Remove(moves.Front()).(board)

			if !next.explored(explored) {
				fringe.PushBack(next)
			}
		}

		if largest < (fringe.Len() + len(explored)) {
			largest = fringe.Len() + len(explored)
		}

		current = fringe.Remove(fringe.Front()).(board)
	}

	return count, largest
}
