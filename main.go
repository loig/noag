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
	"time"
)

func main() {

	rand.Seed(int64(time.Now().Nanosecond()))

	a := genAutomaton([]int{1, 3, 12, 27, 5})
	fmt.Println(a)
	fmt.Println(a.toJSON(3))

}
