package commands

import "github.com/SawitProRecruitment/JuniorBackendEngineering/domain"

type Up struct{}

func MoveUp() Up {
	return Up{}
}

func (u Up) Action(d domain.Droner, dist, meterPerDist int) {
	h, travelDistance := d.Height(), d.TravelDistance()
	d.SetHeight(h + dist)
	d.SetTravelDistance(travelDistance + meterPerDist*dist)
}
