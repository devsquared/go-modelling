package godel

import "fmt"

// TODO: Let's create some custom marshalling so that we can print how we want here.

// StateMachine defines a simple state machine.
type StateMachine struct {
	Name         string `json:"Name"`
	Desc         string `json:"Desc"`
	States       States `json:"States"`
	CurrentState State  `json:"-"`
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
