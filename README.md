# NoAG (Networks of Automata Generation)
A tool for generation of networks of automata

## Configuration
The tool can be configured using a json configuration file, see conf.json for an example.

One can define:
- MinNumStatesPerAutomaton: the minimum number of states in each generated automaton
- MaxNumStatesPerAutomaton: the maximum number of states in each generated automaton
- MinNumGoalStatesPerAutomaton: the minimum number of goal states in each generated automaton,
- MaxNumGoalStatesPerAutomaton: the maximum number of goal states in each generated automaton,
- MinNumLabelsPerAutomaton: the minimum number of different labels used by each generated automaton,
- MaxNumLabelsPerAutomaton: the maximum number of different labels used by each generated automaton,
- MinNumPrivateLabelsPerAutomaton: the minimum number of different private labels used by each generated automaton,
- MaxNumPrivateLabelsPerAutomaton: the maximum number of different private labels used by each generated automaton,
- MinNumTransitionsPerState: the minimum number of transitions going out of each state of each generated automaton,
- MinNumTransitionsPerAutomaton: the minimum number of transitions in each generated automaton,
- NumAutomata: the number of automata to generate

## Important remarks
The automata generated should all be deterministic, non-empty, and their interaction graph should have only one connected component.

MinNumStatesPerAutomaton, MaxNumStatesPerAutomaton, MinNumGoalStatesPerAutomaton, MaxNumGoalStatesPerAutomaton, MinNumLabelsPerAutomaton, MaxNumLabelsPerAutomaton, MinNumPrivateLabelsPerAutomaton, NumAutomata are guaranteed to be respected.

MaxNumPrivateLabelsPerAutomaton, MinNumTransitionsPerState, MinNumTransitionsPerAutomaton can sometimes be impossible to respect (depending on the random values generated from the others parameters for each particular automaton), in these cases they just won't be.

## Usage
In order to generate automata according to the characteristics given in conf.json and store these automata in the file out.json, just use the following command:

./noag -conf conf.json -out out.json
