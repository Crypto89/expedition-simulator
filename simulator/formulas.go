package simulator

import (
	"math"
)

type Ship struct {
	Count           int
	BaseConsumption int
	Speed           int
	Capacity        int
}

func CalculateCapacity(ships []*Ship) int {
	result := 0

	for _, ship := range ships {
		result += ship.Count * ship.Capacity
	}

	return result
}

func CalculateConsumption(ships []*Ship, distance, speedPercent, holdingTime int) int {
	var consumption float64
	var holdingCosts float64
	maxSpeed := getMaxSpeed(ships)
	duration := calculateDuration(distance, maxSpeed, speedPercent)
	speedValue := math.Max(0.5, float64(duration*getFleetSpeedFactor()-10.))

	for _, ship := range ships {
		if ship.Count > 0 {
			shipSpeedValue := 35000. / speedValue * math.Sqrt(float64(distance*10)/float64(ship.Speed))
			holdingCosts += float64(ship.BaseConsumption * ship.Count * holdingTime)
			consumption += math.Max(float64(ship.BaseConsumption*ship.Count*distance)/35000.*math.Pow(shipSpeedValue/10.+1., 2.), 1)
		}
	}

	consumption = math.Round(consumption)
	if holdingTime > 0 {
		consumption += math.Max(math.Floor(holdingCosts/10.), 1)
	}

	return int(consumption)
}

func calculateDuration(distance, maxSpeed, speedPercent int) int {
	return int(math.Max(
		math.Round((35000./float64(speedPercent)*math.Sqrt(float64(distance*10)/float64(maxSpeed))+10.)/float64(getFleetSpeedFactor())),
		1.))
}

func getMaxSpeed(ships []*Ship) int {
	result := math.MaxInt

	for _, ship := range ships {
		if ship.Speed < result {
			result = ship.Speed
		}
	}

	return result
}

func getFleetSpeedFactor() int {
	return 8
}
