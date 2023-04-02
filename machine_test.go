package godel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateMachine_ReceivedEvent(t *testing.T) {
	type testScenario struct {
		name            string
		machine         StateMachine
		eventID         EventIdentifier
		expectedStateID StateIdentifier
		expectedErr     error
	}

	testScenarios := []testScenario{
		{
			name:        "empty identifier given - event not defined",
			machine:     StateMachine{}, // machine does not matter in this case
			eventID:     "",
			expectedErr: ErrEventNotDefined,
		},
		{
			name: "state defined but event not found on current state",
			machine: StateMachine{
				CurrentState: State{
					Events: map[EventIdentifier]StateIdentifier{}, //empty - so event not defined
				},
			},
			eventID:     "something",
			expectedErr: fmt.Errorf("error processing event: %w", ErrEventNotDefined),
		},
		{
			name: "resulting state from event resulting in state not defined",
			machine: StateMachine{
				CurrentState: State{
					Events: Events{
						"something": "someUndefinedState",
					},
				},
			},
			eventID:     "something",
			expectedErr: ErrStateNotDefined,
		},
		{
			name: "machine well defined with events and states",
			machine: StateMachine{
				States: States{
					"someState": State{
						Identifier: "someState",
					},
				},
				CurrentState: State{
					Events: Events{
						"someID": "someState",
					},
				},
			},
			eventID:         "someID",
			expectedStateID: "someState",
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			actualStateID, actualErr := ts.machine.ReceivedEvent(ts.eventID)
			assert.Equal(t, actualErr, ts.expectedErr)
			assert.Equal(t, ts.expectedStateID, actualStateID)
		})
	}
}
