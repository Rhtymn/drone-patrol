package app

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/apperror"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/commands"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/constants"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/entity"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/util"
)

type app struct {
	estate *entity.Estate
	drone  entity.Drone
}

func New() app {
	return app{
		drone: entity.NewDrone(1, 1),
	}
}

func (a *app) Run() error {
	err := a.initializing()
	if err != nil {
		return err
	}

	a.patrol()
	return nil
}

func (a *app) initializing() error {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return apperror.Wrap(err)
	}

	estateInf, err := util.IntSlice(line)
	if err != nil {
		return apperror.Wrap(err)
	}

	if len(estateInf) != 3 {
		return apperror.Wrap(err)
	}

	if !util.Range(estateInf[0], 0, 50_000) || !util.Range(estateInf[1], 0, 50_000) || !util.Range(estateInf[2], 0, 50_000) {
		return apperror.Wrap(err)
	}

	estate, err := entity.NewEstate(estateInf[0], estateInf[1])
	palms := make([]entity.Palm, estateInf[2])

	for i := 0; i < estateInf[2]; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return apperror.Wrap(err)
		}

		palmInf, err := util.IntSlice(line)
		if err != nil {
			return apperror.Wrap(err)
		}

		if len(palmInf) != 3 {
			return apperror.Wrap(err)
		}

		palm, err := entity.NewPalm(palmInf[2], palmInf[0], palmInf[1])
		if err != nil {
			return apperror.Wrap(err)
		}

		palms[i] = *palm
	}

	err = estate.SetPlants(palms)
	if err != nil {
		return apperror.Wrap(err)
	}
	a.estate = estate
	return nil
}

func (a *app) patrol() {
	a.drone.Action(commands.MoveUp(), 1, 1)

	palm := a.estate.Plot(1, 1)
	if palm != nil {
		a.drone.Action(commands.MoveUp(), palm.Height(), 1)
	}

loop:
	for a.drone.PosX()+1 <= a.estate.Width() || a.drone.PosY()+1 <= a.estate.Length() || (a.drone.Facing() == constants.West && a.drone.PosX()-1 >= 1) {
		switch a.drone.Facing() {
		case constants.East:
			{
				if a.drone.PosX()+1 > a.estate.Width() {
					nextPalm := a.estate.Plot(a.drone.PosX(), a.drone.PosY()+1)
					if nextPalm != nil && nextPalm.Height()+1 != a.drone.Height() {
						if nextPalm.Height()+1 > a.drone.Height() {
							a.drone.Action(commands.MoveUp(), nextPalm.Height()+1-a.drone.Height(), 1)
							a.drone.Action(commands.MoveNorth(), 1, 10)
						} else {
							a.drone.Action(commands.MoveNorth(), 1, 10)
							a.drone.Action(commands.MoveDown(), a.drone.Height()-(nextPalm.Height()+1), 1)
						}
						a.drone.SetFacing(constants.West)
						continue
					}

					a.drone.Action(commands.MoveNorth(), 1, 10)
					a.drone.SetFacing(constants.West)
					continue
				}

				nextPalm := a.estate.Plot(a.drone.PosX()+1, a.drone.PosY())
				if nextPalm == nil {
					a.drone.Action(commands.MoveEast(), 1, 10)
					continue
				}

				if a.drone.Height() != nextPalm.Height()+1 {
					if nextPalm.Height()+1 > a.drone.Height() {
						a.drone.Action(commands.MoveUp(), nextPalm.Height()+1-a.drone.Height(), 1)
						a.drone.Action(commands.MoveEast(), 1, 10)
					} else {
						a.drone.Action(commands.MoveEast(), 1, 10)
						a.drone.Action(commands.MoveDown(), a.drone.Height()-(nextPalm.Height()+1), 1)
					}
					continue
				}
				a.drone.Action(commands.MoveEast(), 1, 10)
			}
		case constants.West:
			{
				if a.drone.PosX()-1 < 1 {
					if a.drone.PosY() == a.estate.Length() {
						break loop
					}
					nextPalm := a.estate.Plot(a.drone.PosX(), a.drone.PosY()+1)
					if nextPalm != nil && nextPalm.Height()+1 != a.drone.Height() {
						if nextPalm.Height()+1 > a.drone.Height() {
							a.drone.Action(commands.MoveUp(), nextPalm.Height()+1-a.drone.Height(), 1)
							a.drone.Action(commands.MoveNorth(), 1, 10)
						} else {
							a.drone.Action(commands.MoveNorth(), 1, 10)
							a.drone.Action(commands.MoveDown(), a.drone.Height()-(nextPalm.Height()+1), 1)
						}
						a.drone.SetFacing(constants.East)
						continue
					}

					a.drone.Action(commands.MoveNorth(), 1, 10)
					a.drone.SetFacing(constants.East)
					continue
				}

				nextPalm := a.estate.Plot(a.drone.PosX()-1, a.drone.PosY())
				if nextPalm == nil {
					a.drone.Action(commands.MoveWest(), 1, 10)
					continue
				}

				if nextPalm.Height()+1 != a.drone.Height() {
					if nextPalm.Height()+1 > a.drone.Height() {
						a.drone.Action(commands.MoveUp(), nextPalm.Height()+1-a.drone.Height(), 1)
						a.drone.Action(commands.MoveWest(), 1, 10)
					} else {
						a.drone.Action(commands.MoveWest(), 1, 10)
						a.drone.Action(commands.MoveDown(), a.drone.Height()-(nextPalm.Height()+1), 1)
					}
					continue
				}
				a.drone.Action(commands.MoveWest(), 1, 10)
			}
		}
	}
	a.drone.Action(commands.MoveDown(), a.drone.Height(), 1)

	fmt.Println(a.drone.TravelDistance())
}
