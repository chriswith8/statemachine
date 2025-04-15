package entity

import (
	"context"
	"fmt"

	"github.com/chriswith8/statemachine/internal/myfsm"
)

type Person struct {
	Name  string
	State PersonState
	fsm   *myfsm.MyFSM
}

func NewPerson(name string) *Person {
	p := &Person{Name: name}
	p.fsm = myfsm.NewMyFSM(
		PersonStateStanding.String(),
		p.setState,
		myfsm.Events{
			{Name: "sit", Src: PersonStateStanding.String(), Dst: PersonStateSitting.String()},
			{Name: "stand", Src: PersonStateSitting.String(), Dst: PersonStateStanding.String()},
			{Name: "walk", Src: PersonStateStanding.String(), Dst: PersonStateWalking.String()},
			{Name: "stop", Src: PersonStateWalking.String(), Dst: PersonStateStanding.String(), Condition: p.IsWalking},
		},
	)

	return p
}

func (p *Person) setState(state string) {
	p.State = PersonState(state)
}

func (p *Person) Sit(ctx context.Context) error {
	return p.fsm.SendEvent(ctx, "sit")
}

func (p *Person) Walk(ctx context.Context) error {
	return p.fsm.SendEvent(ctx, "walk")
}

func (p *Person) Stand(ctx context.Context) error {
	return p.fsm.SendEvent(ctx, "stand")
}

func (p *Person) Stop(ctx context.Context) error {
	return p.fsm.SendEvent(ctx, "stop")
}

func (p Person) PrintState() {
	fmt.Printf("Person %s is now %s\n", p.Name, p.State)
}

func (p Person) IsWalking() bool {
	return p.State == PersonStateWalking
}
