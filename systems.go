package main

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type movementSystem struct {
	spaceComponent *common.SpaceComponent
	samus          BaseEntity
}

func (*movementSystem) Type() string { return "movementSystem" }

func (movementSystem *movementSystem) Update(dt float32) {
	// movementSystem.spaceComponent.Position = engo.Point{100, 100}
	// A friendly reminder that **we do NOT do a little trolling**
	if engo.Input.Button("MoveLeft").Down() {
		movementSystem.spaceComponent.Position.X -= 3
	}
	if engo.Input.Button("MoveRight").Down() {
		movementSystem.spaceComponent.Position.X += 3
	}
	if engo.Input.Button("Jump").Down() && movementSystem.samus.totalJump < 100 {
		movementSystem.spaceComponent.Position.Y -= 20
		movementSystem.samus.totalJump += 20
	}
	if engo.Input.Button("Jump").Down() && !movementSystem.samus.canJump {
		log.Printf("%d", movementSystem.samus.totalJump)
	}
}
func (movementSystem *movementSystem) AddEtc(samus BaseEntity) {
	movementSystem.samus = samus
	movementSystem.samus.canJump = true
}
func (movementSystem *movementSystem) Add(basicEntity *ecs.BasicEntity, renderComponent *common.RenderComponent, spaceComponent *common.SpaceComponent) {
	movementSystem.spaceComponent = spaceComponent
}

func (movementSystem *movementSystem) Remove(added ecs.BasicEntity) {
	// nop
}

type menuSystem struct{}

func (*menuSystem) Type() string { return "menuSystem" }
func (sys menuSystem) Update(dt float32) {
	if engo.Input.Button("startgame").JustPressed() {
		engo.Mailbox.Dispatch(&DummyMessage{})
		engo.SetScene(&MainDeckScene{}, true)
	}
}
func (sys menuSystem) Add(e *ecs.BasicEntity) {
	// nop
}

func (sys menuSystem) Remove(e ecs.BasicEntity) {
	// nop
}
