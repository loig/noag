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

	var g graph
	g.automata = make([]automaton, graphNumAutomata)
	g.jsonAutomata = make([]JSONAutomaton, graphNumAutomata)

	lastLabel := 0
	allLabels := make([]int, 0)

	for i := 0; i < graphNumAutomata; i++ {
		log.Print("A", i, ":")
		// determine numbers of labels and private labels
		numLabels := rand.Intn(automatonMaxLabels-automatonMinLabels+1) + automatonMinLabels
		maxPrivateLabels := automatonMaxPrivateLabels
		if maxPrivateLabels > numLabels-1 {
			maxPrivateLabels = numLabels - 1
		}
		minPrivateLabels := automatonMinPrivateLabels
		if minPrivateLabels > maxPrivateLabels {
			minPrivateLabels = maxPrivateLabels
		}
		numPrivateLabels := rand.Intn(maxPrivateLabels-minPrivateLabels+1) + minPrivateLabels
		log.Print(numLabels, " labels, ", numPrivateLabels, " privates")
		// build a set of labels
		labels := make([]int, numLabels)
		if lastLabel == 0 {
			for i := 0; i < numLabels; i++ {
				labels[i] = i
				if i < numLabels-numPrivateLabels {
					log.Print("Maybe shared: a", i)
					allLabels = append(allLabels, i)
				} else {
					log.Print("Private: a", i)
				}
			}
			lastLabel += numLabels
		} else {
			numSharedLabelsFromPrevious := rand.Intn(numLabels-numPrivateLabels) + 1
			if numSharedLabelsFromPrevious > len(allLabels) {
				numSharedLabelsFromPrevious = len(allLabels)
			}
			log.Print(numSharedLabelsFromPrevious, " from previous")
			if numSharedLabelsFromPrevious > 0 {
				rand.Shuffle(len(allLabels), func(i, j int) {
					allLabels[i], allLabels[j] = allLabels[j], allLabels[i]
				})
				for i := 0; i < numSharedLabelsFromPrevious; i++ {
					labels[i] = allLabels[i]
					log.Print("Shared: a", allLabels[i])
				}
			}
			for i := numSharedLabelsFromPrevious; i < numLabels; i++ {
				lastLabel++
				labels[i] = lastLabel
				if i < numLabels-numPrivateLabels {
					log.Print("Maybe shared: a", lastLabel)
					allLabels = append(allLabels, lastLabel)
				} else {
					log.Print("Private: a", lastLabel)
				}
			}
		}
		// generate an automaton
		g.automata[i] = genAutomaton(labels)
		g.jsonAutomata[i] = g.automata[i].toJSON(i)
	}

	return g

}
