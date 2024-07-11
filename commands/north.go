package commands

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/constants"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/domain"
)

// north implement commander interface which have Action method
type north struct{}

// MoveNorth returning north object
func MoveNorth() north {
	return north{}
}

// Action method which represented drone command for moving to north
// each call for this action, the drone will calculating travel distance
// dist representing distance from start position
// meterPerDist representing how much meter for each 1 distance
func (n north) Action(d domain.Droner, dist, meterPerDist int) {
	posY, travelDistance := d.PosY(), d.TravelDistance()
	d.SetPosY(posY + dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
	d.SetFacing(constants.North)
}
