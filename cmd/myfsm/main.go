package main

import (
	"context"
	"fmt"

	"github.com/chriswith8/statemachine/internal/entity"
)

func main() {
	ctx := context.Background()
	p := entity.NewPerson("Chris")
	if err := p.Walk(ctx); err != nil {
		panic(err)
	}

	p.PrintState()

	if err := p.Sit(ctx); err != nil {
		fmt.Println("Person must be standing to sit")
	}

	if err := p.Stop(ctx); err != nil {
		fmt.Println("Person must be walking to stop")
	}

	p.PrintState()
}
