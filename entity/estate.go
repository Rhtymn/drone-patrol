package entity

import "github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"

// Estate representing estate object which have width, length and plots
type Estate struct {
	width  int
	length int

	plots [][]*Palm
}

// NewEstate returning estate object which initialized with given width and length value
func NewEstate(width, length int) (*Estate, error) {
	if width <= 0 || length <= 0 {
		return nil, apperror.NewInvalidArguments("width and length must be more than zero")
	}

	plots := make([][]*Palm, length)
	for i := 0; i < len(plots); i++ {
		plots[i] = make([]*Palm, width)
	}

	return &Estate{
		width:  width,
		length: length,
		plots:  plots,
	}, nil
}

// SetPlants will ploting given palm on estate plots representation
// if palm position given is out of range estate plots, it will return error out of range position
// if estate plots capacity is not enough for palm plotting, it will return error insufficient plots
func (e *Estate) SetPlants(palms []Palm) error {
	if len(palms) > e.length*e.width {
		return apperror.NewInsufficientPlots()
	}

	for i := 0; i < len(palms); i++ {
		palm := palms[i]
		if !e.inEstate(palm.posX, palm.posY) {
			e.resetPlots()
			return apperror.NewOutOfRangePosition("plant position out of range from estate")
		}
		e.plots[palm.posY-1][palm.posX-1] = &palm
	}
	return nil
}

// inEstate checking is given x and y position given is in estate range
func (e *Estate) inEstate(x, y int) bool {
	return x <= e.width && x > 0 && y > 0 && y <= e.length
}

// resetPlots removing all plant plotting
func (e *Estate) resetPlots() {
	plots := make([][]*Palm, e.length)
	for i := 0; i < len(plots); i++ {
		plots[i] = make([]*Palm, e.width)
	}
}

// Plot returning palm which on given position x and y
// if there's no palm on given position, it will returning nil
func (e Estate) Plot(x, y int) *Palm {
	return e.plots[y-1][x-1]
}

// Getter width
func (e Estate) Width() int {
	return e.width
}

// Getter length
func (e Estate) Length() int {
	return e.length
}

// Getter plots
func (e Estate) Plots() [][]*Palm {
	return e.plots
}
