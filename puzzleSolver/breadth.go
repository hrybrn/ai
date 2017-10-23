package main

import (
	"container/list"
	"fmt"
)

func breadthFirst(start, solution board) {
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

		if count%1000000 == 0 {
			fmt.Println(count)
		}
		current = fringe.Remove(fringe.Front()).(board)
	}

	fmt.Println(count)
	current.printParents()
}
