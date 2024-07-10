package entity

type Drone struct {
	travelDistance int
	posX           int
	posY           int
}

func NewDrone(posX, posY int) Drone {
	return Drone{
		posX: posX,
		posY: posY,
	}
}

func (d *Drone) Reset() {
	d.travelDistance = 0
}

func (d Drone) PosX() int {
	return d.posX
}

func (d Drone) PosY() int {
	return d.posY
}

func (d Drone) TravelDistance() int {
	return d.travelDistance
}
