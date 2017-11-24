package main

import (
	"container/list"
	"math/rand"

	shuffle "github.com/shogo82148/go-shuffle"
)

func depthFirst(start, solution board) int {
	count := 0

	current := start

	for !current.equals(solution) {
		count++

		moves := current.moves()

		index := rand.Intn(moves.Len())

		for index > 0 {
			moves.Remove(moves.Front())
			index--
		}

		current = moves.Front().Value.(board)
	}

	return count
	//current.printParents()
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
