package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/devsquared/godel"
	"github.com/devsquared/godel/model"
)

func main() {
	fmt.Println("Welcome to godel!")

	fmt.Println("Let's build out a simple example here.")

	// Model an inventory entry
	type InventoryEntry struct {
		Status   string
		Quantity int
	}

	manualRemovalEvent := godel.Event{
		Identifier:  "Manual Inventory Removal",
		Desc:        "Event to remove an entry from the inventory - possible on all states",
		ResultState: "Removed",
	}

	receivedStatusKnownEvents := map[godel.EventIdentifier]godel.StateIdentifier{
		manualRemovalEvent.Identifier: "Removed",
	}
	receivedStatusState := godel.State{
		Identifier: "ASN received and counted",
		Desc:       "Inbound shipment received at our loading dock and has been counted and put away for waving.",
		Content: InventoryEntry{
			Status:   "received",
			Quantity: 10,
		},
		Events: receivedStatusKnownEvents,
	}

	edi944Action := godel.Event{
		Identifier:  "EDI 944 received",
		Desc:        "We have received the EDI 944.",
		ResultState: receivedStatusState.Identifier,
	}

	expectedStatusStateKnownEvents := map[godel.EventIdentifier]godel.StateIdentifier{
		edi944Action.Identifier:       receivedStatusState.Identifier,
		manualRemovalEvent.Identifier: "Removed",
	}
	expectedStatusState := godel.State{
		Identifier: "ASN Expected",
		Desc:       "We are expecting an ASN and the quantity shows us this in expected. This is typically the initial state.",
		Content: InventoryEntry{
			Status:   "expected",
			Quantity: 10,
		},
		Events: expectedStatusStateKnownEvents,
	}

	machineStates := map[godel.StateIdentifier]godel.State{
		expectedStatusState.Identifier: expectedStatusState,
		receivedStatusState.Identifier: receivedStatusState,
	}
	exampleMachine := godel.StateMachine{
		Name:         "Simple receiving ASN State Machine",
		Desc:         "In this simple example, we receive an inbound shipment to see how that affects the expected quantity.",
		States:       machineStates,
		CurrentState: expectedStatusState, // initial state is here
	}

	fmt.Println("We have an example State Machine called " + exampleMachine.Name)

	fmt.Println("Let's simulate receiving a 944.")
	fmt.Println("We start in the initial state of " + exampleMachine.CurrentState.Identifier)
	fmt.Println(fmt.Sprintf("In this state, we have the content of: %v", exampleMachine.CurrentState.Content))
	fmt.Println("Upon the action of " + edi944Action.Identifier + ", the machine moves to the following state: ")
	newStateID, err := exampleMachine.ReceivedEvent(edi944Action.Identifier)
	if err != nil {
		fmt.Println(" unfortunately an INVALID STATE")
		panic(err.Error())
	}
	fmt.Println(newStateID)
	fmt.Println(fmt.Sprintf("In this state, we have the content of: %v", exampleMachine.CurrentState.Content))

	fmt.Println("Now let's observe godel's ability to marshal our state machine to meaningful JSON: ")
	fmt.Println("------------------------")
	fmt.Println("Marshalling the example to output")

	data, err := json.MarshalIndent(exampleMachine, "", " ")
	if err != nil {
		panic(fmt.Errorf("unable to marshal the machine: %w", err).Error())
	}

	fmt.Println("For the simple example, we will output the data here rather than storing in file.")
	fmt.Println(string(data))

	// Right Arrow
	if err = model.DrawArrow(os.Stdout, model.Right, 30, string(exampleMachine.CurrentState.Identifier)); err != nil {
		panic("drawing sucks")
	}

	fmt.Println()

	// Left Arrow
	if err = model.DrawArrow(os.Stdout, model.Left, 30, string(exampleMachine.CurrentState.Identifier)); err != nil {
		panic("drawing sucks")
	}

	fmt.Println()

	// Up Arrow
	if err = model.DrawArrow(os.Stdout, model.Up, 20, string(exampleMachine.CurrentState.Identifier)); err != nil {
		panic("drawing sucks")
	}

	fmt.Println()

	// Down Arrow
	if err = model.DrawArrow(os.Stdout, model.Down, 20, string(exampleMachine.CurrentState.Identifier)); err != nil {
		panic("drawing sucks")
	}

	fmt.Println()

	// State Node
	if err = model.DrawStateNode(os.Stdout, string(exampleMachine.CurrentState.Identifier), "Status: received"); err != nil {
		panic("drawing sucks")
	}
}
