package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type movementSystem struct {
	spaceComponent *common.SpaceComponent
	totalJump      int
	isJumping      bool
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
	if engo.Input.Button("Jump").Down() && !movementSystem.isJumping && movementSystem.totalJump <= 100 {
		movementSystem.spaceComponent.Position.Y -= 20
		movementSystem.totalJump += 20
		movementSystem.isJumping = true
	}
	if movementSystem.isJumping && movementSystem.totalJump <= 100 && movementSystem.totalJump >= -1 {
		movementSystem.spaceComponent.Position.Y -= 20
		movementSystem.totalJump += 20
	}
	if movementSystem.isJumping && movementSystem.totalJump >= 100 {
		movementSystem.isJumping = false
	}
}

func (movementSystem *movementSystem) Add(basicEntity *ecs.BasicEntity, renderComponent *common.RenderComponent, spaceComponent *common.SpaceComponent) {
	movementSystem.spaceComponent = spaceComponent
	movementSystem.isJumping = false
	movementSystem.totalJump = 0
}

func (movementSystem *movementSystem) Remove(added ecs.BasicEntity) {
	// nop
}
