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
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(int64(time.Now().Nanosecond()))

	var configFileName string
	var outputFileName string
	flag.StringVar(&configFileName, "conf", configFile, "Path to configuration file")
	flag.StringVar(&outputFileName, "out", outputFile, "Path to output file")
	flag.Parse()

	readConfigurationFile(configFileName)

	g := genGraph()
	out, err := json.Marshal(g.jsonAutomata)
	if err != nil {
		log.Fatal("Error: cannot build the json output")
		//log.Panic(err)
	}

	log.Print("Writing automata into ", outputFileName)
	err = ioutil.WriteFile(outputFileName, out, 0644)
	if err != nil {
		log.Fatal("Error: cannot write to output file (", outputFileName, ")")
		log.Panic(err)
	}
}
