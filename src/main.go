package main

import (
	"fmt"
	"os"
	"time"
)

func usage() string {
	return fmt.Sprintf(
		"usage: taquin ALGORITHM <pathToFile>\nAvailable algorithms:\n%s%s%s%s",
		" - bfs\n",
		" - misplaced\n",
		" - manhattan\n",
		" - misplaced+manhattan\n",
	)
}

var totalNodesExplored = 0

func main() {
	var puzzle *Puzzle
	var err error

	puzzle, err = handleArgs()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	if !puzzle.isSolvable() {
		fmt.Println("️Puzzle is not solvable... 😔")
		return
	}
	start := time.Now()
	pathToSolution := puzzle.Solve()
	defer fmt.Printf("Total nodes explored: %v\nPath to solution: %v\nTook %v to resolve puzzle", totalNodesExplored, pathToSolution, time.Since(start))
	return
}

func handleArgs() (*Puzzle, error) {
	argIdx := 1

	if len(os.Args) == 1 {
		return nil, fmt.Errorf(usage())
	}

	if len(os.Args) == argIdx {
		return nil, fmt.Errorf(usage())
	}
	switch os.Args[argIdx] {
	case MANHATTAN:
		selectedAlgorithm = MANHATTAN
	case BFS:
		selectedAlgorithm = BFS
	case MISPLACED:
		selectedAlgorithm = MISPLACED
	case MISPLACEDMANHATTAN:
		selectedAlgorithm = MISPLACEDMANHATTAN
	default:
		return nil, fmt.Errorf(usage())
	}

	if len(os.Args) == argIdx+2 {
		return Parse(os.Args[argIdx+1])
	} else {
		return nil, fmt.Errorf(usage())
	}

	return nil, nil
}

