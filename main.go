package main

import (
	"github.com/crypto89/expedition-simulator/simulator"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func main() {
	spew.Dump(simulator.Consumption(50, 400))
	return

	sim, _ := simulator.NewSimulator(&simulator.Config{
		Rounds:           4000,
		ExpeditionPoints: 9000,
		SpeedFactor:      8,
		Cargo:            simulator.CARGO_LC,
		CargoShips:       400,
	})

	result := sim.Run()

	spew.Dump(result)
	logrus.Infof("resources found in %d expeditions: %d", 4000, result.MSE)
}
