package constants

// Direction is predeclared identifier representing basic compass points type
type Direction uint

// It's constants for representing compass point
const (
	North = iota + 1
	East
	South
	West
)
