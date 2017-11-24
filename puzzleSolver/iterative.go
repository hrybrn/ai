package main

import (
	"container/list"
	"fmt"
)

func iterativeDeepening(start, solution board) {
	fringe := list.New()
	fringe.PushFront(start)

	count := 0

	depthLimit := 2.0

	current := fringe.Remove(fringe.Front()).(board)

	for !current.equals(solution) {
		count++

		depth := current.numberParents() + 1

		moves := current.moves()

		if depth < depthLimit {
			for i := 0; moves.Len() > 0; i++ {
				next := moves.Remove(moves.Front()).(board)

				fringe.PushFront(next)
			}

		} else {
			for i := 0; moves.Len() > 0; i++ {
				next := moves.Remove(moves.Front()).(board)

				fringe.PushBack(next)
			}
		}

		if depth == depthLimit + 1 {
			depthLimit += 3
			fmt.Println(depthLimit)
		}

		if count%10000000 == 0 {
			fmt.Println(count)
		}
		current = fringe.Remove(fringe.Front()).(board)

	}

	fmt.Println(count)
	current.printParents()
}
