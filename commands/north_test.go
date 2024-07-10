package commands_test

import (
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/commands"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/constants"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/entity"
)

func TestNorthAction(t *testing.T) {
	t.Run("should move drone to the correct position", func(t *testing.T) {
		posX, posY, dist := 3, 1, 2
		drone := entity.NewDrone(posX, posY)
		n := commands.MoveNorth()

		n.Action(&drone, dist, 10)

		if drone.PosY() != posY+dist || drone.PosX() != posX {
			t.Fatalf("drone not move to the expected position")
		}
	})

	t.Run("should correctly calculate drone travel distance", func(t *testing.T) {
		posX, posY, dist := 3, 1, 2
		drone := entity.NewDrone(posX, posY)
		w := commands.MoveNorth()

		w.Action(&drone, dist, 10)

		if drone.TravelDistance() != 10*dist {
			t.Fatalf("incorrect calculating drone travel distance")
		}
	})

	t.Run("should make drone to face on north", func(t *testing.T) {
		posX, posY, dist := 1, 1, 1
		drone := entity.NewDrone(posX, posY)
		e := commands.MoveNorth()

		e.Action(&drone, dist, 10)

		if drone.Facing() != constants.North {
			t.Fatalf("incorrect drone facing")
		}
	})
}
