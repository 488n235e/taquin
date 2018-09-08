package main

import (
	"fmt"
	"math"
	"time"
)

const (
	LEFT  = "left"
	RIGHT = "right"
	UP    = "up"
	DOWN  = "down"
)

type Puzzle struct {
	board     [][]int
	path      []int
	dimension int
	lastMove  int
}

func createPuzzle(input [][]int) *Puzzle {
	return &Puzzle{board: input, dimension: len(input)}
}

func (puzzle Puzzle) getBlankSpacePosition() []int {
	for i := 0; i < puzzle.dimension; i++ {
		for j := 0; j < puzzle.dimension; j++ {
			if puzzle.board[i][j] == 0 {
				return []int{i, j}
			}
		}
	}
	return nil
}

func (puzzle Puzzle) swap(i1 int, j1 int, i2 int, j2 int) {
	puzzle.board[i1][j1], puzzle.board[i2][j2] = puzzle.board[i2][j2], puzzle.board[i1][j1]
}

func (puzzle Puzzle) getMove(piece int) string {
	var blankSpacePosition = puzzle.getBlankSpacePosition()
	var line = blankSpacePosition[0]
	var column = blankSpacePosition[1]
	switch {
	case line > 0 && piece == puzzle.board[line-1][column]:
		return DOWN
	case line < puzzle.dimension-1 && piece == puzzle.board[line+1][column]:
		return UP
	case column > 0 && piece == puzzle.board[line][column-1]:
		return RIGHT
	case column < puzzle.dimension-1 && piece == puzzle.board[line][column+1]:
		return LEFT
	}
	return ""
}

func (puzzle *Puzzle) isGoalState() bool {

	for i := 0; i < puzzle.dimension; i++ {
		for j := 0; j < puzzle.dimension; j++ {
			piece := puzzle.board[i][j]
			if piece != 0 {
				originalLine := int(math.Floor(float64((piece - 1) / puzzle.dimension)))
				originalColumn := (piece - 1) % puzzle.dimension
				if i != originalLine || j != originalColumn {
					return false
				}
			}
		}
	}
	return true
}

func (puzzle Puzzle) getCopy() *Puzzle {
	newPuzzle := new(Puzzle)
	n, m := len(puzzle.board), len(puzzle.board[0])

	newBoardDuplicate := make([][]int, n)
	data := make([]int, n*m)
	for i := range puzzle.board {
		start := i * m
		end := start + m
		newBoardDuplicate[i] = data[start:end:end]
		copy(newBoardDuplicate[i], puzzle.board[i])
	}
	newPathDuplicate := make([]int, len(puzzle.path))
	copy(newPathDuplicate, puzzle.path)

	newPuzzle.board = newBoardDuplicate
	newPuzzle.path = newPathDuplicate
	newPuzzle.dimension = puzzle.dimension
	return newPuzzle
}

func (puzzle Puzzle) getAllowedMoves() []int {
	var allowedMoves = make([]int, 0)

	for i := 0; i < puzzle.dimension; i++ {
		for j := 0; j < puzzle.dimension; j++ {
			piece := puzzle.board[i][j]
			if puzzle.getMove(piece) != "" {
				allowedMoves = append(allowedMoves, piece)
			}
		}
	}
	return allowedMoves
}

func (puzzle Puzzle) visit() []*Puzzle {
	var children = make([]*Puzzle, 0)
	allowedMoves := puzzle.getAllowedMoves()

	for i := 0; i < len(allowedMoves); i++ {
		var move = allowedMoves[i]
		if move != puzzle.lastMove {
			var newInstance = puzzle.getCopy()
			newInstance.move(move)
			newInstance.lastMove = move
			newInstance.path = append(newInstance.path, move)
			children = append(children, newInstance)
		}
	}
	return children
}

func (puzzle Puzzle) move(piece int) string {
	var move = puzzle.getMove(piece)
	if move != "" {
		blankSpacePosition := puzzle.getBlankSpacePosition()
		line, column := blankSpacePosition[0], blankSpacePosition[1]
		switch move {
		case LEFT:
			puzzle.swap(line, column, line, column+1)
		case RIGHT:
			puzzle.swap(line, column, line, column-1)
		case UP:
			puzzle.swap(line, column, line+1, column)
		case DOWN:
			puzzle.swap(line, column, line-1, column)
		}

		return move
	}
	return ""
}

var totalNodesExplored = 0

func main() {
	var puzzle *Puzzle

	//input := [][]int{{13, 2, 3, 12}, {9, 11, 1, 10}, {0, 6, 4, 14}, {15, 8, 7, 5}}
	input := [][]int{{1, 3, 7, 4}, {6, 0, 2, 8}, {5, 9, 10, 11}, {13, 14, 15, 12}}
	//input := [][]int{{5, 1, 3, 4}, {2, 6, 7, 8}, {9, 10, 0, 12}, {13, 14, 11, 15}}
//	input := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 0, 15}}

	puzzle = createPuzzle(input)

	start := time.Now()
	pathToSolution := puzzle.solveBFS()
	defer fmt.Printf("Total nodes explored: %v\nPath to solution: %v\nTook %v to resolve puzzle", totalNodesExplored, pathToSolution, time.Since(start))
	return
}

func (puzzle *Puzzle) solveBFS() []int {

	states := make([]*Puzzle, 0)
	states = append(states, puzzle.getCopy())
	for len(states) > 0 {
		state := states[0]
		totalNodesExplored += 1
		if state.isGoalState() {
			return state.path
		}
		states = states[1:]
		states = append(states, state.visit()...)
	}

	return nil
}


