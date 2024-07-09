package entity

type Palm struct {
	height int
	posX   int
	posY   int
}

func NewPalm(height, posX, posY int) Palm {
	return Palm{
		height: height,
		posX:   posX,
		posY:   posY,
	}
}

func (p Palm) PosX() int {
	return p.posX
}

func (p Palm) PosY() int {
	return p.posY
}

func (p Palm) Height() int {
	return p.height
}
