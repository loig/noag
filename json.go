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
)

type JSONAutomaton struct {
	Name         string          `json:"name"`
	States       []string        `json:"states"`
	InputSymbols []string        `json:"input_symbols"`
	Transitions  JSONTransitions `json:"transitions"`
	InitialState string          `json:"initial_state"`
	FinalStates  []string        `json:"final_states"`
}

type JSONTransitions struct {
	Content map[string][]JSONTransition
}

type JSONTransition struct {
	To    string
	Label string
}

func (jsonTrans JSONTransitions) MarshalJSON() ([]byte, error) {

	var asJSON string
	asJSON += "{"
	firstLoop := true
	for from, transitions := range jsonTrans.Content {
		if !firstLoop {
			asJSON += ","
		} else {
			firstLoop = false
		}
		asJSON += "\"" + from + "\":"
		asJSON += "{"
		for i, transition := range transitions {
			if i != 0 {
				asJSON += ","
			}
			asJSON += "\"" + transition.Label + "\":"
			asJSON += "\"" + transition.To + "\""
		}
		asJSON += "}"
	}
	asJSON += "}"

	return []byte(asJSON), nil
}

func (a automaton) toJSON(id int) JSONAutomaton {

	// Name
	var jAutomaton JSONAutomaton
	jAutomaton.Name = fmt.Sprint(automatonName, id)

	// States
	jAutomaton.States = make([]string, a.numStates)
	for i := 0; i < a.numStates; i++ {
		jAutomaton.States[i] = fmt.Sprint(stateName, i)
	}

	// InputSymbols
	jAutomaton.InputSymbols = make([]string, len(a.labels))
	for i, label := range a.labels {
		jAutomaton.InputSymbols[i] = fmt.Sprint(actionName, label)
	}

	// Transitions
	jAutomaton.Transitions.Content = make(map[string][]JSONTransition)
	for _, transition := range a.transitions {
		from := fmt.Sprint(stateName, transition.from)
		jTransition := JSONTransition{
			To:    fmt.Sprint(stateName, transition.to),
			Label: fmt.Sprint(actionName, transition.label),
		}
		jTransitions, found := jAutomaton.Transitions.Content[from]
		if !found {
			jAutomaton.Transitions.Content[from] = []JSONTransition{jTransition}
		} else {
			jAutomaton.Transitions.Content[from] = append(jTransitions, jTransition)
		}
	}

	// InitialState
	jAutomaton.InitialState = fmt.Sprint(stateName, 0)

	//FinalStates
	jAutomaton.FinalStates = make([]string, len(a.goalStates))
	for i, stateNum := range a.goalStates {
		jAutomaton.FinalStates[i] = fmt.Sprint(stateName, stateNum)
	}

	return jAutomaton

}
