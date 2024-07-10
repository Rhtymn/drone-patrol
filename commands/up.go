package commands

import "github.com/SawitProRecruitment/JuniorBackendEngineering/domain"

type up struct{}

func MoveUp() up {
	return up{}
}

func (u up) Action(d domain.Droner, dist, meterPerDist int) {
	h, travelDistance := d.Height(), d.TravelDistance()
	d.SetHeight(h + dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
}
