package commands

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/constants"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/domain"
)

// west implement commander interface which have Action method
type west struct{}

// MoveWest returning up object
func MoveWest() west {
	return west{}
}

// Action method which represented drone command for moving to west
// each call for this action, the drone will calculating travel distance
// dist representing distance from start position
// meterPerDist representing how much meter for each 1 distance
func (w west) Action(d domain.Droner, dist, meterPerDist int) {
	posX, travelDistance := d.PosX(), d.TravelDistance()
	d.SetPosX(posX - dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
	d.SetFacing(constants.West)
}
