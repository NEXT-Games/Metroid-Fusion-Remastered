package main

import (
	"log"

	"github.com/EngoEngine/engo"
	"github.com/Noofbiz/engoBox2dSystem"
)

func addListeners(s *entityHolder) {
	engo.Mailbox.Listen("CollisionStartMessage", func(message engo.Message) {
		c, isCollision := message.(engoBox2dSystem.CollisionStartMessage)
		if isCollision {
			a := c.Contact.GetFixtureA().GetBody().GetUserData()
			b := c.Contact.GetFixtureB().GetBody().GetUserData()
			if c.Contact.IsTouching() {
				for i1, e1 := range s.entities {
					log.Printf("i1: %d", i1)
					if e1.BasicEntity.ID() == a || e1.BasicEntity.ID() == b {
						log.Println("e1 obtained")
						for i2, e2 := range s.entities {
							if i1 == i2 {
								log.Println("i1 == i2")
								continue
							}
							log.Printf("i2: %d", i1)
							if e2.BasicEntity.ID() == a || e2.BasicEntity.ID() == b {
								// This means samus has hit the floor and is no longer jumping.
								// Stop the jump
								log.Println("e1 and e2 obtained")
								if e1.entity.spriteMeta == "samus" {
									log.Println("e1")
									e1.entity.totalJump = 0
								}
								if e2.entity.spriteMeta == "samus" {
									log.Println("e2")
									e2.entity.totalJump = 0
								}
							}
						}
					}
				}
			}
		}
	})
}
