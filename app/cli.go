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

// app representing cli application
type app struct {
	estate *entity.Estate
	drone  entity.Drone
}

// New returning app object with drone already created
func New() app {
	return app{
		drone: entity.NewDrone(1, 1),
	}
}

// Run will be start application
func (a *app) Run() error {
	err := a.initializing()
	if err != nil {
		return err
	}

	a.patrol()
	return nil
}

// initializing will do estate initialization and palm plotting
func (a *app) initializing() error {
	reader := bufio.NewReader(os.Stdin)
	// read first line input which representing estate information
	line, err := reader.ReadString('\n')
	if err != nil {
		return apperror.Wrap(err)
	}

	// transform string to slice of int
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

	// initializing estate from given input
	estate, err := entity.NewEstate(estateInf[0], estateInf[1])
	// create array of palms with length from given input
	palms := make([]entity.Palm, estateInf[2])

	// loop for reading next line until last line
	for i := 0; i < estateInf[2]; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return apperror.Wrap(err)
		}

		// transform string to slice of int
		palmInf, err := util.IntSlice(line)
		if err != nil {
			return apperror.Wrap(err)
		}

		if len(palmInf) != 3 {
			return apperror.Wrap(err)
		}

		// create palm from given input information
		palm, err := entity.NewPalm(palmInf[2], palmInf[0], palmInf[1])
		if err != nil {
			return apperror.Wrap(err)
		}

		// add created palm to array
		palms[i] = *palm
	}

	// for each created palms, plot to estate
	err = estate.SetPlants(palms)
	if err != nil {
		return apperror.Wrap(err)
	}
	a.estate = estate
	return nil
}

// patrol will start drone patrol for estate which already initialized
// in the end of patrol, it will print drone travel distance
func (a *app) patrol() {
	// move up to 1 meter height at first
	a.drone.Action(commands.MoveUp(), 1, 1)

	palm := a.estate.Plot(1, 1)
	// if there's palm on first plot, move up drone until height is palm heigth plus one
	// drone heigth = palm height + 1
	if palm != nil {
		a.drone.Action(commands.MoveUp(), palm.Height(), 1)
	}

	// doing patrol until last plot
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
	// move drone to down until ground
	a.drone.Action(commands.MoveDown(), a.drone.Height(), 1)

	// print drone travel distance
	fmt.Println(a.drone.TravelDistance())
}
