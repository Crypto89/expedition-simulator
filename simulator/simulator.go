package simulator

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Rounds int

	Cargo  Cargo
	Reward Reward
}

type Simulator struct {
	r *rand.Rand

	config *Config
}

type SimulationResult struct {
}

func NewSimulator(c *Config) (*Simulator, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	return &Simulator{r, c}, nil
}

func (s *Simulator) Run() []*SimulationResult {
	factor := s.rollFactor()

	for n := 0; n < s.config.Rounds; n++ {
		s.simulate(n)
		logrus.Infof("factor: %d", factor)
	}

	return nil
}

func (s *Simulator) simulate(round int) *SimulationResult {
	result := &SimulationResult{}

	return result
}

func (s *Simulator) rollFactor() int {
	switch roll := s.r.Float32(); {
	case roll < 0.01:
		return s.rollFactorValue(100, 200, 2)
	case roll < 0.11:
		return s.rollFactorValue(50, 100, 2)
	default:
		return s.rollFactorValue(10, 50, 2)
	}
}

func (s *Simulator) rollFactorValue(min, max, step int) int {
	return min + (s.r.Int()%((max-min)/step))*step
}
