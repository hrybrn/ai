package main

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

	//breadthFirst(start, solution)
	concurrentBreadthFirst(start, solution)
	//depthFirst(start, solution)
	//astar(start, solution)
}
