package fsm_test

import (
	"errors"
	"testing"

	cachestore "github.com/fsm/cache-store"
	"github.com/fsm/fsm"
	"github.com/fsm/test"
)

var (
	buzzIntent = &fsm.Intent{
		Slug: "buzz",
		Utterances: []string{
			"buzz",
		},
	}

	catchAllIntent = &fsm.Intent{
		Slug: "catch-all",
		Slots: map[string]*fsm.Type{
			"input": {
				Slug: "literal",
				IsValid: func(string) bool {
					return true
				},
			},
		},
		Utterances: []string{
			"{input}",
		},
	}

	simpleMachine = []fsm.BuildState{
		getStartState,
		getStateA,
		getStateB,
		getStateC,
		getStateD,
	}
)

func getStartState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: fsm.StartState,
		Entry: func(isReentry bool) error {
			return nil
		},
		ValidIntents: func() []*fsm.Intent {
			return []*fsm.Intent{
				catchAllIntent,
			}
		},
		Transition: func(*fsm.Intent, map[string]string) *fsm.State {
			return getStateA(emitter, traverser)
		},
	}
}

func getStateA(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: "A",
		Entry: func(isReentry bool) error {
			return nil
		},
		ValidIntents: func() []*fsm.Intent {
			return []*fsm.Intent{
				catchAllIntent,
			}
		},
		Transition: func(intent *fsm.Intent, params map[string]string) *fsm.State {
			if params["input"] == "hello" {
				return getStateB(emitter, traverser)
			}
			return nil
		},
	}
}

func getStateB(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: "B",
		Entry: func(isReentry bool) error {
			return nil
		},
		ValidIntents: func() []*fsm.Intent {
			return []*fsm.Intent{
				buzzIntent,
			}
		},
		Transition: func(intent *fsm.Intent, params map[string]string) *fsm.State {
			if intent == buzzIntent {
				return getStateC(emitter, traverser)
			}
			return nil
		},
	}
}

func getStateC(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: "C",
		Entry: func(isReentry bool) error {
			traverser.SetCurrentState(getStateD(emitter, traverser).Slug)
			return nil
		},
		ValidIntents: func() []*fsm.Intent {
			return nil
		},
		Transition: func(intent *fsm.Intent, params map[string]string) *fsm.State {
			return nil
		},
	}
}

func getStateD(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: "D",
		Entry: func(isReentry bool) error {
			return errors.New("END")
		},
		ValidIntents: func() []*fsm.Intent {
			return nil
		},
		Transition: func(intent *fsm.Intent, params map[string]string) *fsm.State {
			return nil
		},
	}
}

func TestStateMap(t *testing.T) {
	stateMap := fsm.GetStateMap(simpleMachine)

	// Test state map was generated properly
	if stateMap["A"](nil, nil).Slug != "A" {
		t.Fail()
	}
	if stateMap["B"](nil, nil).Slug != "B" {
		t.Fail()
	}
	if len(stateMap) != 5 {
		t.Fail()
	}
}

func TestStep(t *testing.T) {
	traverser := test.New(simpleMachine, cachestore.New())

	// Testing that saying anything shifts us to 'A' state
	traverser.Send("Hello")
	if traverser.CurrentState() != "A" {
		t.Fail()
	}

	// Testing that saying 'world' does not shift us to any state and
	// we remain in state A.
	traverser.Send("world")
	if traverser.CurrentState() != "A" {
		t.Fail()
	}

	// Testing that "hello" moves us to 'B' state
	traverser.Send("hello")
	if traverser.CurrentState() != "B" {
		t.Fail()
	}

	// Testing that saying 'hello' does not shift us out of B state
	// (only buzz can move us out)
	traverser.Send("hello")
	if traverser.CurrentState() != "B" {
		t.Fail()
	}

	// Testing that 'buzz' shifts us out of B state to C state,
	// but then immediately redirects to the D state.
	traverser.Send("buzz")
	if traverser.CurrentState() != "D" {
		t.Fail()
	}
}
