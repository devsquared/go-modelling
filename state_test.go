package godel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_ActedUponBy(t *testing.T) {
	type testScenario struct {
		name            string
		startingState   State
		eventID         EventIdentifier
		expectedStateID StateIdentifier
		expectedError   error
	}

	var NilStateID StateIdentifier

	testScenarios := []testScenario{
		{
			name: "no events defined",
			startingState: State{
				Events: map[EventIdentifier]StateIdentifier{},
			},
			eventID:         "",
			expectedError:   ErrEventNotDefined,
			expectedStateID: NilStateID,
		},
		{
			name: "event defined",
			startingState: State{
				Events: map[EventIdentifier]StateIdentifier{
					"event": "someState",
				},
			},
			eventID:         "event",
			expectedStateID: "someState",
		},
		{
			name: "events defined; eventID given not valid",
			startingState: State{
				Events: map[EventIdentifier]StateIdentifier{
					"event": "someState",
				},
			},
			eventID:         "invalid",
			expectedError:   ErrEventNotDefined,
			expectedStateID: NilStateID,
		},
	}

	for _, ts := range testScenarios {
		t.Run(ts.name, func(t *testing.T) {
			actualStateID, actualErr := ts.startingState.ActedUponBy(ts.eventID)
			assert.Equal(t, ts.expectedError, actualErr)
			assert.Equal(t, ts.expectedStateID, actualStateID)
		})
	}
}
