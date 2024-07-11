package entity

import "github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"

// Palm representing palm plants which have height, posX, and posY
type Palm struct {
	height int
	posX   int
	posY   int
}

// NewPalm returning palm object which has initialized with given height, posX, and posY value
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

// Getter posX
func (p Palm) PosX() int {
	return p.posX
}

// Getter posY
func (p Palm) PosY() int {
	return p.posY
}

// Getter height
func (p Palm) Height() int {
	return p.height
}
