package godel

import "fmt"

// StateMachine defines a simple state machine.
type StateMachine struct {
	Name         string  `json:"name"`
	Desc         string  `json:"desc"`
	States       []State `json:"states"`
	CurrentState State   `json:"currentState"`
}

// ReceivedAction kicks off the state machines processing of the action given.
func (m *StateMachine) ReceivedAction(action Action) error {
	resultState, err := m.CurrentState.ActedUponBy(action.Name)
	if err != nil {
		return fmt.Errorf("error processing action: %w", err)
	}

	m.CurrentState = resultState
	return nil
}
