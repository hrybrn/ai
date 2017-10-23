package main

import (
	"container/list"
	"fmt"

	shuffle "github.com/shogo82148/go-shuffle"
)

func main() {
	startA := point{0, 0}
	startB := point{1, 0}
	startC := point{2, 0}
	startAgent := point{3, 0}

	start := board{startA, startB, startC, startAgent, nil}

	solutionA := point{1, 2}
	solutionB := point{1, 1}
	solutionC := point{1, 0}
	solutionAgent := point{3, 0}

	solution := board{solutionA, solutionB, solutionC, solutionAgent, nil}

	breadthFirst(start, solution)
	//depthFirst(start, solution)
}

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
