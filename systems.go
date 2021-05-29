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

func (self *movementSystem) Update(dt float32) {
	// self.spaceComponent.Position = engo.Point{100, 100}
	// A friendly reminder that **we do NOT do a little trolling**
	if engo.Input.Button("MoveLeft").Down() {
		self.spaceComponent.Position.X -= 3
	}
	if engo.Input.Button("MoveRight").Down() {
		self.spaceComponent.Position.X += 3
	}
	if engo.Input.Button("Jump").Down() && !self.isJumping && self.totalJump <= 100 {
		self.spaceComponent.Position.Y -= 20
		self.totalJump += 20
		self.isJumping = true
	}
	if self.isJumping && self.totalJump <= 100 && self.totalJump >= -1 {
		self.spaceComponent.Position.Y -= 20
		self.totalJump += 20
	}
	if self.isJumping && self.totalJump >= 100 {
		self.isJumping = false
	}
}

func (self *movementSystem) Add(basicEntity *ecs.BasicEntity, renderComponent *common.RenderComponent, spaceComponent *common.SpaceComponent) {
	self.spaceComponent = spaceComponent
	self.isJumping = false
	self.totalJump = 0
}

func (self *movementSystem) Remove(added ecs.BasicEntity) {
	// nop
}
