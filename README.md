# godel
Godel is pronounced like yodel. Or maybe like yah-del. Or, if you are a mathematician, like the greatest logician to 
ever walk with Einstein.

Simple and basic modelling in go. 

## Reasoning
This repo was created with the goal of creating an easy way to model state machines in go. I found myself struggling
with other modelling languages and wanted to explore potentially interop with go applications themselves through annotations,
code generators, test assistant code, or similar. I also wanted to just understand state machines and modelling better; no better way
than doing it yourself. 


## TODO: 
- [x] Make sure we can marshal and unmarshal JSON effectively
  - [x] How can we clean this up? Are all actions and states needed? Better way to display?
- [ ] Create and model out how to create and spec out a machine
  - [ ] Make sure that it is easy to execute the machine for testing and such
- [ ] Explore some kind of simple graphical output
  - [ ] Start with arrows, then nodes, and then the whole model
    - [ ] Output to markdown with a simple model and description would be very helpful
- [ ] Could we explore annotating go code to read and produce a state machine model?


## Scratch Notes
- Define interfaces for states and events.
- Want user to define events, states, and specifications.
- How do we let a user define how to test the system?
- How do we make assertions against the specs?

- need to have an OnEvent method for each event for a state. Then when in a state, we clearly see what happens when an event happens
- could just have an interface for state that defines OnEvent - the implementation is then in charge of sorting which event and what to do