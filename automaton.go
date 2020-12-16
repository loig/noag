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
Generate an automaton
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
	numLabelsUsed := 0
	allLabelsUsed := false
	labelsUsedPerState := make([][]bool, numStates)
	for i := 0; i < numStates; i++ {
		labelsUsedPerState[i] = make([]bool, len(labels))
	}
	numLabelsUsedPerState := make([]int, numStates)
	blockedStates := make([]bool, numStates)
	numBlockedStates := 0
	minNumTransitions := automatonMinTransitions
	if minNumTransitions > numStates*len(labels) {
		minNumTransitions = numStates * len(labels)
	}
	enoughTransitions := false
	minNumTransitionsPerState := automatonMinTransitionsPerState
	if minNumTransitionsPerState > len(labels) {
		minNumTransitionsPerState = len(labels)
	}
	enoughTransitionsStates := make([]bool, numStates)
	numEnoughTransitionsStates := 0
	enoughTransitionsPerState := false
	for !allStatesReached || !allLabelsUsed ||
		!enoughTransitions || !enoughTransitionsPerState {
		// choose a reachable state
		var state int
		var stateNum int
		if allStatesReached {
			stateNum = rand.Intn(numStates - numBlockedStates)
		} else {
			stateNum = rand.Intn(nextStatePos - numBlockedStates)
		}
		stateCount := 0
		statePos := 0
		for stateCount <= stateNum {
			if !blockedStates[statePos] {
				stateCount++
				state = statePos
			}
			statePos++
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
		labelNum := rand.Intn(len(labels) - numLabelsUsedPerState[state])
		labelCount := 0
		labelPos := 0
		var label int
		for labelCount <= labelNum {
			if !labelsUsedPerState[state][labelPos] {
				labelCount++
				label = labels[labelPos]
			}
			labelPos++
		}
		labelsUsedPerState[state][labelPos-1] = true
		numLabelsUsedPerState[state]++
		if numLabelsUsedPerState[state] >= len(labels) {
			blockedStates[state] = true
			numBlockedStates++
		}
		if !enoughTransitionsStates[state] && numLabelsUsedPerState[state] >= minNumTransitionsPerState {
			enoughTransitionsStates[state] = true
			numEnoughTransitionsStates++
			enoughTransitionsPerState = numEnoughTransitionsStates >= numStates
		}
		if !usedLabels[labelPos-1] {
			usedLabels[labelPos-1] = true
			numLabelsUsed++
			allLabelsUsed = numLabelsUsed >= len(labels)
		}
		// add a transition between the two states
		transitions = append(transitions, transition{
			from:  state,
			to:    nextState,
			label: label,
		})
		// count this transition
		enoughTransitions = len(transitions) >= minNumTransitions
	}

	return automaton{
		numStates:   numStates,
		labels:      labels,
		goalStates:  goalStates,
		transitions: transitions,
	}
}
