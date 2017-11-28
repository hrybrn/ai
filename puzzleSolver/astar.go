package main

import (
	priorityQueue "github.com/jupp0r/go-priority-queue"
)

func astar(start, solution board) (int, int) {
	fringe := priorityQueue.New()

	startMD := start.manhattanDistance(solution)

	fringe.Insert(start, startMD)

	largest := 1

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

			if largest < fringe.Len() {
				largest = fringe.Len()
			}

			inter, err := fringe.Pop()

			if err == nil {
				current = inter.(board)
			}
		}
		//current.printParents()
		return count, largest
	}
	return 0, 0
}

func astarGraph(start, solution board) (int, int) {
	explored := make([]board, 1)
	fringe := priorityQueue.New()

	largest := 1

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
			explored = append(explored, current)

			if largest < (fringe.Len() + len(explored)) {
				largest = fringe.Len() + len(explored)
			}

			for current.explored(explored) {
				inter, _ := fringe.Pop()

				current = inter.(board)
			}

		}
		return count, largest
	}
	return 0, 0
}
