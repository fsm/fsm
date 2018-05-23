package fsm

// GetStateMap converts a StateMachine into a StateMap
func GetStateMap(stateMachine StateMachine) StateMap {
	stateMap := make(StateMap, 0)
	for _, buildState := range stateMachine {
		stateMap[buildState(nil, nil).Slug] = buildState
	}
	return stateMap
}

// Step performs a single step through a StateMachine.
//
// This function handles the nuance of the logic for a single step through a state machine.
// ALL fsm-target's should call Step directly, and not attempt to handle the process of stepping through
// the StateMachine, so all platforms function with the same logic.
func Step(platform, uuid string, input interface{}, InputTransformer InputTransformer, store Store, emitter Emitter, stateMap StateMap) {
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
	intent, params := InputTransformer(input, currentState.ValidIntents())
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

// performEntryAction handles the logic of switching states and calling the Entry function.
//
// It is handled via this function, as a state can manually switch states in the Entry function.
// If that occurs, we then perform the Entry function of that state.  This continues until we land
// in a state whose Entry action doesn't shift us to a new state.
func performEntryAction(state *State, emitter Emitter, traverser Traverser, stateMap StateMap) error {
	err := state.Entry(false)
	if err != nil {
		return err
	}

	// If we switch states in Entry action, we want to perform
	// the next states Entry action.
	currentState := traverser.CurrentState()
	if currentState != state.Slug {
		shiftedState := stateMap[currentState](emitter, traverser)
		performEntryAction(shiftedState, emitter, traverser, stateMap)
	}
	return nil
}
