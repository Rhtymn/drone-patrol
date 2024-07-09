package entity_test

import (
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/entity"
)

func TestNewEstate(t *testing.T) {
	t.Run("should return estate with correct plotWidth and plotLength", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate := entity.NewEstate(plotWidth, plotLength)

		if estate.PlotWidth() != plotWidth || estate.PlotLength() != plotLength {
			t.Fatalf("estate should have plotWidth %d and plotLength %d, but got plotWidth %d and plotLength %d",
				plotWidth, plotLength, estate.PlotWidth(), estate.PlotLength())
		}
	})

	t.Run("should return estate with correct plots capacity", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate := entity.NewEstate(plotWidth, plotLength)
		plots := estate.Plots()

		if len(plots) != plotLength || len(plots[0]) != plotWidth {
			t.Fatalf("estate should have plots with length %d and width %d, but got length %d and width %d",
				plotLength, plotWidth, len(plots), len(plots[0]))
		}
	})
}

func TestEstatePlotWidth(t *testing.T) {
	t.Run("should return correct estate plot width", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate := entity.NewEstate(plotWidth, plotLength)

		if estate.PlotWidth() != plotWidth {
			t.Fatalf("estate should have plot width %d, but got %d",
				plotWidth, estate.PlotWidth())
		}
	})
}

func TestEstatePloteLength(t *testing.T) {
	t.Run("should return correct estate plot length", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate := entity.NewEstate(plotWidth, plotLength)

		if estate.PlotLength() != plotLength {
			t.Fatalf("estate should have plot length %d, but got %d",
				plotLength, estate.PlotLength())
		}
	})
}

func TestEstatePlots(t *testing.T) {
	t.Run("should return correct estate plots capacity", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate := entity.NewEstate(plotWidth, plotLength)
		plots := estate.Plots()

		if len(plots) != plotLength || len(plots[0]) != plotWidth {
			t.Fatalf("estate should have plots with length %d and width %d, but got length %d and width %d",
				plotLength, plotWidth, len(plots), len(plots[0]))
		}
	})

	t.Run("should not return nil", func(t *testing.T) {
		plotWidth, plotLength := 2, 3

		estate := entity.NewEstate(plotWidth, plotLength)
		plots := estate.Plots()

		if plots == nil {
			t.Fatalf("plots should not nil, but got nil")
		}
	})
}
