package simulator

import "fmt"

type SimulationResult struct {
	Reward          Reward
	Factor          int
	ResourceType    Resource
	Resources       int
	FuelConsumption int
}

func (sr *SimulationResult) GetMSE() int {
	switch sr.ResourceType {
	case RESOURCE_DEUTERIUM:
		return sr.Resources * 3
	case RESOURCE_CRYSTAL:
		return int(float64(sr.Resources) * 1.5)
	default:
		return sr.Resources
	}
}

func (sr *SimulationResult) String() string {
	return fmt.Sprintf("[%s] roll %d yields %d %s", sr.Reward, sr.Factor, sr.Resources, sr.ResourceType)
}

type AggregateResult struct {
	MSE             int
	FuelConsumption int
	ShipCount       int

	Results []*SimulationResult
	Rewards map[Reward]int
}

func (ar *AggregateResult) Gain() int {
	return ar.MSE - ar.FuelConsumptionMSE()
}

func (ar *AggregateResult) FuelConsumptionMSE() int {
	return ar.FuelConsumption * 3
}
