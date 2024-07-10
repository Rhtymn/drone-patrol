package entity

import "github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"

type Palm struct {
	height int
	posX   int
	posY   int
}

func NewPalm(height, posX, posY int) (*Palm, error) {
	if height < 0 {
		return nil, apperror.NewInvalidArguments("height must be more than or equal zero")
	}
	return &Palm{
		height: height,
		posX:   posX,
		posY:   posY,
	}, nil
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
