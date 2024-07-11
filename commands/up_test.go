package commands_test

import (
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/commands"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/entity"
)

func TestUpAction(t *testing.T) {
	t.Run("should move drone to the correct position", func(t *testing.T) {
		posX, posY, dist := 3, 3, 2
		drone := entity.NewDrone(posX, posY)
		u := commands.MoveUp()

		u.Action(&drone, dist, 1)

		if drone.Height() != dist {
			t.Fatalf("drone not move to the expected position")
		}
	})

	t.Run("should correctly calculate drone travel distance", func(t *testing.T) {
		posX, posY, dist := 3, 3, 2
		drone := entity.NewDrone(posX, posY)
		u := commands.MoveUp()

		u.Action(&drone, dist, 1)

		if drone.TravelDistance() != dist {
			t.Fatalf("incorrect calculating drone travel distance")
		}
	})
}
