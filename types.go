package main

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/Noofbiz/engoBox2dSystem"
)

type BaseEntity struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	engoBox2dSystem.Box2dComponent
	spriteMeta string
	totalJump  int
	canJump    bool
}
type entityType struct {
	ecs.BasicEntity
	*engoBox2dSystem.Box2dComponent
	entity BaseEntity
}

type entityHolder struct {
	entities []*entityType
	msys     *movementSystem
}

type DummyMessage struct{}

type Text struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

func (entity *entityType) Debug() {
	log.Printf("debug entityType totaljump: %d", entity.entity.totalJump)
	if entity.entity.canJump {
		log.Println("debug entityType success canJump")
	}
}
func (holder *entityHolder) Add(e *entityType) {
	holder.entities = append(holder.entities, e)
}
func (holder *entityHolder) SetMsys(sys *movementSystem) {
	holder.msys = sys
}
func (holder *entityType) SetCanJump() {
	holder.entity.canJump = true
	holder.entity.totalJump = 0
}

func (*DummyMessage) Type() string {
	return "menuswitch"
}
