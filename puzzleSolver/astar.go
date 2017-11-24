package main

import (
	priorityQueue "github.com/jupp0r/go-priority-queue"
)

func astar(start, solution board) int {
	fringe := priorityQueue.New()

	startMD := start.manhattanDistance(solution)

	fringe.Insert(start, startMD)

	inter, err := fringe.Pop()

	current := inter.(board)
	count := 0

	if err == nil {
		for !current.equals(solution) {
			count++

			moves := current.moves()

			for i := 0; moves.Len() > 0; i++ {
				next := moves.Remove(moves.Front()).(board)

				md := next.manhattanDistance(solution)
				numParents := next.numberParents()

				fringe.Insert(next, md+numParents)
			}

			inter, err := fringe.Pop()

			if err == nil {
				current = inter.(board)
			}
		}
		return count
		//current.printParents()
	}
	return 0
}
