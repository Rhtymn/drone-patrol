package commands

import "github.com/SawitProRecruitment/JuniorBackendEngineering/domain"

type West struct{}

func MoveWest() West {
	return West{}
}

func (w West) Action(d domain.Droner, dist, meterPerDist int) {
	posX, travelDistance := d.PosX(), d.TravelDistance()
	d.SetPosX(posX - dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
}
