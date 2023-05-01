package main

import (
	"time"

	"github.com/crypto89/expedition-simulator/simulator"
	"github.com/sirupsen/logrus"
)

func createShips(cargoCount int) []*simulator.Ship {
	return []*simulator.Ship{
		// large cargo
		{
			Count:           937,
			BaseConsumption: 50,
			Speed:           16500,
			Capacity:        37500,
		},
		// pathfinder
		{
			Count:           1,
			BaseConsumption: 300,
			Speed:           40800,
			Capacity:        15000,
		},
		// reaper
		{
			Count:           1,
			BaseConsumption: 1100,
			Speed:           23800,
			Capacity:        15000,
		},
		// espionage probe
		{
			Count:           1,
			BaseConsumption: 1,
			Speed:           220000000,
			Capacity:        0,
		},
	}
}

func main() {
	ships := createShips(900)

	// spew.Dump(simulator.CalculateConsumption(ships, 1040, 10, 1))
	config := &simulator.Config{
		Rounds:           1000000,
		ExpeditionPoints: 12000,
		SpeedFactor:      8,
	}

	sim, _ := simulator.NewSimulator(config, ships)

	start := time.Now()
	result := sim.Run()
	logrus.Infof("simulation took: %s", time.Since(start))

	logrus.Infof("resources found in %d expeditions: %d (fuel: %d)", 1000000, result.MSE, result.FuelConsumption)
	logrus.Infof("average per expedition: %d", result.MSE/config.Rounds)
}
