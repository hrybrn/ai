package main

import (
	"container/list"
)

func iterativeDeepening(start, solution board) int {

	startingDepth := 1.0

	count, found := iterative(start, solution, startingDepth)

	for !found {
		startingDepth++
		currentCount, currentFound := iterative(start, solution, startingDepth)

		count += currentCount
		found = currentFound
	}

	return count
}

func iterative(start, solution board, depthLimit float64) (int, bool) {
	fringe := list.New()
	fringe.PushFront(start)

	count := 0

	current := start

	for fringe.Len() > 0 && !current.equals(solution) {
		current = fringe.Remove(fringe.Front()).(board)
		count++

		depth := current.numberParents() + 1

		moves := current.moves()

		if depth < depthLimit {
			for i := 0; moves.Len() > 0; i++ {
				next := moves.Remove(moves.Front()).(board)

				fringe.PushFront(next)
			}

		}

		//current.printParents()
	}

	return count, (current.equals(solution))
}
