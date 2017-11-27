package main

import (
	"container/list"
)

func iterativeDeepening(start, solution board) (int, int) {

	startingDepth := 1.0

	count, largest, found := iterative(start, solution, startingDepth)

	for !found {
		startingDepth++
		currentCount, currentLargest, currentFound := iterative(start, solution, startingDepth)

		count += currentCount
		found = currentFound

		if largest < currentLargest {
			largest = currentLargest
		}
	}

	return count, largest
}

func iterative(start, solution board, depthLimit float64) (int, int, bool) {
	fringe := list.New()
	fringe.PushFront(start)

	largest := 1
	count := 0

	current := start

	for fringe.Len() > 0 && !current.equals(solution) {
		current = fringe.Remove(fringe.Front()).(board)

		depth := current.numberParents() + 1

		moves := current.moves()

		if depth < depthLimit {
			count++

			for i := 0; moves.Len() > 0; i++ {
				next := moves.Remove(moves.Front()).(board)

				fringe.PushFront(next)
			}

		}

		if largest < fringe.Len() {
			largest = fringe.Len()
		}

		//current.printParents()
	}

	return count, largest, (current.equals(solution))
}
