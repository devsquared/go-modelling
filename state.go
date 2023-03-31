package godel

// State defines a node in a state machine. A state has some amount of actions that can happen to it that could cause
// it to change state.
type State struct {
	Name    string   `json:"name"`
	Desc    string   `json:"desc"`
	Content any      `json:"content"` // this could be anything within the state
	Actions []Action `json:"actions"`
}

// ActedUponBy defines the transformation that happens when the state is acted on by an action.
func (s *State) ActedUponBy(actionName string) (State, error) {
	for _, a := range s.Actions {
		if a.Name == actionName {
			return a.ResultState, nil
		}
	}

	return State{}, ErrActionNotDefined
}
