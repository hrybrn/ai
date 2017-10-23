package main

import (
	"container/list"
	"fmt"

	shuffle "github.com/shogo82148/go-shuffle"
)

func depthFirst(start, solution board) {
	fringe := list.New()
	fringe.PushFront(start)

	count := 0

	current := fringe.Remove(fringe.Front()).(board)

	for !current.equals(solution) {
		count++

		moves := current.moves()

		*moves = shuffleList(moves)

		for i := 0; moves.Len() > 0; i++ {
			next := moves.Remove(moves.Front())
			fringe.PushFront(next)
			//fmt.Println(count, next.agent.x, next.agent.y)
		}

		if count%1000000 == 0 {
			fmt.Println(count)
		}

		current = fringe.Remove(fringe.Front()).(board)
	}

	fmt.Println(count)
	current.printParents()
}

func shuffleList(input *list.List) list.List {
	size := input.Len()
	intermediate := make([]board, size)

	for i := 0; i < size; i++ {
		current := input.Front()
		input.Remove(current)
		intermediate[i] = current.Value.(board)
	}

	slice := intermediate[0:len(intermediate)]
	shuffle.Slice(slice)

	for i := 0; i < size; i++ {
		input.PushFront(slice[i])
	}

	return *input
}
