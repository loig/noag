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
	"fmt"
	"math/rand"
)

/*
Structure for representing automata.
States are positive integers from 0 to numStates - 1.
The initial state is always 0.
*/
type automaton struct {
	numStates   uint
	goalStates  []uint
	transitions []transition
}

type transition struct {
	from  uint
	to    uint
	label uint
}

/*
Generate an automaton by performing a search from its initial state
*/
func genAutomaton(minLabel, maxLabel uint) automaton {

	// number of states
	numStates := rand.Intn(automatonMaxNumStates-automatonMinNumStates+1) + automatonMinNumStates

	// number of goal states
	numGoalStates := rand.Intn(automatonMaxGoalStates-automatonMinGoalStates+1) + automatonMinGoalStates
	if numGoalStates > numStates {
		numGoalStates = numStates
	}

	// set of goal states
	allStates := make([]uint, numStates)
	for i := 0; i < numStates; i++ {
		allStates[i] = uint(i)
	}
	rand.Shuffle(numStates, func(i, j int) {
		allStates[i], allStates[j] = allStates[j], allStates[i]
	})
	goalStates := make([]uint, numGoalStates)
	copy(goalStates, allStates[:numGoalStates])
	fmt.Println(goalStates)

	// set of transitions
	transitions := make([]transition, 0)
	nextStatePos := 1
	allStatesReached := false
	usedLabels := make([]bool, int(maxLabel-minLabel+1))
	numLabelUsed := 0
	allLabelsUsed := false
	for !allStatesReached || !allLabelsUsed {
		// choose a reachable state
		var state uint
		if allStatesReached {
			state = uint(rand.Intn(numStates))
		} else {
			state = uint(rand.Intn(nextStatePos))
		}
		// choose a state to reach from it
		var nextState uint
		if allStatesReached {
			nextState = allStates[nextStatePos]
		} else {
			nextState = uint(nextStatePos)
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
		label := uint(rand.Intn(int(maxLabel-minLabel+1)) + int(minLabel))
		if !usedLabels[int(label-minLabel)] {
			usedLabels[int(label-minLabel)] = true
			numLabelUsed++
			if numLabelUsed >= int(maxLabel-minLabel+1) {
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
		numStates:   uint(numStates),
		goalStates:  goalStates,
		transitions: transitions,
	}
}
