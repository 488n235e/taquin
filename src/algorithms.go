package main

import (
	"container/heap"
)

var selectedAlgorithm string

func (puzzle Puzzle) Solve() []int {

	if selectedAlgorithm == MANHATTAN || selectedAlgorithm == MISPLACED {
		return puzzle.AStar()
	} else {
		return puzzle.BFS()
	}
	return nil
}

func (puzzle *Puzzle) BFS() []int {

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

func (puzzle *Puzzle) AStar() []int {
	states := make(PriorityQueue, 0)
	newInstance := puzzle.getCopy()
	states = append(states, newInstance)
	for len(states) > 0 {
		state := heap.Pop(&states).(*Puzzle)
		if state.isGoalState() {
			return state.path
		}
		children := state.visit()
		for i := 0; i < len(children); i++ {
			totalNodesExplored += 1
			var child = children[i]
			child.distance = len(child.path) + child.getCost()
			states.Push(child)
		}
	}
	return nil
}
