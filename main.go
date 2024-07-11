package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/commands"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/constants"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/entity"
	"github.com/SawitProRecruitment/JuniorBackendEngineering/util"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}

	estateInf, err := util.IntSlice(line)
	if err != nil {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}

	if len(estateInf) != 3 {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}

	if !util.Range(estateInf[0], 0, 50_000) || !util.Range(estateInf[1], 0, 50_000) || !util.Range(estateInf[2], 0, 50_000) {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}

	estate, err := entity.NewEstate(estateInf[0], estateInf[1])
	palms := make([]entity.Palm, estateInf[2])

	for i := 0; i < estateInf[2]; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}

		palmInf, err := util.IntSlice(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}

		if len(palmInf) != 3 {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}

		palm, err := entity.NewPalm(palmInf[2], palmInf[0], palmInf[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "FAIL")
			os.Exit(1)
		}

		palms[i] = *palm
	}

	err = estate.SetPlants(palms)
	if err != nil {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}

	drone := entity.NewDrone(1, 1)
	drone.Action(commands.MoveUp(), 1, 1)

	palm := estate.Plot(1, 1)
	if palm != nil {
		drone.Action(commands.MoveUp(), palm.Height(), 1)
	}

loop:
	for drone.PosX()+1 <= estate.Width() || drone.PosY()+1 <= estate.Length() || (drone.Facing() == constants.West && drone.PosX()-1 >= 1) {
		switch drone.Facing() {
		case constants.East:
			{
				if drone.PosX()+1 > estate.Width() {
					nextPalm := estate.Plot(drone.PosX(), drone.PosY()+1)
					if nextPalm != nil && nextPalm.Height()+1 != drone.Height() {
						if nextPalm.Height()+1 > drone.Height() {
							drone.Action(commands.MoveUp(), nextPalm.Height()+1-drone.Height(), 1)
							drone.Action(commands.MoveNorth(), 1, 10)
						} else {
							drone.Action(commands.MoveNorth(), 1, 10)
							drone.Action(commands.MoveDown(), drone.Height()-(nextPalm.Height()+1), 1)
						}
						drone.SetFacing(constants.West)
						continue
					}

					drone.Action(commands.MoveNorth(), 1, 10)
					drone.SetFacing(constants.West)
					continue
				}

				nextPalm := estate.Plot(drone.PosX()+1, drone.PosY())
				if nextPalm == nil {
					drone.Action(commands.MoveEast(), 1, 10)
					continue
				}

				if drone.Height() != nextPalm.Height()+1 {
					if nextPalm.Height()+1 > drone.Height() {
						drone.Action(commands.MoveUp(), nextPalm.Height()+1-drone.Height(), 1)
						drone.Action(commands.MoveEast(), 1, 10)
					} else {
						drone.Action(commands.MoveEast(), 1, 10)
						drone.Action(commands.MoveDown(), drone.Height()-(nextPalm.Height()+1), 1)
					}
					continue
				}
				drone.Action(commands.MoveEast(), 1, 10)
			}
		case constants.West:
			{
				if drone.PosX()-1 < 1 {
					if drone.PosY() == estate.Length() {
						break loop
					}
					nextPalm := estate.Plot(drone.PosX(), drone.PosY()+1)
					if nextPalm != nil && nextPalm.Height()+1 != drone.Height() {
						if nextPalm.Height()+1 > drone.Height() {
							drone.Action(commands.MoveUp(), nextPalm.Height()+1-drone.Height(), 1)
							drone.Action(commands.MoveNorth(), 1, 10)
						} else {
							drone.Action(commands.MoveNorth(), 1, 10)
							drone.Action(commands.MoveDown(), drone.Height()-(nextPalm.Height()+1), 1)
						}
						drone.SetFacing(constants.East)
						continue
					}

					drone.Action(commands.MoveNorth(), 1, 10)
					drone.SetFacing(constants.East)
					continue
				}

				nextPalm := estate.Plot(drone.PosX()-1, drone.PosY())
				if nextPalm == nil {
					drone.Action(commands.MoveWest(), 1, 10)
					continue
				}

				if nextPalm.Height()+1 != drone.Height() {
					if nextPalm.Height()+1 > drone.Height() {
						drone.Action(commands.MoveUp(), nextPalm.Height()+1-drone.Height(), 1)
						drone.Action(commands.MoveWest(), 1, 10)
					} else {
						drone.Action(commands.MoveWest(), 1, 10)
						drone.Action(commands.MoveDown(), drone.Height()-(nextPalm.Height()+1), 1)
					}
					continue
				}
				drone.Action(commands.MoveWest(), 1, 10)
			}
		}
	}
	drone.Action(commands.MoveDown(), drone.Height(), 1)

	fmt.Println(drone.TravelDistance())
}
