package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/Noofbiz/engoBox2dSystem"
)

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

func addListeners(s *entityHolder) {
	engo.Mailbox.Listen("CollisionStartMessage", func(message engo.Message) {
		c, isCollision := message.(engoBox2dSystem.CollisionStartMessage)
		if isCollision {
			a := c.Contact.GetFixtureA().GetBody().GetUserData()
			b := c.Contact.GetFixtureB().GetBody().GetUserData()
			if c.Contact.IsTouching() {
				for i1, e1 := range s.entities {
					if e1.BasicEntity.ID() == a || e1.BasicEntity.ID() == b {
						for i2, e2 := range s.entities {
							if i1 == i2 {
								continue
							}
							if e2.BasicEntity.ID() == a || e2.BasicEntity.ID() == b {
								// This means samus has hit the floor and is no longer jumping.
								// Stop the jump
								if e1.entity.objMeta == "samus" {
									e1.entity.isJumping = false
								}
							}
						}
					}
				}
			}
		}
	})
}
