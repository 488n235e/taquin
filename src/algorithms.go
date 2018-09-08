package main

import (
	"container/heap"
)

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

func (puzzle *Puzzle) solveAStarManhattan() []int {
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
			child.distance = len(child.path) + child.getManhattanDistance()
			states.Push(child)
		}
	}
	return nil
}