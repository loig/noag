/*
noag, generation of networks of automata
Copyright (C) 2020 Loïg Jezequel

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"math/rand"
)

/*
Structure for representing automata.
States are positive integers from 0 to numStates - 1.
The initial state is always 0.
*/
type automaton struct {
	numStates   int
	labels      []int
	goalStates  []int
	transitions []transition
}

type transition struct {
	from  int
	to    int
	label int
}

/*
Generate an automaton by performing a search from its initial state
*/
func genAutomaton(labels []int) automaton {

	// number of states
	numStates := rand.Intn(automatonMaxNumStates-automatonMinNumStates+1) + automatonMinNumStates

	// number of goal states
	numGoalStates := rand.Intn(automatonMaxGoalStates-automatonMinGoalStates+1) + automatonMinGoalStates
	if numGoalStates > numStates {
		numGoalStates = numStates
	}

	// set of goal states
	allStates := make([]int, numStates)
	for i := 0; i < numStates; i++ {
		allStates[i] = i
	}
	rand.Shuffle(numStates, func(i, j int) {
		allStates[i], allStates[j] = allStates[j], allStates[i]
	})
	goalStates := make([]int, numGoalStates)
	copy(goalStates, allStates[:numGoalStates])

	// set of transitions
	transitions := make([]transition, 0)
	nextStatePos := 1
	allStatesReached := false
	usedLabels := make([]bool, len(labels))
	numLabelUsed := 0
	allLabelsUsed := false
	for !allStatesReached || !allLabelsUsed {
		// choose a reachable state
		var state int
		if allStatesReached {
			state = rand.Intn(numStates)
		} else {
			state = rand.Intn(nextStatePos)
		}
		// choose a state to reach from it
		var nextState int
		if allStatesReached {
			nextState = allStates[nextStatePos]
		} else {
			nextState = nextStatePos
		}
		nextStatePos++
		if nextStatePos >= numStates {
			rand.Shuffle(numStates, func(i, j int) {
				allStates[i], allStates[j] = allStates[j], allStates[i]
			})
			nextStatePos = 0
			allStatesReached = true
		}
		// choose a label
		// WARNING: this does not guarantee determinism of automata, does not
		// avoid duplicate transitions, and may even lead to infinite time
		// for automaton generation
		labelNum := rand.Intn(len(labels))
		label := labels[labelNum]
		if !usedLabels[labelNum] {
			usedLabels[labelNum] = true
			numLabelUsed++
			if numLabelUsed >= len(labels) {
				allLabelsUsed = true
			}
		}
		// add a transition between the two states
		transitions = append(transitions, transition{
			from:  state,
			to:    nextState,
			label: label,
		})
	}

	return automaton{
		numStates:   numStates,
		labels:      labels,
		goalStates:  goalStates,
		transitions: transitions,
	}
}
