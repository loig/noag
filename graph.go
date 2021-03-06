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
	"log"
	"math/rand"
)

type graph struct {
	automata     []automaton
	jsonAutomata []JSONAutomaton
}

func genGraph() graph {
	log.Print("Starting generation of ", config.NumAutomata, " automata")

	var g graph
	g.automata = make([]automaton, config.NumAutomata)
	g.jsonAutomata = make([]JSONAutomaton, config.NumAutomata)

	lastLabel := 0
	allLabels := make([]int, 0)

	for i := 0; i < config.NumAutomata; i++ {
		// determine numbers of labels and private labels
		numLabels := rand.Intn(config.MaxNumLabelsPerAutomaton-config.MinNumLabelsPerAutomaton+1) + config.MinNumLabelsPerAutomaton
		maxPrivateLabels := config.MaxNumPrivateLabelsPerAutomaton
		if maxPrivateLabels > numLabels-1 {
			maxPrivateLabels = numLabels - 1
		}
		minPrivateLabels := config.MinNumPrivateLabelsPerAutomaton
		if minPrivateLabels > maxPrivateLabels {
			minPrivateLabels = maxPrivateLabels
		}
		numPrivateLabels := rand.Intn(maxPrivateLabels-minPrivateLabels+1) + minPrivateLabels
		// build a set of labels
		labels := make([]int, numLabels)
		if lastLabel == 0 {
			for i := 0; i < numLabels; i++ {
				labels[i] = i
				if i < numLabels-numPrivateLabels {
					allLabels = append(allLabels, i)
				}
			}
			lastLabel += numLabels
		} else {
			numSharedLabelsFromPrevious := rand.Intn(numLabels-numPrivateLabels) + 1
			if numSharedLabelsFromPrevious > len(allLabels) {
				numSharedLabelsFromPrevious = len(allLabels)
			}
			if numSharedLabelsFromPrevious > 0 {
				rand.Shuffle(len(allLabels), func(i, j int) {
					allLabels[i], allLabels[j] = allLabels[j], allLabels[i]
				})
				for i := 0; i < numSharedLabelsFromPrevious; i++ {
					labels[i] = allLabels[i]
				}
			}
			for i := numSharedLabelsFromPrevious; i < numLabels; i++ {
				lastLabel++
				labels[i] = lastLabel
				if i < numLabels-numPrivateLabels {
					allLabels = append(allLabels, lastLabel)
				}
			}
		}
		// generate an automaton
		log.Print("Starting generation of automaton ", automatonName, i)
		g.automata[i] = genAutomaton(labels)
		g.jsonAutomata[i] = g.automata[i].toJSON(i)
		log.Print("Labels: ", g.jsonAutomata[i].InputSymbols)
		log.Print("Automaton ", automatonName, i, " generated")
	}

	log.Print("Generation complete")
	return g
}
