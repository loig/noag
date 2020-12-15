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
	"encoding/json"
	"fmt"
	"log"
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
	Content []JSONTransition
}

type JSONTransition struct {
	From  string
	To    string
	Label string
}

func (jsonTrans JSONTransitions) MarshalJSON() ([]byte, error) {

	var asJSON string
	asJSON += "{"
	for i, transition := range jsonTrans.Content {
		if i != 0 {
			asJSON += ","
		}
		asJSON += "\"" + transition.From + "\":"
		asJSON += "{"
		asJSON += "\"" + transition.Label + "\":"
		asJSON += "\"" + transition.To + "\""
		asJSON += "}"
	}
	asJSON += "}"

	return []byte(asJSON), nil
}

func (a automaton) toJSON(id int) string {

	// Name
	var jAutomaton JSONAutomaton
	jAutomaton.Name = fmt.Sprint(automatonName, id)

	// States
	jAutomaton.States = make([]string, a.numStates)
	for i := 0; i < a.numStates; i++ {
		jAutomaton.States[i] = fmt.Sprint(stateName, i)
	}

	// InputSymbols
	jAutomaton.InputSymbols = make([]string, a.maxLabel-a.minLabel+1)
	for i := a.minLabel; i < a.maxLabel+1; i++ {
		jAutomaton.InputSymbols[i-a.minLabel] = fmt.Sprint(actionName, i)
	}

	// Transitions
	jAutomaton.Transitions.Content = make([]JSONTransition, len(a.transitions))
	for i, transition := range a.transitions {
		jAutomaton.Transitions.Content[i] = JSONTransition{
			From:  fmt.Sprint(stateName, transition.from),
			To:    fmt.Sprint(stateName, transition.to),
			Label: fmt.Sprint(actionName, transition.label),
		}
	}

	// InitialState
	jAutomaton.InitialState = fmt.Sprint(stateName, 0)

	//FinalStates
	jAutomaton.FinalStates = make([]string, len(a.goalStates))
	for i, stateNum := range a.goalStates {
		jAutomaton.FinalStates[i] = fmt.Sprint(stateName, stateNum)
	}

	// Get result as JSON
	out, err := json.Marshal(jAutomaton)
	if err != nil {
		log.Panic(err)
	}

	return string(out)

}
