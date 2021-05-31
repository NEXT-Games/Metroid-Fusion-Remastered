package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type movementSystem struct {
	spaceComponent *common.SpaceComponent
	totalJump      int
	samus          *BaseEntity
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
	if engo.Input.Button("Jump").Down() && !movementSystem.samus.isJumping {
		movementSystem.spaceComponent.Position.Y -= 20
		movementSystem.totalJump += 20
		movementSystem.samus.isJumping = true
	}
	if movementSystem.samus.isJumping && movementSystem.totalJump <= 100 && movementSystem.totalJump >= -1 {
		movementSystem.spaceComponent.Position.Y -= 20
		movementSystem.totalJump += 20
	}
}
func (movementSystem *movementSystem) AddEtc(samus *BaseEntity) {
	movementSystem.samus = samus
}
func (movementSystem *movementSystem) Add(basicEntity *ecs.BasicEntity, renderComponent *common.RenderComponent, spaceComponent *common.SpaceComponent) {
	movementSystem.spaceComponent = spaceComponent
}

func (movementSystem *movementSystem) Remove(added ecs.BasicEntity) {
	// nop
}
