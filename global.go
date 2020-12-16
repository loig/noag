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

// characteristics of the generated automata
const (
	automatonMinNumStates           = 5  // must be at least 1
	automatonMaxNumStates           = 10 // must be greater or equal than the above
	automatonMinGoalStates          = 2  // must be at least 1
	automatonMaxGoalStates          = 7  // must be greater or equal than the above
	automatonMinLabels              = 3  // must be at least 1
	automatonMaxLabels              = 8  // must be greater or equal than the above
	automatonMinPrivateLabels       = 2  // can be 0
	automatonMaxPrivateLabels       = 7  // notice that the number of private labels is at most the number of labels minus 1 because at least one label must be shared
	automatonMinTransitionsPerState = 2  // could be 0, will never be more than the number of labels in an automaton
	automatonMinTransitions         = 10 // could be 0, will neve be more than the number of states * the number of labels in an automaton
	graphNumAutomata                = 5
)

// names of things
const (
	automatonName = "A"
	stateName     = "s"
	actionName    = "a"
)
