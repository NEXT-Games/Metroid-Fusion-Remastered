package main

import (
	"log"

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
	if engo.Input.Button("Jump").Down() && movementSystem.samus.canJump {
		movementSystem.samus.Body.ApplyLinearImpulseToCenter(box2d.B2Vec2{X: 0, Y: -1000}, true)
		movementSystem.samus.totalJump += 10
	}
	if movementSystem.samus.totalJump >= 100 && movementSystem.samus.canJump {
		log.Println("*unjumps your samus*")
		movementSystem.samus.canJump = false
	}
}
func (movementSystem *movementSystem) AddEtc(samus *BaseEntity) {
	movementSystem.samus = samus
	movementSystem.samus.canJump = true
}
func (movementSystem *movementSystem) Add(basicEntity *ecs.BasicEntity, renderComponent *common.RenderComponent, spaceComponent *common.SpaceComponent) {
	movementSystem.spaceComponent = spaceComponent
}

func (movementSystem *movementSystem) Remove(added ecs.BasicEntity) {
	// nop
}
