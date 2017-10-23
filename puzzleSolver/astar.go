package main

import (
	priorityQueue "github.com/jupp0r/go-priority-queue"
)

func astar(start, solution board) {
	fringe := priorityQueue.New()

	startMD := start.manhattanDistance(solution)

	fringe.Insert(start, startMD)

	inter, err := fringe.Pop()

	current := inter.(board)

	if err == nil {
		for !current.equals(solution) {
			moves := current.moves()

			for i := 0; moves.Len() > 0; i++ {
				next := moves.Remove(moves.Front()).(board)

				md := next.manhattanDistance(solution)
				fringe.Insert(next, md)
			}

			inter, err := fringe.Pop()

			if err == nil {
				current = inter.(board)
			}
		}

		current.printParents()
	}
}
