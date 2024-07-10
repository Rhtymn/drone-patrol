package entity

import "github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"

type Palm struct {
	height int
	posX   int
	posY   int
}

func NewPalm(height, posX, posY int) (*Palm, error) {
	if height < 1 || height > 30 {
		return nil, apperror.NewInvalidArguments("height must be in range from 1 to 30 (inclusive)")
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
