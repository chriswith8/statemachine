package main

import (
	"context"
	"fmt"

	"github.com/chriswith8/statemachine/internal/entity"
	"github.com/looplab/fsm"
)

func main() {
	light := entity.NewLight()
	err := light.TurnOn(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Light is on state: %s\n", light.Status)

	err = light.TurnOn(context.Background())
	if err != nil && err.Error() == "event turn_on inappropriate in current state on" {
		fmt.Println("Light is already on")
	}

	err = light.TurnOff(context.Background())

	view, err := fsm.VisualizeWithType(light.FSM, fsm.MermaidFlowChart)
	if err != nil {
		panic(err)
	}

	fmt.Println(view)
}
