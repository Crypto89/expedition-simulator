package main

import (
	"math"
	"sync"
	"time"

	"github.com/crypto89/expedition-simulator/simulator"
	"github.com/sirupsen/logrus"
)

func createShips(cargoCount int) []*simulator.Ship {
	return []*simulator.Ship{
		// large cargo
		{
			Count:           cargoCount,
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
	logrus.SetLevel(logrus.DebugLevel)
	// spew.Dump(simulator.CalculateConsumption(ships, 1040, 10, 1))
	config := &simulator.Config{
		Rounds:           1000000,
		ExpeditionPoints: 12000,
		SpeedFactor:      8,
	}

	start := time.Now()

	var results []*simulator.AggregateResult
	var mu sync.Mutex
	var thread int64
	wg := &sync.WaitGroup{}
	// simulationCount := int64(runtime.GOMAXPROCS(0) * 2)
	simulationCount := int64(1)
	logrus.Infof("simulating %d runs with %d expeditions", simulationCount, config.Rounds)
	for thread = 0; thread < simulationCount; thread++ {
		wg.Add(1)
		go func(thread int64) {
			defer wg.Done()
			result := simulate(thread, config)
			mu.Lock()
			defer mu.Unlock()
			results = append(results, result)
		}(thread)
	}

	wg.Wait()

	largeCargo := 0
	for _, result := range results {
		largeCargo += result.ShipCount
	}

	logrus.Infof("simulation took: %s", time.Since(start))
	logrus.Infof("best expedition with %f Large Cargos", math.Ceil(float64(largeCargo)/float64(simulationCount)))
}

func simulate(thread int64, config *simulator.Config) *simulator.AggregateResult {
	var best *simulator.AggregateResult

	sim, _ := simulator.NewSimulator(config, thread)

	for count := 500; count <= 1440; count += 10 {
		ships := createShips(count)
		result := sim.Run(ships)
		if best == nil || result.Gain() > best.Gain() {
			best = result
			best.ShipCount = count
		}
		logrus.Debugf("expedition with %d large cargo gained %f on average", count, float64(result.Gain())/float64(config.Rounds))
	}

	// logrus.Infof("best rolls: %v", best.Rewards)

	return best
}
