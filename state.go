package godel

// StateIdentifier defines a simple identifier for a state.
type StateIdentifier string

// States defines the states that are registered.
type States map[StateIdentifier]MachineState

// MachineState is the definition for a state in a machine. A state needs to react to events and that is defined in OnEvent.
type MachineState interface {
	OnEvent(event Event) (StateIdentifier, error)
	Identify() string
}
