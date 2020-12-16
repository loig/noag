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
	"io/ioutil"
	"log"
)

type Configuration struct {
	MinNumStatesPerAutomaton        int
	MaxNumStatesPerAutomaton        int
	MinNumGoalStatesPerAutomaton    int
	MaxNumGoalStatesPerAutomaton    int
	MinNumLabelsPerAutomaton        int
	MaxNumLabelsPerAutomaton        int
	MinNumPrivateLabelsPerAutomaton int
	MaxNumPrivateLabelsPerAutomaton int
	MinNumTransitionsPerState       int
	MinNumTransitionsPerAutomaton   int
	NumAutomata                     int
}

func readConfigurationFile(file string) {
	log.Print("Reading configuration file ", file)

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error: cannot open configuration file ", file)
		//log.Panic(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Error: cannot parse configuration file ", file)
		//log.Panic(err)
	}

	// at least one state per automaton
	if config.MinNumStatesPerAutomaton < 1 {
		log.Print(
			"Warning, MinNumStatesPerAutomaton (",
			config.MinNumStatesPerAutomaton,
			") should be at least 1, automatically set to 1",
		)
		config.MinNumStatesPerAutomaton = 1
	}

	// max number of states greater than min number of states
	if config.MaxNumStatesPerAutomaton < config.MinNumStatesPerAutomaton {
		log.Print(
			"Warning: MaxNumStatesPerAutomaton (",
			config.MaxNumStatesPerAutomaton,
			") should be at least equal to MinNumStatesPerAutomaton (",
			config.MinNumStatesPerAutomaton,
			"), automatically set to ",
			config.MinNumStatesPerAutomaton,
		)
		config.MaxNumStatesPerAutomaton = config.MinNumStatesPerAutomaton
	}

	// at least one goal state per automaton
	if config.MinNumGoalStatesPerAutomaton < 1 {
		log.Print(
			"Warning, MinNumGoalStatesPerAutomaton (",
			config.MinNumGoalStatesPerAutomaton,
			") should be at least 1, automatically set to 1",
		)
		config.MinNumGoalStatesPerAutomaton = 1
	}

	// min number of goal states smaller than max number of states
	if config.MinNumGoalStatesPerAutomaton > config.MaxNumStatesPerAutomaton {
		log.Print(
			"Warning: MinNumGoalStatesPerAutomaton (",
			config.MinNumGoalStatesPerAutomaton,
			") should be smaller or equal than MaxNumStatesPerAutomaton (",
			config.MaxNumStatesPerAutomaton,
			"), automatically set to ",
			config.MaxNumStatesPerAutomaton,
		)
		config.MinNumGoalStatesPerAutomaton = config.MaxNumStatesPerAutomaton
	}

	// max number of goal states greater than min number of goal states
	if config.MaxNumGoalStatesPerAutomaton < config.MinNumGoalStatesPerAutomaton {
		log.Print(
			"Warning: MaxNumGoalStatesPerAutomaton (",
			config.MaxNumGoalStatesPerAutomaton,
			") should be at least equal to MinNumGoalStatesPerAutomaton (",
			config.MinNumGoalStatesPerAutomaton,
			"), automatically set to ",
			config.MinNumGoalStatesPerAutomaton,
		)
		config.MaxNumGoalStatesPerAutomaton = config.MinNumGoalStatesPerAutomaton
	}

	// max number of goal states smaller than max number of states
	if config.MaxNumGoalStatesPerAutomaton > config.MaxNumStatesPerAutomaton {
		log.Print(
			"Warning: MaxNumGoalStatesPerAutomaton (",
			config.MaxNumGoalStatesPerAutomaton,
			") should be smaller or equal than MaxNumStatesPerAutomaton (",
			config.MaxNumStatesPerAutomaton,
			"), automatically set to ",
			config.MaxNumStatesPerAutomaton,
		)
		config.MaxNumGoalStatesPerAutomaton = config.MaxNumStatesPerAutomaton
	}

	// at least one label per automaton
	if config.MinNumLabelsPerAutomaton < 1 {
		log.Print(
			"Warning, MinNumLabelsPerAutomaton (",
			config.MinNumLabelsPerAutomaton,
			") should be at least 1, automatically set to 1",
		)
		config.MinNumLabelsPerAutomaton = 1
	}

	// max number of labels greater than min number of labels
	if config.MaxNumLabelsPerAutomaton < config.MinNumLabelsPerAutomaton {
		log.Print(
			"Warning: MaxNumLabelsPerAutomaton (",
			config.MaxNumLabelsPerAutomaton,
			") should be at least equal to MinNumLabelsPerAutomaton (",
			config.MinNumLabelsPerAutomaton,
			"), automatically set to ",
			config.MinNumLabelsPerAutomaton,
		)
		config.MaxNumLabelsPerAutomaton = config.MinNumLabelsPerAutomaton
	}

	// at least zero private label per automaton
	if config.MinNumPrivateLabelsPerAutomaton < 0 {
		log.Print(
			"Warning, MinNumPrivateLabelsPerAutomaton (",
			config.MinNumPrivateLabelsPerAutomaton,
			") should not be negative, automatically set to 0",
		)
		config.MinNumPrivateLabelsPerAutomaton = 0
	}

	// min number of private labels smaller than max number of labels - 1
	if config.MinNumPrivateLabelsPerAutomaton > config.MaxNumLabelsPerAutomaton-1 {
		log.Print(
			"Warning: MinNumPrivateLabelsPerAutomaton (",
			config.MinNumPrivateLabelsPerAutomaton,
			") should be strictly smaller than MaxNumLabelsPerAutomaton (",
			config.MaxNumLabelsPerAutomaton,
			"), automatically set to ",
			config.MaxNumLabelsPerAutomaton-1,
		)
		config.MinNumPrivateLabelsPerAutomaton = config.MaxNumLabelsPerAutomaton - 1
	}

	// max number of private labels greater than min number of private labels
	if config.MaxNumPrivateLabelsPerAutomaton < config.MinNumPrivateLabelsPerAutomaton {
		log.Print(
			"Warning: MaxNumPrivateLabelsPerAutomaton (",
			config.MaxNumPrivateLabelsPerAutomaton,
			") should be at least equal to MinNumPrivateLabelsPerAutomaton (",
			config.MinNumPrivateLabelsPerAutomaton,
			"), automatically set to ",
			config.MinNumPrivateLabelsPerAutomaton,
		)
		config.MaxNumPrivateLabelsPerAutomaton = config.MinNumPrivateLabelsPerAutomaton
	}

	// max number of private labels smaller than max number of labels - 1
	if config.MaxNumPrivateLabelsPerAutomaton > config.MaxNumLabelsPerAutomaton-1 {
		log.Print(
			"Warning: MaxNumPrivateLabelsPerAutomaton (",
			config.MaxNumPrivateLabelsPerAutomaton,
			") should be strictly smaller than MaxNumLabelsPerAutomaton (",
			config.MaxNumLabelsPerAutomaton,
			"), automatically set to ",
			config.MaxNumLabelsPerAutomaton-1,
		)
		config.MaxNumPrivateLabelsPerAutomaton = config.MaxNumLabelsPerAutomaton - 1
	}

	// at least zero transition per state
	if config.MinNumTransitionsPerState < 0 {
		log.Print(
			"Warning, MinNumTransitionsPerState (",
			config.MinNumTransitionsPerState,
			") should not be negative, automatically set to 0",
		)
		config.MinNumTransitionsPerState = 0
	}

	// no more transitions per state than labels
	if config.MinNumTransitionsPerState > config.MaxNumLabelsPerAutomaton {
		log.Print(
			"Warning: MinNumTransitionsPerState (",
			config.MinNumTransitionsPerState,
			") should be smaller or equal than MaxNumLabelsPerAutomaton (",
			config.MaxNumLabelsPerAutomaton,
			"), automatically set to ",
			config.MaxNumLabelsPerAutomaton,
		)
		config.MinNumTransitionsPerState = config.MaxNumLabelsPerAutomaton
	}

	// at least one transition per automaton
	if config.MinNumTransitionsPerAutomaton < 1 {
		log.Print(
			"Warning, MinNumTransitionsPerAutomaton (",
			config.MinNumTransitionsPerAutomaton,
			") should be at least 1, automatically set to 1",
		)
		config.MinNumTransitionsPerAutomaton = 1
	}

	// no more transitions per automaton than labels * states
	if config.MinNumTransitionsPerAutomaton > config.MaxNumLabelsPerAutomaton*config.MaxNumStatesPerAutomaton {
		log.Print(
			"Warning: MinNumTransitionsPerAutomaton (",
			config.MinNumTransitionsPerAutomaton,
			") should be smaller or equal than MaxNumLabelsPerAutomaton * MaxNumStatesPerAutomaton (",
			config.MaxNumLabelsPerAutomaton, " * ", config.MaxNumStatesPerAutomaton,
			"), automatically set to ",
			config.MaxNumLabelsPerAutomaton*config.MaxNumStatesPerAutomaton,
		)
		config.MinNumTransitionsPerAutomaton = config.MaxNumLabelsPerAutomaton * config.MaxNumStatesPerAutomaton
	}

	// at least (transitions per state * states) transitions in an automaton
	if config.MinNumTransitionsPerAutomaton < config.MinNumTransitionsPerState*config.MinNumStatesPerAutomaton {
		log.Print(
			"Warning: MinNumTransitionsPerAutomaton (",
			config.MinNumTransitionsPerAutomaton,
			") should be greater or equal than MinNumTransitionsPerState * MinNumStatesPerAutomaton (",
			config.MinNumTransitionsPerState, " * ", config.MinNumStatesPerAutomaton,
			"), automatically set to ",
			config.MinNumTransitionsPerState*config.MinNumStatesPerAutomaton,
		)
		config.MinNumTransitionsPerAutomaton = config.MinNumTransitionsPerState * config.MinNumStatesPerAutomaton
	}

	// at least one automaton
	if config.NumAutomata < 1 {
		log.Print(
			"Warning, NumAutomata (",
			config.NumAutomata,
			") should be at least 1, automatically set to 1",
		)
		config.NumAutomata = 1
	}
}
