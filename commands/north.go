package commands

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/constants"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/domain"
)

type north struct{}

func MoveNorth() north {
	return north{}
}

func (n north) Action(d domain.Droner, dist, meterPerDist int) {
	posY, travelDistance := d.PosY(), d.TravelDistance()
	d.SetPosY(posY + dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
	d.SetFacing(constants.North)
}
