package simulator

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

type Simulator struct {
	r      *rand.Rand
	config *Config
}

func NewSimulator(c *Config, seed int64) (*Simulator, error) {
	r := rand.New(rand.NewSource(time.Now().Unix() + seed))

	return &Simulator{r, c}, nil
}

func (s *Simulator) Run(ships []*Ship) *AggregateResult {
	aggregate := &AggregateResult{
		Results: make([]*SimulationResult, s.config.Rounds),
		Rewards: map[Reward]int{
			REWARD_NONE:      0,
			REWARD_RESOURCES: 0,
			REWARD_SHIPS:     0,
			REWARD_BLACKHOLE: 0,
		},
	}

	for n := 0; n < s.config.Rounds; n++ {
		res := s.simulate(ships)
		aggregate.Results[n] = res
		aggregate.MSE += res.GetMSE()
		aggregate.FuelConsumption += res.FuelConsumption
		aggregate.Rewards[res.Reward] += 1
	}

	return aggregate
}

func (s *Simulator) simulate(ships []*Ship) *SimulationResult {
	reward := s.rollReward()
	result := &SimulationResult{
		Reward:          reward,
		FuelConsumption: CalculateConsumption(ships, 1040, 10, 1),
	}

	switch reward {
	case REWARD_RESOURCES:
		result.Factor = s.rollFactor()
		resourcePoints := s.calculateResources(result.Factor)
		result.ResourceType = s.rollResourceType()

		switch result.ResourceType {
		case RESOURCE_METAL:
			result.Resources = resourcePoints
		case RESOURCE_CRYSTAL:
			result.Resources = resourcePoints / 2
		case RESOURCE_DEUTERIUM:
			result.Resources = resourcePoints / 3
		}

		storageCapacity := CalculateCapacity(ships)

		if storageCapacity < result.Resources {
			logrus.Tracef("rolled value %d over ship capacity %d", result.Resources, storageCapacity)
			result.Resources = storageCapacity
		}
	case REWARD_SHIPS:
		result.Factor = s.rollFactor()
		structuralIntegrity := float64(s.calculateStructuralIntegrity(result.Factor))
		result.Resources = int(structuralIntegrity*0.54 +
			structuralIntegrity*0.46*1.5 +
			structuralIntegrity*0.1*3.)
	case REWARD_BLACKHOLE:
		result.Resources = ships[0].Count * -150000
	}

	return result
}

func (s *Simulator) rollFactor() int {
	switch roll := s.r.Float32(); {
	case roll < 0.01:
		return s.rollFactorValue(100, 200, 2)
	case roll < 0.11:
		// large
		return s.rollFactorValue(50, 100, 2)
	default:
		// normal
		return s.rollFactorValue(10, 50, 2)
	}
}

func (s *Simulator) rollReward() Reward {
	switch roll := s.r.Float64(); {
	case roll < 0.325:
		return REWARD_RESOURCES
	case roll < 0.325+0.22:
		return REWARD_SHIPS
	case roll >= 1.-.0033:
		return REWARD_BLACKHOLE
	default:
		return REWARD_NONE
	}
}

func (s *Simulator) rollResourceType() Resource {
	switch roll := s.r.Float32(); {
	case roll < 0.50:
		return RESOURCE_METAL
	case roll < 0.50+0.33:
		return RESOURCE_CRYSTAL
	default:
		return RESOURCE_DEUTERIUM
	}
}

func (s *Simulator) calculateResources(factor int) int {
	return int(float64(factor*s.config.ExpeditionPoints*s.config.SpeedFactor*2) * 1.5)
}

func (s *Simulator) calculateStructuralIntegrity(factor int) int {
	return ((factor * s.config.ExpeditionPoints) / 2) * s.config.SpeedFactor
}

func (s *Simulator) rollFactorValue(min, max, step int) int {
	return min + (s.r.Int()%((max-min)/step))*step
}
