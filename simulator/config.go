package simulator

type Config struct {
	Rounds               int
	ExpeditionPoints     int
	SpeedFactor          int
	HyperspaceTechnology int

	Cargo      Cargo
	CargoShips int
	Reward     Reward
}
