package commands

import "github.com/SawitProRecruitment/JuniorBackendEngineering/domain"

type North struct{}

func MoveNorth() North {
	return North{}
}

func (n North) Action(d domain.Droner, dist, meterPerDist int) {
	posY, travelDistance := d.PosY(), d.TravelDistance()
	d.SetPosY(posY + dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
}
