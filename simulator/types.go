package simulator

type Cargo int

func (c Cargo) Capacity(ht int) int {
	switch c {
	case CARGO_SC:
		return int(5000 * (1. + float64(ht)*0.05))
	case CARGO_LC:
		return int(25000 * (1. + float64(ht)*0.05))
	default:
		return 0
	}
}

const (
	CARGO_NONE Cargo = iota
	CARGO_SC
	CARGO_LC
)

type Reward int

func (r Reward) String() string {
	switch r {
	case REWARD_NONE:
		return "REWARD_NONE"
	case REWARD_RESOURCES:
		return "REWARD_RESOURCES"
	case REWARD_SHIPS:
		return "REWARD_SHIPS"
	case REWARD_BLACKHOLE:
		return "REWARD_BLACKHOLE"
	default:
		return "REWARD_UNKNOWN"
	}
}

const (
	REWARD_NONE Reward = iota
	REWARD_RESOURCES
	REWARD_SHIPS
	REWARD_BLACKHOLE
)

type Bucket int

const (
	BUCKET_NORMAL Bucket = iota
	BUCKET_LARGE
	BUCKET_XLARGE
)

type Resource int

func (r Resource) String() string {
	switch r {
	case RESOURCE_DEUTERIUM:
		return "RESOURCE_DEUTERIUM"
	case RESOURCE_CRYSTAL:
		return "RESOURCE_CRYSTAL"
	case RESOURCE_METAL:
		return "RESOURCE_METAL"
	default:
		return "RESOURCE_NONE"
	}
}

const (
	RESOURCE_NONE Resource = iota
	RESOURCE_DEUTERIUM
	RESOURCE_CRYSTAL
	RESOURCE_METAL
)
