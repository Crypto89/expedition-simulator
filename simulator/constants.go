package simulator

type Cargo int

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
	default:
		return "REWARD_UNKNOWN"
	}
}

const (
	REWARD_NONE Reward = iota
	REWARD_RESOURCES
	REWARD_SHIPS
)

type Bucket int

const (
	BUCKET_NORMAL Bucket = iota
	BUCKET_LARGE
	BUCKET_XLARGE
)
