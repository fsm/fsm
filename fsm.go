package fsm

// StartState is a constant for defining the slug of
// the start state for all StateMachines.
const StartState = "start"

// StateMachine is an array of all BuildState functions
type StateMachine []BuildState

// StateMap is a k:v map for all BuildState functions
// in a StateMachine.  This is exclusively utilized
// by the internal workings of targets.
type StateMap map[string]BuildState

// BuildState is a function that generates a State
// with access to a specific Emitter and Traverser
type BuildState func(Emitter, Traverser) *State

// State represents an individual state in a larger state machine
type State struct {
	Slug         string
	Entry        func(isReentry bool) error
	ValidIntents func() []*Intent
	Transition   func(*Intent, map[string]string) *State
}

// Emitter is a generic interface to output arbitrary data.
// Emit is generally called from State.EntryAction.
type Emitter interface {
	Emit(interface{}) error
}

// A Store is a generic interface responsible for managing
// The fetching and creation of traversers
type Store interface {
	FetchTraverser(uuid string) (Traverser, error)
	CreateTraverser(uuid string) (Traverser, error)
}

// A Traverser is an individual that is traversing the
// StateMachine.  This interface that is responsible
// for managing the state of that individual
type Traverser interface {
	// UUID
	UUID() string
	SetUUID(string)

	// Platform
	Platform() string
	SetPlatform(string)

	// State
	CurrentState() string
	SetCurrentState(string)

	// Data
	Upsert(key string, value interface{})
	Fetch(key string) interface{}
	Delete(key string)
}
