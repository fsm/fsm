package fsm

// StateMachine is a k:v map that maps the
// slug of States to BuildState functions
type StateMachine map[string]BuildState

// BuildState is a function that generates a State
// with access to a specific Emitter and Traverser
type BuildState func(Emitter, Traverser) *State

// State represents an individual state in a larger state machine
type State struct {
	Slug          string
	EntryAction   func() error
	ReentryAction func() error
	Transition    func(interface{}) *State
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
	UUID() string
	SetUUID(string)
	CurrentState() string
	SetCurrentState(string)
	Upsert(key string, value interface{}) error
	Fetch(key string) (interface{}, error)
	Delete(key string) error
}
