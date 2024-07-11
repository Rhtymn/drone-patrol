package entity_test

import (
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/entity"
)

func TestNewEstate(t *testing.T) {
	t.Run("should return estate with correct width and length", func(t *testing.T) {
		width, length := 2, 3

		estate, _ := entity.NewEstate(width, length)

		if estate.Width() != width || estate.Length() != length {
			t.Fatalf("estate should have width %d and length %d, but got width %d and length %d",
				width, length, estate.Width(), estate.Length())
		}
	})

	t.Run("should return estate with correct plots capacity", func(t *testing.T) {
		width, length := 2, 3

		estate, _ := entity.NewEstate(width, length)
		plots := estate.Plots()

		if len(plots) != length || len(plots[0]) != width {
			t.Fatalf("estate should have plots with length %d and width %d, but got length %d and width %d",
				length, width, len(plots), len(plots[0]))
		}
	})

	t.Run("should return error InvalidArguments when given width or length with value less than zero", func(t *testing.T) {
		width, length := -2, 3

		_, err := entity.NewEstate(width, length)

		if err == nil || !apperror.ErrorIs(err, apperror.InvalidArguments) {
			t.Fatalf("should return error invalid arguments")
		}
	})
}

func TestSetPlants(t *testing.T) {
	t.Run("should put plant to the correct position", func(t *testing.T) {
		width, length := 2, 2
		h, x, y := 1, 2, 2
		palm, _ := entity.NewPalm(h, x, y)
		palms := []entity.Palm{*palm}
		estate, _ := entity.NewEstate(width, length)

		estate.SetPlants(palms)
		res := estate.Plot(x, y)

		if (res.PosX() != x && res.PosY() != y && res.Height() != h) || res == nil {
			t.Fatalf("plant doesn't plot in the correct position")
		}
	})

	t.Run("should return error out of range position when given palm out of range from estate", func(t *testing.T) {
		width, length := 2, 2
		h, x, y := 1, 2, 3
		palm, _ := entity.NewPalm(h, x, y)
		palms := []entity.Palm{*palm}
		estate, _ := entity.NewEstate(width, length)

		err := estate.SetPlants(palms)

		if err == nil || !apperror.ErrorIs(err, apperror.OutOfRangePosition) {
			t.Fatalf("should return error out of range position")
		}
	})

	t.Run("should return error insufficient plots when given palm more than plot numbers", func(t *testing.T) {
		width, length := 1, 1
		p1, _ := entity.NewPalm(10, 1, 1)
		p2, _ := entity.NewPalm(10, 1, 1)
		palms := []entity.Palm{*p1, *p2}
		estate, _ := entity.NewEstate(width, length)

		err := estate.SetPlants(palms)

		if err == nil || !apperror.ErrorIs(err, apperror.InsufficientPlots) {
			t.Fatalf("should return error insufficient plots")
		}
	})
}

func TestEstateWidth(t *testing.T) {
	t.Run("should return correct estate width", func(t *testing.T) {
		width, length := 2, 3

		estate, _ := entity.NewEstate(width, length)

		if estate.Width() != width {
			t.Fatalf("estate should have width %d, but got %d",
				width, estate.Width())
		}
	})
}

func TestEstateLength(t *testing.T) {
	t.Run("should return correct estate length", func(t *testing.T) {
		plotWidth, length := 2, 3

		estate, _ := entity.NewEstate(plotWidth, length)

		if estate.Length() != length {
			t.Fatalf("estate should have length %d, but got %d",
				length, estate.Length())
		}
	})
}

func TestEstatePlots(t *testing.T) {
	t.Run("should return correct estate plots capacity", func(t *testing.T) {
		width, length := 2, 3

		estate, _ := entity.NewEstate(width, length)
		plots := estate.Plots()

		if len(plots) != length || len(plots[0]) != width {
			t.Fatalf("estate should have plots with length %d and width %d, but got length %d and width %d",
				length, width, len(plots), len(plots[0]))
		}
	})

	t.Run("should not return nil", func(t *testing.T) {
		width, length := 2, 3

		estate, _ := entity.NewEstate(width, length)
		plots := estate.Plots()

		if plots == nil {
			t.Fatalf("plots should not nil, but got nil")
		}
	})
}
