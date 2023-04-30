package simulator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

type Simulator struct {
	r *rand.Rand

	config *Config
}

type SimulationResult struct {
	Reward       Reward
	Factor       int
	ResourceType Resource
	Resources    int
}

func (sr *SimulationResult) GetMSE() int {
	switch sr.ResourceType {
	case RESOURCE_DEUTERIUM:
		return sr.Resources * 3
	case RESOURCE_CRYSTAL:
		return sr.Resources * 2
	default:
		return sr.Resources
	}
}

type AggregateResult struct {
	MSE int

	Results []*SimulationResult
}

func (sr *SimulationResult) String() string {
	return fmt.Sprintf("[%s] roll %d yields %d %s", sr.Reward, sr.Factor, sr.Resources, sr.ResourceType)
}

func NewSimulator(c *Config) (*Simulator, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	return &Simulator{r, c}, nil
}

func (s *Simulator) Run() *AggregateResult {
	aggregate := &AggregateResult{
		Results: make([]*SimulationResult, s.config.Rounds),
	}

	for n := 0; n < s.config.Rounds; n++ {
		res := s.simulate(n)
		aggregate.Results[n] = res
		aggregate.MSE += res.GetMSE()
	}

	return aggregate
}

func (s *Simulator) simulate(round int) *SimulationResult {
	reward := s.rollReward()
	result := &SimulationResult{Reward: reward}

	switch reward {
	case REWARD_RESOURCES:
		result.Factor = s.rollFactor()
		resourcePoints := s.calculateResources(result.Factor)
		result.ResourceType = s.rollResourceType()
		result.Resources = int(result.ResourceType) * resourcePoints

		if cap := s.config.Cargo.Capacity(s.config.HyperspaceTechnology); cap*s.config.CargoShips < result.Resources {
			logrus.Warnf("rolled value %d over ship capacity %d", result.Resources, cap*s.config.CargoShips)
			result.Resources = cap * s.config.CargoShips
		}
	}

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

func (s *Simulator) rollReward() Reward {
	switch roll := s.r.Float64(); {
	case roll < 0.325:
		return REWARD_RESOURCES
	case roll < 0.545:
		return REWARD_SHIPS
	default:
		return REWARD_NONE
	}
}

func (s *Simulator) rollResourceType() Resource {
	switch roll := s.r.Float32(); {
	case roll < 0.685:
		return RESOURCE_METAL
	case roll < 0.685+0.24:
		return RESOURCE_CRYSTAL
	default:
		return RESOURCE_DEUTERIUM
	}
}

func (s *Simulator) calculateResources(factor int) int {
	return factor * s.config.ExpeditionPoints * s.config.SpeedFactor
}

func (s *Simulator) rollFactorValue(min, max, step int) int {
	return min + (s.r.Int()%((max-min)/step))*step
}
