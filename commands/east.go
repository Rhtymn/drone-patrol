package commands

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/constants"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/domain"
)

type east struct{}

func MoveEast() east {
	return east{}
}

func (e east) Action(d domain.Droner, dist, meterPerDist int) {
	posX, travelDistance := d.PosX(), d.TravelDistance()
	d.SetPosX(posX + dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
	d.SetFacing(constants.East)
}
