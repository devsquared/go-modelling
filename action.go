package godel

// Action defines something that happens to a state. For the sake of simplicity here, a State knows and defines the actions that
// can happen to it. This is because the context of the original State is important for the outcome; not simply the action itself.
type Action struct {
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	ResultState State  `json:"resultState"`
}
