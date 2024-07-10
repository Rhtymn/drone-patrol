package commands_test

import (
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/commands"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/entity"
)

func TestWestAction(t *testing.T) {
	t.Run("should move drone to the correct position", func(t *testing.T) {
		posX, posY, dist := 3, 1, 2
		drone := entity.NewDrone(posX, posY)
		w := commands.MoveWest()

		w.Action(&drone, dist, 10)

		if drone.PosX() != posX-dist || drone.PosY() != posY {
			t.Fatalf("drone not move to the expected position")
		}
	})

	t.Run("should correctly calculate drone travel distance", func(t *testing.T) {
		posX, posY, dist := 3, 1, 2
		drone := entity.NewDrone(posX, posY)
		w := commands.MoveWest()

		w.Action(&drone, dist, 10)

		if drone.TravelDistance() != 10*dist {
			t.Fatalf("incorrect calculating drone travel distance")
		}
	})
}
