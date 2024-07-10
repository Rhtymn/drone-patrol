package commands

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/domain"
)

type Down struct{}

func MoveDown() Down {
	return Down{}
}

func (dn Down) Action(d domain.Droner, dist, meterPerDist int) {
	h, travelDistance := d.Height(), d.TravelDistance()
	d.SetHeight(h - dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
}
