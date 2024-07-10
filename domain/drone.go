package domain

import "github.com/SawitProRecruitment/JuniorBackendEngineering/constants"

type Droner interface {
	Height() int
	PosX() int
	PosY() int
	TravelDistance() int

	SetHeight(h int)
	SetPosX(x int)
	SetPosY(y int)
	SetTravelDistance(dist int)
	SetFacing(dir constants.Direction)
}
