package godel

import "fmt"

// StateMachine defines a simple state machine.
type StateMachine struct {
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	States       States `json:"-"`
	CurrentState State  `json:"currentState"`
}

// ReceivedEvent kicks off the state machines processing of the event given.
func (m *StateMachine) ReceivedEvent(event Event) error {
	if event.Identifier == "" {
		return ErrEventNotDefined
	}

	resultStateID, err := m.CurrentState.ActedUponBy(event.Identifier)
	if err != nil {
		return fmt.Errorf("error processing event: %w", err)
	}

	if resultState, ok := m.States[resultStateID]; ok {
		m.CurrentState = resultState
		return nil
	}

	return ErrStateNotDefined
}
