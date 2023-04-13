package main

import (
	"encoding/json"
	"fmt"
	"github.com/devsquared/godel"
	"os"
)

type LightPower string

const LightOn LightPower = "on"
const LightOff LightPower = "off"

// LightOnState is a basic example of a struct implementing the State interface.
type LightOnState struct {
	Name        string                `json:"Name"`
	Identifier  godel.StateIdentifier `json:"-"`
	Desc        string                `json:"Desc"`
	KnownEvents godel.Events          `json:"KnownEvents"`
	Content     LightPower            `json:"Content"`
}

// LightOffState is another basic example.
type LightOffState struct {
	Name        string                `json:"Name"`
	Identifier  godel.StateIdentifier `json:"-"`
	Desc        string                `json:"Desc"`
	KnownEvents godel.Events          `json:"KnownEvents"`
	Content     LightPower            `json:"Content"`
}

// OnEvent deals with the event that comes into the machine and affects the state. As events are known, add a handler here.
func (s *LightOnState) OnEvent(event godel.Event) (godel.StateIdentifier, error) {
	fmt.Println("received " + event.Identifier + " event")

	// check knownEvents for the event needed
	if resultState, exists := s.KnownEvents[event.Identifier]; exists {
		return resultState, nil
	} else {
		return "", godel.ErrEventNotDefined
	}
}

func (s *LightOnState) Identify() godel.StateIdentifier {
	return "lightOn"
}

func (s *LightOffState) OnEvent(event godel.Event) (godel.StateIdentifier, error) {
	fmt.Println("received " + event.Identifier + " event")

	// check knownEvents for the event needed
	if resultState, exists := s.KnownEvents[event.Identifier]; exists {
		return resultState, nil
	} else {
		return "", godel.ErrEventNotDefined
	}
}

func (s *LightOffState) Identify() godel.StateIdentifier {
	return "lightOff"
}

func main() {
	switchFlipEvent := godel.Event{
		Identifier: "flipSwitch",
	}

	lightOnState := LightOnState{
		Name:       "light is on state",
		Identifier: "lightOn",
		KnownEvents: map[godel.EventIdentifier]godel.StateIdentifier{
			switchFlipEvent.Identifier: "lightOff",
		},
		Content: LightOn,
	}

	lightOffState := LightOffState{
		Name:       "light is off state",
		Identifier: "lightOff",
		KnownEvents: map[godel.EventIdentifier]godel.StateIdentifier{
			switchFlipEvent.Identifier: "lightOn",
		},
		Content: LightOff,
	}

	exampleLightSwitchMachine := godel.NewStateMachineBuilder("Simple Light Switch State Machine", &lightOffState).
		WithDesc("A simple state machine simulating a simple light switch!").
		AddState("lightOn", &lightOnState).
		AddState("lightOff", &lightOffState).
		BuildStateMachine()

	fmt.Println("We are starting in the " + exampleLightSwitchMachine.CurrentState.Identify())
	fmt.Println("Sending the light switch flip event.")
	if err := exampleLightSwitchMachine.SendEvent(switchFlipEvent); err != nil {
		panic("welp this is awkward: " + err.Error())
	}

	fmt.Println("Have ended up in this state: " + exampleLightSwitchMachine.CurrentState.Identify())

	fmt.Println("Let's marshal to JSON to conveniently see the machine.")

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(&exampleLightSwitchMachine)
}
