package domain

type Droner interface {
	Height() int
	PosX() int
	PosY() int
	TravelDistance() int

	SetHeight(h int)
	SetPosX(x int)
	SetPosY(y int)
	SetTravelDistance(dist int)
}
