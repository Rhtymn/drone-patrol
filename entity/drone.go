package entity

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/constants"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/domain"
)

type Drone struct {
	travelDistance int
	posX           int
	posY           int
	height         int
	facing         constants.Direction
}

func NewDrone(posX, posY int) Drone {
	return Drone{
		posX:   posX,
		posY:   posY,
		facing: constants.East,
	}
}

func (d *Drone) Reset() {
	d.travelDistance = 0
	d.facing = constants.East
	d.height = 0
	d.posX = 1
	d.posY = 1
}

func (d *Drone) SetPosX(x int) {
	d.posX = x
}

func (d *Drone) SetPosY(y int) {
	d.posY = y
}

func (d *Drone) SetTravelDistance(dist int) {
	d.travelDistance = dist
}

func (d *Drone) SetHeight(h int) {
	d.height = h
}

func (d *Drone) SetFacing(dir constants.Direction) {
	d.facing = dir
}

func (d *Drone) Action(command domain.Commander, dist, meterPerDist int) {
	command.Action(d, dist, meterPerDist)
}

func (d Drone) PosX() int {
	return d.posX
}

func (d Drone) PosY() int {
	return d.posY
}

func (d Drone) Height() int {
	return d.height
}

func (d Drone) TravelDistance() int {
	return d.travelDistance
}

func (d Drone) Facing() constants.Direction {
	return d.facing
}
