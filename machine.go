package godel

import "fmt"

// Need a way to execute many, many instances of a machine to test the specifications.

// StateMachine defines a simple state machine instance.
type StateMachine struct {
	Name         string `json:"Name"`
	Desc         string `json:"Desc"`
	States       States `json:"States"`
	CurrentState State  `json:"-"`
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

type StateMachineBuilder struct {
	machine *StateMachine
}

func NewStateMachineBuilder(name string, currentState State) *StateMachineBuilder {
	return &StateMachineBuilder{machine: &StateMachine{Name: name, States: make(States, 0), CurrentState: currentState}}
}

func (b *StateMachineBuilder) AddState(id StateIdentifier, state State) *StateMachineBuilder {
	b.machine.States[id] = state // simple set. it is clear if you add multiple with same id

	return b
}

func (b *StateMachineBuilder) WithDesc(desc string) *StateMachineBuilder {
	b.machine.Desc = desc

	return b
}

func (b *StateMachineBuilder) BuildStateMachine() *StateMachine {
	return b.machine
}
