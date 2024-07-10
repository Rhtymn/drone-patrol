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

func (e *Estate) SetPlants(palms []Palm) error {
	for i := 0; i < len(palms); i++ {
		palm := palms[i]
		if !e.inEstate(palm.posX, palm.posY) {
			e.resetPlots()
			return apperror.NewOutOfRangePosition("plant position out of range from estate")
		}
		e.plots[palm.posX-1][palm.posY-1] = &palm
	}
	return nil
}

func (e *Estate) inEstate(x, y int) bool {
	return x <= e.plotLength && x > 0 && y > 0 && y <= e.plotWidth
}

func (e *Estate) resetPlots() {
	plots := make([][]*Palm, e.plotLength)
	for i := 0; i < len(plots); i++ {
		plots[i] = make([]*Palm, e.plotWidth)
	}
}

func (e Estate) Plot(x, y int) *Palm {
	return e.plots[x-1][y-1]
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
