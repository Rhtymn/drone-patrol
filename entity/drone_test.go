package entity_test

import (
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/entity"
)

func TestNewDrone(t *testing.T) {
	t.Run("should return drone with zero value on each attribute", func(t *testing.T) {
		posX, posY := 1, 1

		drone := entity.NewDrone(posX, posY)

		if drone.PosX() != 0 && drone.PosY() != 0 && drone.TravelDistance() != 0 {
			t.Fatalf("drone should have pos (0, 0) and travel distance 0 but got (%d, %d) and %d",
				drone.PosX(), drone.PosY(), drone.TravelDistance())
		}
	})
}

func TestDronePosX(t *testing.T) {
	t.Run("should return correct drone position x", func(t *testing.T) {

	})
}

func TestDronePosY(t *testing.T) {
	t.Run("should return correct drone position y", func(t *testing.T) {

	})
}

func TestDroneTravelDistance(t *testing.T) {
	t.Run("should return correct drone travel distance", func(t *testing.T) {})
}
