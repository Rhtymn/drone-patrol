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

func TestSetPosX(t *testing.T) {
	t.Run("should set drone position x to the correct position", func(t *testing.T) {
		posX, posY := 1, 1
		newPosX := 2
		drone := entity.NewDrone(posX, posY)

		drone.SetPosX(newPosX)

		if drone.PosX() != newPosX {
			t.Fatalf("incorrect drone position x")
		}
	})
}

func TestSetPosY(t *testing.T) {
	t.Run("should set drone position y to the correct position", func(t *testing.T) {
		posX, posY := 1, 1
		newPosY := 2
		drone := entity.NewDrone(posX, posY)

		drone.SetPosY(newPosY)

		if drone.PosY() != newPosY {
			t.Fatalf("incorrect drone position y")
		}
	})
}

func TestSetTravelDistance(t *testing.T) {
	t.Run("should set drone travel distance to the correct value given", func(t *testing.T) {
		posX, posY := 1, 1
		travelDistance := 3
		drone := entity.NewDrone(posX, posY)

		drone.SetTravelDistance(travelDistance)

		if drone.TravelDistance() != travelDistance {
			t.Fatalf("incorrect drone travel distance")
		}
	})
}

func TestSetHeight(t *testing.T) {
	t.Run("should set drone height to the correct value given", func(t *testing.T) {
		posX, posY := 1, 1
		height := 3
		drone := entity.NewDrone(posX, posY)

		drone.SetHeight(height)

		if drone.Height() != height {
			t.Fatalf("incorrect drone height")
		}
	})
}

func TestDronePosX(t *testing.T) {
	t.Run("should return correct drone position x", func(t *testing.T) {
		posX, posY := 1, 1
		drone := entity.NewDrone(posX, posY)
		drone.SetPosX(2)

		res := drone.PosX()

		if res != 2 {
			t.Fatalf("incorrect drone position x")
		}
	})
}

func TestDronePosY(t *testing.T) {
	t.Run("should return correct drone position y", func(t *testing.T) {
		posX, posY := 1, 1
		drone := entity.NewDrone(posX, posY)
		drone.SetPosY(3)

		res := drone.PosY()

		if res != 3 {
			t.Fatalf("incorrect drone position x")
		}
	})
}

func TestDroneTravelDistance(t *testing.T) {
	t.Run("should return correct drone travel distance", func(t *testing.T) {
		posX, posY := 1, 1
		drone := entity.NewDrone(posX, posY)
		drone.SetTravelDistance(4)

		res := drone.TravelDistance()

		if res != 4 {
			t.Fatalf("incorrect drone travel distance")
		}
	})
}

func TestDroneHeight(t *testing.T) {
	t.Run("should return correct drone height", func(t *testing.T) {
		posX, posY := 1, 1
		drone := entity.NewDrone(posX, posY)
		drone.SetHeight(5)

		res := drone.Height()

		if res != 5 {
			t.Fatalf("incorrect drone height")
		}
	})
}
