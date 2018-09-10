package main

import (
	"fmt"
)

var selectedAlgorithm string

func (puzzle Puzzle) Solve() []int {

	if selectedAlgorithm == MANHATTAN || selectedAlgorithm == MISPLACED || selectedAlgorithm == MISPLACEDMANHATTAN {
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
	states := NewImplicitHeapMin(true)
	newInstance := puzzle.getCopy()
	states.Push(0, newInstance)
	for states.Len() > 0 {
		state, ok := states.Pop()
		s := state.(*Puzzle)
		if !ok {
			fmt.Errorf("error")
		}
		if s.isGoalState(){
			return s.path
		}
		children := s.visit()
		for i := 0; i < len(children); i++ {
			totalNodesExplored += 1
			var child = children[i]
			child.distance = len(child.path) + child.getCost()
			states.Push(child.distance, child)
		}
	}
	return nil
}
