package godel

import "fmt"

// TODO: Need to create a good builder pattern here so that creating machine with well-defined states and events is easy and intuitive

// StateMachine defines a simple state machine instance.
type StateMachine struct {
	Name         string       `json:"Name"`
	Desc         string       `json:"Desc"`
	States       States       `json:"States"`
	CurrentState MachineState `json:"-"`
}

func (m *StateMachine) SendEvent(event Event) error {
	newStateID, err := m.CurrentState.OnEvent(event)
	if err != nil {
		return fmt.Errorf("failed responding to event: %w", err)
	}

	if newState, exists := m.States[newStateID]; exists {
		m.CurrentState = newState
		return nil
	} else {
		return ErrStateNotDefined
	}
}

// Need a way to execute many, many instances of a machine to test the specifications.
