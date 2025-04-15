package entity

type PersonState string

func (p PersonState) String() string {
	return string(p)
}

const (
	PersonStateSitting  PersonState = "sitting"
	PersonStateStanding PersonState = "standing"
	PersonStateWalking  PersonState = "walking"
)
