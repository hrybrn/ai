package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

func main() {
	startA := point{0, 0}
	startB := point{1, 0}
	startC := point{2, 0}
	startAgent := point{3, 0}

	start := board{startA, startB, startC, startAgent}

	solutionA := point{0, 0}
	solutionB := point{1, 0}
	solutionC := point{3, 2}
	solutionAgent := point{3, 0}

	solution := board{solutionA, solutionB, solutionC, solutionAgent}

	//breadthFirst(start, solution)
	depthFirst(start, solution)
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
	fmt.Println(current.a.x, current.a.y)
	fmt.Println(current.b.x, current.b.y)
	fmt.Println(current.c.x, current.c.y)
	fmt.Println(current.agent.x, current.agent.y)
}

func depthFirst(start, solution board) {
	fringe := list.New()
	fringe.PushFront(start)

	count := 0

	current := fringe.Remove(fringe.Front()).(board)

	for !current.equals(solution) {
		count++

		moves := current.moves()

		shuffleList(moves)

		for i := 0; moves.Len() > 0; i++ {
			next := moves.Remove(moves.Front()).(board)
			fringe.PushFront(next)
			//fmt.Println(count, next.agent.x, next.agent.y)
		}

		if count%1000000 == 0 {
			fmt.Println(count)
		}

		current = fringe.Remove(fringe.Front()).(board)
	}

	fmt.Println(count)
	fmt.Println(current.a.x, current.a.y)
	fmt.Println(current.b.x, current.b.y)
	fmt.Println(current.c.x, current.c.y)
	fmt.Println(current.agent.x, current.agent.y)
}

func shuffleList(input *list.List) {
	intermediate := list.New()

	rand := rand.Intn(input.Len())

	for i := 0; i < rand; i++ {
		current := input.Remove(input.Front())
		intermediate.PushBack(current)
	}

	for intermediate.Len() > 0 {
		current := intermediate.Remove(intermediate.Front())
		input.PushFront(current)
	}
}
