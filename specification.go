package godel

// TODO: lots of work needed here

// Specification is a definition on what is expected out of a StateMachine. Specifications observe events and the
// results caused by that action in the machine.
type Specification struct {
	Name      string
	BelongsTo StateMachine
	Observes  EventIdentifier
}
