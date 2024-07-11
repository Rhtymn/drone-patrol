package commands

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/constants"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/domain"
)

// south implement commander interface which have Action method
type south struct{}

// MoveSouth returning north object
func MoveSouth() south {
	return south{}
}

// Action method which represented drone command for moving to south
// each call for this action, the drone will calculating travel distance
// dist representing distance from start position
// meterPerDist representing how much meter for each 1 distance
func (s south) Action(d domain.Droner, dist, meterPerDist int) {
	posY, travelDistance := d.PosY(), d.TravelDistance()
	d.SetPosY(posY - dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
	d.SetFacing(constants.South)
}
