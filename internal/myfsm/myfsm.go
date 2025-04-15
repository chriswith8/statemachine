package myfsm

import (
	"context"
)

type (
	eKey struct {
		name string
		src  string
	}
	Event struct {
		Name      string
		Src       string
		Dst       string
		Condition func() bool
		OnEnter   func(ctx context.Context)
	}
	Events []Event
	MyFSM  struct {
		current     string
		transitions map[eKey]Event
		setter      func(state string)
	}
)

func NewMyFSM(initial string, setter func(state string), events Events) *MyFSM {
	setter(initial)

	f := &MyFSM{
		current:     initial,
		setter:      setter,
		transitions: make(map[eKey]Event),
	}

	for _, event := range events {
		f.transitions[eKey{name: event.Name, src: event.Src}] = event
	}
	return f
}

func (f *MyFSM) doTransition(newState string) {
	f.setter(newState)
	f.current = newState
}

func (f *MyFSM) SendEvent(ctx context.Context, name string) error {
	e, ok := f.transitions[eKey{name: name, src: f.current}]
	if !ok {
		return NewMachineError(ErrImpossibleTransition)
	}

	if e.Condition != nil && !e.Condition() {
		return NewMachineError(ErrConditionIsFalse)
	}

	f.doTransition(e.Dst)

	if e.OnEnter != nil {
		e.OnEnter(ctx)
	}

	return nil
}
