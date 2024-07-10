package entity

import "github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"

type Estate struct {
	plotWidth  int
	plotLength int

	plots [][]*Palm
}

func NewEstate(plotWidth, plotLength int) (*Estate, error) {
	if plotWidth <= 0 || plotLength <= 0 {
		return nil, apperror.NewInvalidArguments("plotWidth and plotLength must be more than zero")
	}

	plots := make([][]*Palm, plotLength)
	for i := 0; i < len(plots); i++ {
		plots[i] = make([]*Palm, plotWidth)
	}

	return &Estate{
		plotWidth:  plotWidth,
		plotLength: plotLength,
		plots:      plots,
	}, nil
}

func (e *Estate) SetPlants(palms []Palm) {
	for i := 0; i < len(palms); i++ {
		palm := palms[i]
		e.plots[palm.posX][palm.posY] = &palm
	}
}

func (e Estate) PlotWidth() int {
	return e.plotWidth
}

func (e Estate) PlotLength() int {
	return e.plotLength
}

func (e Estate) Plots() [][]*Palm {
	return e.plots
}
