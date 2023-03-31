package main

import (
	"encoding/json"
	"fmt"

	"github.com/devsquared/godel"
)

func main() {
	fmt.Println("Welcome to godel!")

	fmt.Println("Let's build out a simple example here.")

	// Model an inventory entry
	type InventoryEntry struct {
		Status   string
		Quantity int
	}

	manualRemovalAction := godel.Action{
		Name:        "Manual Inventory Removal",
		Desc:        "Action to remove an entry from the inventory - possible on all states",
		ResultState: godel.State{},
	}

	receivedStatusState := godel.State{
		Name: "ASN received and counted",
		Desc: "Inbound shipment received at our loading dock and has been counted and put away for waving.",
		Content: InventoryEntry{
			Status:   "received",
			Quantity: 10,
		},
		Actions: []godel.Action{manualRemovalAction},
	}

	edi944Action := godel.Action{
		Name:        "EDI 944 received",
		Desc:        "We have received the EDI 944.",
		ResultState: receivedStatusState,
	}

	expectedStatusState := godel.State{
		Name: "ASN Expected",
		Desc: "We are expecting an ASN and the quantity shows us this in expected. This is typically the initial state.",
		Content: InventoryEntry{
			Status:   "expected",
			Quantity: 10,
		},
		Actions: []godel.Action{edi944Action, manualRemovalAction},
	}

	exampleMachine := godel.StateMachine{
		Name:         "Simple receiving ASN State Machine",
		Desc:         "In this simple example, we receive an inbound shipment to see how that affects the expected quantity.",
		States:       []godel.State{expectedStatusState, receivedStatusState},
		CurrentState: expectedStatusState, // initial state is here
	}

	fmt.Println("We have an example State Machine called " + exampleMachine.Name)

	fmt.Println("Let's simulate receiving a 944.")
	fmt.Println("We start in the initial state of " + exampleMachine.CurrentState.Name)
	fmt.Println(fmt.Sprintf("In this state, we have the content of: %v", exampleMachine.CurrentState.Content))
	fmt.Println("Upon the action of " + edi944Action.Name + ", the machine moves to the following state: ")
	err := exampleMachine.ReceivedAction(edi944Action)
	if err != nil {
		fmt.Println(" unfortunately an INVALID STATE")
		panic(err.Error())
	}
	fmt.Println(exampleMachine.CurrentState.Name)
	fmt.Println(fmt.Sprintf("In this state, we have the content of: %v", exampleMachine.CurrentState.Content))

	fmt.Println("Now let's observe godel's ability to marshal our state machine to meaningful JSON: ")
	fmt.Println("------------------------")
	fmt.Println("Marshalling the example to output")

	data, err := json.MarshalIndent(exampleMachine, "", " ")
	if err != nil {
		panic(fmt.Errorf("unable to marshal the machine: %w", err).Error())
	}
	
	fmt.Println("For the simple example, we will output the data here rather than storing in file.")
	fmt.Println(data)
}
