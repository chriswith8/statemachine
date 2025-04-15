package entity

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

type Light struct {
	Status LightStatus `json:"status"`
	FSM    *fsm.FSM    `json:"fsm"`
}

func NewLight() *Light {
	light := &Light{Status: LightStatusOff}
	light.FSM = fsm.NewFSM(
		light.Status.String(),
		fsm.Events{
			{Name: "turn_on", Src: []string{LightStatusOff.String()}, Dst: LightStatusOn.String()},
			{Name: "turn_off", Src: []string{LightStatusOn.String()}, Dst: LightStatusOff.String()},
			{Name: "hit_someone", Src: []string{LightStatusOn.String(), LightStatusOff.String()}, Dst: LightStatusBroken.String()},
		},
		fsm.Callbacks{
			"enter_state": func(_ context.Context, e *fsm.Event) { fmt.Printf("Event %s was sent\n", e.Event) },
		},
	)

	return light
}

func (l *Light) TurnOn(ctx context.Context) error {
	err := l.FSM.Event(ctx, "turn_on")
	if err == nil {
		l.Status = LightStatusOn

	}
	return err

}
func (l *Light) TurnOff(ctx context.Context) error {
	err := l.FSM.Event(ctx, "turn_off")
	if err == nil {
		l.Status = LightStatusOff
	}
	return err
}
func (l *Light) HitSomeone(ctx context.Context) error {
	err := l.FSM.Event(ctx, "hit_someone")
	if err == nil {
		l.Status = LightStatusBroken
	}
	return err
}
