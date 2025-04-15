package entity

type LightStatus string

const (
	LightStatusOn     LightStatus = "on"
	LightStatusOff    LightStatus = "off"
	LightStatusBroken LightStatus = "broken"
)

func (l LightStatus) String() string {
	return string(l)
}
