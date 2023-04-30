package main

import (
	"math/rand"
	"time"

	"github.com/crypto89/expedition-simulator/simulator"
	"github.com/sirupsen/logrus"
)

const (
	TYPE_NONE int = iota
	TYPE_RESOURCES
	TYPE_SHIPS
)

const (
	BUCKET_NORMAL int = iota
	BUCKET_LARGE
	BUCKET_XLARGE
)

var (
	s                rand.Source
	expeditionPoints = 9000
	speedFactor      = 8
)

func main() {
	sim, _ := simulator.NewSimulator(&simulator.Config{Rounds: 4000})

	sim.Run()
}

func main2() {
	s = rand.NewSource(time.Now().Unix())

	now := time.Now()

	for i := 0; i < 4000; i++ {
		simulate(i)
	}

	logrus.Infof("time: %s", time.Since(now))
}

func simulate(i int) {
	etype := rollType()

	switch etype {
	case TYPE_RESOURCES:
		logrus.WithField("round", i).Infof("roll type: %d", etype)
		simulateResources()
	}
}

func rollType() int {
	r := rand.New(s)

	switch roll := r.Float64(); {
	case roll < 0.325:
		return TYPE_RESOURCES
	case roll < 0.545:
		return TYPE_SHIPS
	default:
		return TYPE_NONE
	}
}

func rollBucket() (int, int) {
	r := rand.New(s)

	switch roll := r.Float32(); {
	case roll < 0.01:
		return BUCKET_XLARGE, rollFactor(100, 200, 2)
	case roll < 0.11:
		return BUCKET_LARGE, rollFactor(50, 100, 2)
	default:
		return BUCKET_NORMAL, rollFactor(10, 50, 2)
	}
}

func rollFactor(min, max, step int) int {
	r := rand.New(s)

	return min + (r.Int()%((max-min)/step))*step
}

func simulateResources() {
	resBucket, factor := rollBucket()
	resources := factor * expeditionPoints * speedFactor

	logrus.Infof("rolled bucket %d, with factor %d yields %d", resBucket, factor, resources)
}
