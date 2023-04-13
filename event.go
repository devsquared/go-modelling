package godel

// EventIdentifier is a simple identifier for an event.
type EventIdentifier string

// Events defines a registry for events to the identifier for the resulting state.
type Events map[EventIdentifier]StateIdentifier

// Event defines something that happens to a state. For the sake of simplicity here, a State knows and defines the actions that
// can happen to it. This is because the context of the original State is important for the outcome; not simply the action itself.
type Event struct {
	Identifier EventIdentifier `json:"Identifier"`
	Desc       string          `json:"Desc"`
}
