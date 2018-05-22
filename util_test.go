package fsm_test

import (
	"testing"

	"github.com/fsm/fsm"
)

var (
	getStartState = func(fsm.Emitter, fsm.Traverser) *fsm.State {
		return &fsm.State{
			Slug: fsm.StartState,
			Entry: func(isReentry bool) error {
				return nil
			},
			ValidIntents: func() []*fsm.Intent {
				return nil
			},
			Transition: func(*fsm.Intent, map[string]string) *fsm.State {
				return nil
			},
		}
	}

	getStateA = func(fsm.Emitter, fsm.Traverser) *fsm.State {
		return &fsm.State{
			Slug: "A",
			Entry: func(isReentry bool) error {
				return nil
			},
			ValidIntents: func() []*fsm.Intent {
				return nil
			},
			Transition: func(*fsm.Intent, map[string]string) *fsm.State {
				return nil
			},
		}
	}

	getStateB = func(fsm.Emitter, fsm.Traverser) *fsm.State {
		return &fsm.State{
			Slug: "B",
			Entry: func(isReentry bool) error {
				return nil
			},
			ValidIntents: func() []*fsm.Intent {
				return nil
			},
			Transition: func(*fsm.Intent, map[string]string) *fsm.State {
				return nil
			},
		}
	}

	simpleMachine = []fsm.BuildState{
		getStartState,
		getStateA,
		getStateB,
	}
)

func TestStateMap(t *testing.T) {
	stateMap := fsm.GetStateMap(simpleMachine)

	// Test state map was generated properly
	if stateMap["A"](nil, nil).Slug != "A" {
		t.Fail()
	}
	if stateMap["B"](nil, nil).Slug != "B" {
		t.Fail()
	}
	if len(stateMap) != 3 {
		t.Fail()
	}
}
