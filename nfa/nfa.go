package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

// You may define helper functions here.
func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO: Write the Reachable function,
	// return true if the nfa accepts the input and can reach the final state with that input,
	// return false otherwise - path by path or back tracking (same logic oppoite direction) needed

	// Boundary condition - when no more symbols to transit, possibly empty list of symbols, it's the last state where it could get
	if len(input) == 0 {
		if start == final {
			return true
		} else {
			return false
		}
	} else {
		// assign the next state 
		// (st state, sym rune) []state
		next := transitions(start, input[0])
		for i := 0; i < len(next); i++ {
			start = next[i]
			// Recursion call to the next states
			if Reachable(transitions, start, final, input[1:]){
				return true
			}
		}
	}

	// Place holder: panic("TODO: implement this!")
	return false
}

