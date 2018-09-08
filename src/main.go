package main

import (
	"fmt"
	"time"
)

const (
	LEFT  = "left"
	RIGHT = "right"
	UP    = "up"
	DOWN  = "down"
)

const (
	BFS        = "BFS"
	AStarManhattan = "A*: Manhattan distance"
)

func createPuzzle(input [][]int) *Puzzle {
	return &Puzzle{board: input, dimension: len(input)}
}

var totalNodesExplored = 0

func main() {
	var puzzle *Puzzle

	//input := [][]int{{13, 2, 3, 12}, {9, 11, 1, 10}, {0, 6, 4, 14}, {15, 8, 7, 5}}
	//input := [][]int{{1, 3, 7, 4}, {6, 0, 2, 8}, {5, 9, 10, 11}, {13, 14, 15, 12}}
	input := [][]int{{5, 1, 3, 4}, {2, 6, 7, 8}, {9, 10, 0, 12}, {13, 14, 11, 15}}
	//	input := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 0, 15}}

	puzzle = createPuzzle(input)

	start := time.Now()
	pathToSolution := puzzle.solveAStarManhattan()
	defer fmt.Printf("Total nodes explored: %v\nPath to solution: %v\nTook %v to resolve puzzle", totalNodesExplored, pathToSolution, time.Since(start))
	return
}
