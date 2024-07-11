package entity_test

import (
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/entity"
)

func TestNewPalm(t *testing.T) {
	t.Run("should return palm with correct height and position", func(t *testing.T) {
		h, posX, posY := 5, 1, 2

		palm, _ := entity.NewPalm(h, posX, posY)

		if palm.Height() != h || palm.PosX() != posX || palm.PosY() != posY {
			t.Fatalf("palm should have height %d and position (%d, %d), but got %d and (%d, %d)",
				h, posX, posY, palm.Height(), palm.PosX(), palm.PosY())
		}
	})

	t.Run("should return error InvalidArguments if given height with value less than zero", func(t *testing.T) {
		h, posX, posY := -3, 0, 1

		_, err := entity.NewPalm(h, posX, posY)

		if err == nil || !apperror.ErrorIs(err, apperror.InvalidArguments) {
			t.Fatalf("should return error invalid arguments")
		}
	})

	t.Run("should return error InvalidArguments if given height with value not in range 1 to 30 (inclusive)", func(t *testing.T) {
		h, posX, posY := 40, 1, 1

		_, err := entity.NewPalm(h, posX, posY)

		if err == nil || !apperror.ErrorIs(err, apperror.InvalidArguments) {
			t.Fatalf("should return error invalid arguments")
		}
	})

	t.Run("should not return error if given height with value range 1 to 30 (inclusive)", func(t *testing.T) {
		h, posX, posY := 15, 1, 1

		_, err := entity.NewPalm(h, posX, posY)

		if err != nil {
			t.Fatalf("should not return any error")
		}
	})
}

func TestPalmPosX(t *testing.T) {
	t.Run("should return correct palm position x", func(t *testing.T) {
		h, posX, posY := 5, 1, 2

		palm, _ := entity.NewPalm(h, posX, posY)

		if palm.PosX() != posX {
			t.Fatalf("palm should return position of x %d, but got %d",
				posX, palm.PosX())
		}
	})
}

func TestPalmPosY(t *testing.T) {
	t.Run("should return correct palm position y", func(t *testing.T) {
		h, posX, posY := 5, 1, 2

		palm, _ := entity.NewPalm(h, posX, posY)

		if palm.PosY() != posY {
			t.Fatalf("palm should return position of y %d, but got %d",
				posY, palm.PosY())
		}
	})
}

func TestHeight(t *testing.T) {
	t.Run("should return correct palm height", func(t *testing.T) {
		h, posX, posY := 5, 1, 2

		palm, _ := entity.NewPalm(h, posX, posY)

		if palm.Height() != h {
			t.Fatalf("palm should return height %d, but got %d",
				h, palm.Height())
		}
	})
}
