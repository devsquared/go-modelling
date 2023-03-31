package godel

// StateIdentifier defines a simple identifier for a state.
type StateIdentifier string

// States defines the states that are registered.
type States map[StateIdentifier]State

// State defines a node in a state machine. A state has some amount of actions that can happen to it that could cause
// it to change state.
type State struct {
	Identifier StateIdentifier `json:"name"`
	Desc       string          `json:"desc"`
	Content    any             `json:"content"` // this could be anything within the state
	Events     Events          `json:"events"`
}

// ActedUponBy defines the transformation that happens when the state is acted on by an action.
func (s *State) ActedUponBy(event EventIdentifier) (StateIdentifier, error) {
	if stateID, ok := s.Events[event]; ok {
		return stateID, nil
	}

	var NilID StateIdentifier
	return NilID, ErrEventNotDefined
}
