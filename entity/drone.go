package entity

type Drone struct {
	travelDistance int
	posX           int
	posY           int
	height         int
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
