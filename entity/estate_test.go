package entity_test

import (
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/entity"
)

func TestNewEstate(t *testing.T) {
	t.Run("should return estate with correct plotWidth and plotLength", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate, _ := entity.NewEstate(plotWidth, plotLength)

		if estate.PlotWidth() != plotWidth || estate.PlotLength() != plotLength {
			t.Fatalf("estate should have plotWidth %d and plotLength %d, but got plotWidth %d and plotLength %d",
				plotWidth, plotLength, estate.PlotWidth(), estate.PlotLength())
		}
	})

	t.Run("should return estate with correct plots capacity", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate, _ := entity.NewEstate(plotWidth, plotLength)
		plots := estate.Plots()

		if len(plots) != plotLength || len(plots[0]) != plotWidth {
			t.Fatalf("estate should have plots with length %d and width %d, but got length %d and width %d",
				plotLength, plotWidth, len(plots), len(plots[0]))
		}
	})

	t.Run("should return error InvalidArguments when given plotWidth or plotLength with value less than zero", func(t *testing.T) {
		plotWidth, plotLength := -2, 3

		_, err := entity.NewEstate(plotWidth, plotLength)

		if err == nil || !apperror.ErrorIs(err, apperror.InvalidArguments) {
			t.Fatalf("should return error invalid arguments")
		}
	})
}

func TestSetPlants(t *testing.T) {
	t.Run("should put plant to the correct position", func(t *testing.T) {
		plotWidth, plotLength := 2, 2
		h, x, y := 1, 2, 2
		palm, _ := entity.NewPalm(h, x, y)
		palms := []entity.Palm{*palm}
		estate, _ := entity.NewEstate(plotWidth, plotLength)

		estate.SetPlants(palms)
		res := estate.Plot(x, y)

		if (res.PosX() != x && res.PosY() != y && res.Height() != h) || res == nil {
			t.Fatalf("plant doesn't plot in the correct position")
		}
	})

	t.Run("should return error out of range position when given palm out of range from estate", func(t *testing.T) {
		plotWidth, plotLength := 2, 2
		h, x, y := 1, 2, 3
		palm, _ := entity.NewPalm(h, x, y)
		palms := []entity.Palm{*palm}
		estate, _ := entity.NewEstate(plotWidth, plotLength)

		err := estate.SetPlants(palms)

		if err == nil || !apperror.ErrorIs(err, apperror.OutOfRangePosition) {
			t.Fatalf("should return error out of range position")
		}
	})

	t.Run("should return error insufficient plots when given palm more than plot numbers", func(t *testing.T) {
		plotWidth, plotLength := 1, 1
		p1, _ := entity.NewPalm(10, 1, 1)
		p2, _ := entity.NewPalm(10, 1, 1)
		palms := []entity.Palm{*p1, *p2}
		estate, _ := entity.NewEstate(plotWidth, plotLength)

		err := estate.SetPlants(palms)

		if err == nil || !apperror.ErrorIs(err, apperror.InsufficientPlots) {
			t.Fatalf("should return error insufficient plots")
		}
	})
}

func TestEstatePlotWidth(t *testing.T) {
	t.Run("should return correct estate plot width", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate, _ := entity.NewEstate(plotWidth, plotLength)

		if estate.PlotWidth() != plotWidth {
			t.Fatalf("estate should have plot width %d, but got %d",
				plotWidth, estate.PlotWidth())
		}
	})
}

func TestEstatePloteLength(t *testing.T) {
	t.Run("should return correct estate plot length", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate, _ := entity.NewEstate(plotWidth, plotLength)

		if estate.PlotLength() != plotLength {
			t.Fatalf("estate should have plot length %d, but got %d",
				plotLength, estate.PlotLength())
		}
	})
}

func TestEstatePlots(t *testing.T) {
	t.Run("should return correct estate plots capacity", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate, _ := entity.NewEstate(plotWidth, plotLength)
		plots := estate.Plots()

		if len(plots) != plotLength || len(plots[0]) != plotWidth {
			t.Fatalf("estate should have plots with length %d and width %d, but got length %d and width %d",
				plotLength, plotWidth, len(plots), len(plots[0]))
		}
	})

	t.Run("should not return nil", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate, _ := entity.NewEstate(plotWidth, plotLength)
		plots := estate.Plots()

		if plots == nil {
			t.Fatalf("plots should not nil, but got nil")
		}
	})
}
