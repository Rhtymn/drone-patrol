package domain

// Commander is interface for creating an drone commands
// Action method expected for implementing certain drone action
// such as move to west, north, south, and many more
type Commander interface {
	Action(d Droner, dist, meterPerDist int)
}
