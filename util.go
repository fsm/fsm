package fsm

const FSMPlatform = "fsm.platform"

// GetStateMap converts a fsm.StateMachine into a fsm.StateMap
func GetStateMap(stateMachine StateMachine) StateMap {
	stateMap := make(StateMap, 0)
	for _, buildState := range stateMachine {
		stateMap[buildState(nil, nil).Slug] = buildState
	}
	return stateMap
}

// Step performs a single step through a StateMachine
func Step(platform, uuid string, intent *Intent, store Store, emitter Emitter, stateMap StateMap) {
	// Get Traverser
	newTraverser := false
	traverser, err := store.FetchTraverser(uuid)
	if err != nil {
		traverser, _ = store.CreateTraverser(uuid)
		traverser.SetCurrentState("start")
		traverser.Upsert("fsm.platform", platform)
		newTraverser = true
	}

	// Get current state
	currentState := stateMap[traverser.CurrentState()](emitter, traverser)
	if newTraverser {
		performEntryAction(currentState, emitter, traverser, stateMap)
	}

	// Transition
	newState := currentState.Transition(intent)
	if newState != nil {
		traverser.SetCurrentState(newState.Slug)
		performEntryAction(newState, emitter, traverser, stateMap)
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
