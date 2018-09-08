package main

import (
	"fmt"
	"os"
	"time"
)

func usage() string {
	return fmt.Sprintf(
		"usage: taquin HEURISTIC [FILE]\nAvailable heuristics:\n%s%s",
		" - manhattan\n",
		" - bfs\n",
	)
}

func createPuzzle(input [][]int) *Puzzle {
	return &Puzzle{board: input, dimension: len(input)}
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

	start := time.Now()
	pathToSolution := puzzle.Solve()
	defer fmt.Printf("Total nodes explored: %v\nPath to solution: %v\nTook %v to resolve puzzle", totalNodesExplored, pathToSolution, time.Since(start))
	return
}

func handleArgs() (*Puzzle, error) {
	var input = [][]int{{5, 1, 3, 4}, {2, 6, 7, 8}, {9, 10, 0, 12}, {13, 14, 11, 15}}
	//input := [][]int{{13, 2, 3, 12}, {9, 11, 1, 10}, {0, 6, 4, 14}, {15, 8, 7, 5}}
	//input := [][]int{{1, 3, 7, 4}, {6, 0, 2, 8}, {5, 9, 10, 11}, {13, 14, 15, 12}}

	//	input := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 0, 15}}
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
	default:
		return nil, fmt.Errorf(usage())
	}

	if len(os.Args) == argIdx+1 {
		return createPuzzle(input), nil
	}

	return nil, nil
}

