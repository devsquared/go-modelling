package main

import (
	"fmt"
	"github.com/devsquared/godel"
)

type LightPower string

const LightOn LightPower = "on"
const LightOff LightPower = "off"

// State is a basic example of a struct implementing the MachineState interface.
type State struct {
	name        string
	identifier  godel.StateIdentifier
	desc        string
	knownEvents godel.Events
	content     LightPower // simple On or Off
}

// OnEvent deals with the event that comes into the machine and affects the state. As events are known, add a handler here.
func (s *State) OnEvent(event godel.Event) (godel.StateIdentifier, error) {
	fmt.Println("received " + event.Identifier + " event")

	// check knownEvents for the event needed
	if resultState, exists := s.knownEvents[event.Identifier]; exists {
		return resultState, nil
	} else {
		return "", godel.ErrEventNotDefined
	}
}

func (s *State) Identify() string {
	return string(s.identifier)
}

func main() {
	switchFlipEvent := godel.Event{
		Identifier: "flipSwitch",
	}

	lightOnState := State{
		name:       "light is on state",
		identifier: "lightOn",
		knownEvents: map[godel.EventIdentifier]godel.StateIdentifier{
			switchFlipEvent.Identifier: "lightOff",
		},
		content: LightOn,
	}

	lightOffState := State{
		name:       "light is off state",
		identifier: "lightOff",
		knownEvents: map[godel.EventIdentifier]godel.StateIdentifier{
			switchFlipEvent.Identifier: "lightOn",
		},
		content: LightOff,
	}

	exampleLightSwitchMachine := godel.StateMachine{
		Name: "Simple Light Switch State Machine",
		States: map[godel.StateIdentifier]godel.MachineState{
			"lightOn":  &lightOnState,
			"lightOff": &lightOffState,
		},
		CurrentState: &lightOffState,
	}

	fmt.Println("We are starting in the " + exampleLightSwitchMachine.CurrentState.Identify())
	fmt.Println("Sending the light switch flip event.")
	if err := exampleLightSwitchMachine.SendEvent(switchFlipEvent); err != nil {
		panic("welp this is awkward: " + err.Error())
	}

	fmt.Println("Have ended up in this state: " + exampleLightSwitchMachine.CurrentState.Identify())
}
