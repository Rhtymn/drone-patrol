package domain

type Commander interface {
	Action(d Droner, dist, meterPerDist int)
}
