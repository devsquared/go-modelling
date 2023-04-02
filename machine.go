package godel

import "fmt"

// TODO: Need to create a good builder pattern here so that creating machine with well-defined states and events is easy and intuitive

// StateMachine defines a simple state machine.
type StateMachine struct {
	Name         string `json:"Name"`
	Desc         string `json:"Desc"`
	States       States `json:"States"`
	CurrentState State  `json:"-"`
}

// ReceivedEvent kicks off the state machines processing of the event given.
func (m *StateMachine) ReceivedEvent(eventID EventIdentifier) (StateIdentifier, error) {
	var NilStateID StateIdentifier
	if eventID == "" {
		return NilStateID, ErrEventNotDefined
	}

	resultStateID, err := m.CurrentState.ActedUponBy(eventID)
	if err != nil {
		return NilStateID, fmt.Errorf("error processing event: %w", err)
	}

	if resultState, ok := m.States[resultStateID]; ok {
		m.CurrentState = resultState
		return m.CurrentState.Identifier, nil
	}

	return NilStateID, ErrStateNotDefined
}
