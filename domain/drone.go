package domain

import "github.com/SawitProRecruitment/JuniorBackendEngineering/constants"

// Droner is interface for creating drone object/struct
// It's contains basic getter and setter which expected
// each drone have this functionality
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
