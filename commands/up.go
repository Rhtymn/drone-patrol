package commands

import "github.com/SawitProRecruitment/JuniorBackendEngineering/domain"

// up implement commander interface which have Action method
type up struct{}

// MoveUp returning up object
func MoveUp() up {
	return up{}
}

// Action method which represented drone command for moving to up
// each call for this action, the drone will calculating travel distance
// dist representing distance from start position
// meterPerDist representing how much meter for each 1 distance
func (u up) Action(d domain.Droner, dist, meterPerDist int) {
	h, travelDistance := d.Height(), d.TravelDistance()
	d.SetHeight(h + dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
}
