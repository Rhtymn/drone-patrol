package commands

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/domain"
)

type down struct{}

func MoveDown() down {
	return down{}
}

func (dn down) Action(d domain.Droner, dist, meterPerDist int) {
	h, travelDistance := d.Height(), d.TravelDistance()
	d.SetHeight(h - dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
}
