package main

import (
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
}

func (holder *entityHolder) Add(e *entityType) {
	holder.entities = append(holder.entities, e)
}
