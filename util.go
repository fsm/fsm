package fsm

import (
	"regexp"
	"strings"
)

var (
	cleanInputRegex  = regexp.MustCompile("[^a-z0-9 ]+")
	doubleSpaceRegex = regexp.MustCompile(" +")
)

// CleanInput converts the input string to only the following:
// - Lowercase Letters (a-z)
// - Numbers (0-9)
// - Spaces ( )
//
// Uppercase letters are converted to lowercase letters, but any character outside
// of what is noted above is stripped from the string. Double (or more) spaces are
// converted into a single space.
func CleanInput(input string) string {
	return doubleSpaceRegex.ReplaceAllString(cleanInputRegex.ReplaceAllString(strings.ToLower(input), ""), " ")
}

// GetStateMap converts a fsm.StateMachine into a fsm.StateMap
func GetStateMap(stateMachine StateMachine) StateMap {
	stateMap := make(StateMap, 0)
	for _, buildState := range stateMachine {
		stateMap[buildState(nil, nil).Slug] = buildState
	}
	return stateMap
}

// Step performs a single step through a StateMachine
func Step(platform, uuid string, input interface{}, inputToIntentTransformer InputToIntentTransformer, store Store, emitter Emitter, stateMap StateMap) {
	// Get Traverser
	newTraverser := false
	traverser, err := store.FetchTraverser(uuid)
	if err != nil {
		traverser, _ = store.CreateTraverser(uuid)
		traverser.SetCurrentState(StartState)
		traverser.SetPlatform(platform)
		newTraverser = true
	}

	// Get current state
	currentState := stateMap[traverser.CurrentState()](emitter, traverser)
	if newTraverser {
		performEntryAction(currentState, emitter, traverser, stateMap)
	}

	// Transition
	intent, params := inputToIntentTransformer(input, currentState.ValidIntents())
	if intent != nil {
		newState := currentState.Transition(intent, params)
		if newState != nil {
			traverser.SetCurrentState(newState.Slug)
			performEntryAction(newState, emitter, traverser, stateMap)
		} else {
			currentState.Entry(true)
		}
	} else {
		currentState.Entry(true)
	}
}

func performEntryAction(state *State, emitter Emitter, traverser Traverser, stateMap StateMap) error {
	err := state.Entry(false)
	if err != nil {
		return err
	}

	// If we switch states in EntryAction, we want to perform
	// the next states EntryAction
	currentState := traverser.CurrentState()
	if currentState != state.Slug {
		shiftedState := stateMap[currentState](emitter, traverser)
		performEntryAction(shiftedState, emitter, traverser, stateMap)
	}
	return nil
}
