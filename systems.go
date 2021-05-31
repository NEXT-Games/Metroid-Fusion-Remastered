package main

import (
	"github.com/ByteArena/box2d"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type movementSystem struct {
	spaceComponent *common.SpaceComponent
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
		movementSystem.samus.Body.ApplyLinearImpulseToCenter(box2d.B2Vec2{X: 0, Y: -500}, true)
		movementSystem.samus.isJumping = true
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
