package main

import (
	"fmt"
	"os"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/app"
)

func main() {
	app := app.New()

	err := app.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}
}
