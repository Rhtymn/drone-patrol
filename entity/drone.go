package entity

import (
	"github.com/SawitProRecruitment/JuniorBackendEngineering/constants"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/domain"
)

// Drone representing drone object which have some attribute
type Drone struct {
	travelDistance int                 // drone travel/move distance
	posX           int                 // drone position x
	posY           int                 // drone position y
	height         int                 // drone height
	facing         constants.Direction // drone facing direction
}

// NewDrone returning drone which initialized with given posX and posY value
func NewDrone(posX, posY int) Drone {
	return Drone{
		posX:   posX,
		posY:   posY,
		facing: constants.East,
	}
}

// Reset all drone attribute
func (d *Drone) Reset() {
	d.travelDistance = 0
	d.facing = constants.East
	d.height = 0
	d.posX = 1
	d.posY = 1
}

// Setter posX
func (d *Drone) SetPosX(x int) {
	d.posX = x
}

// Setter posY
func (d *Drone) SetPosY(y int) {
	d.posY = y
}

// Setter travelDistance
func (d *Drone) SetTravelDistance(dist int) {
	d.travelDistance = dist
}

// Setter height
func (d *Drone) SetHeight(h int) {
	d.height = h
}

// Setter facing
func (d *Drone) SetFacing(dir constants.Direction) {
	d.facing = dir
}

func (d *Drone) Action(command domain.Commander, dist, meterPerDist int) {
	command.Action(d, dist, meterPerDist)
}

// Getter posX
func (d Drone) PosX() int {
	return d.posX
}

// Getter posY
func (d Drone) PosY() int {
	return d.posY
}

// Getter height
func (d Drone) Height() int {
	return d.height
}

// Getter travelDistance
func (d Drone) TravelDistance() int {
	return d.travelDistance
}

// Getter facing
func (d Drone) Facing() constants.Direction {
	return d.facing
}
