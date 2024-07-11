package commands

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/domain"
)

// down implement commander interface which have Action method
type down struct{}

// MoveDown returning down object
func MoveDown() down {
	return down{}
}

// Action method which represented drone command for moving to down
// each call for this action, the drone will calculating travel distance
// dist representing distance from start position
// meterPerDist representing how much meter for each 1 distance
func (dn down) Action(d domain.Droner, dist, meterPerDist int) {
	h, travelDistance := d.Height(), d.TravelDistance()
	d.SetHeight(h - dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
}
