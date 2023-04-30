package simulator

import (
	"math"

	"github.com/sirupsen/logrus"
)

// set the distance to a fixed distance from pos 8
const DISTANCE = 1040
const SPEED_PERCENTAGE float64 = 10

func shipSpeedFactor(distance, shipSpeed int) float64 {
	result := math.Sqrt(float64(distance*10) / float64(shipSpeed))
	logrus.Infof("shipSpeedFactor: %f", result)
	return result
}

func duration(distance, maxSpeed int) float64 {
	result := (35000/float64(distance))*shipSpeedFactor(distance, maxSpeed) + 10
	logrus.Infof("duration: %f", result)
	return result
}

func speed() float64 {
	return (35000 / (duration(DISTANCE, 16500) - 10)) * shipSpeedFactor(DISTANCE, 16500)
}

func Consumption(baseConsumption, count int) float64 {
	return (float64((baseConsumption * count * DISTANCE)) / 35000.) * math.Pow(speed()/10+1, 2.)
}
