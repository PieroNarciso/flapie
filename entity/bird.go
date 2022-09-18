package entity

import (
	"time"

	"github.com/PieroNarciso/flapie/constants"
)

type Player interface {
	SetX(x int) Player
	SetY(y int) Player
	X() int
	Y() int
	Speed() float64
	Jump() Player
	Update(t time.Duration) Player
}

type bird struct {
	pos   Point
	speed float64
}

func (b *bird) Jump() Player {
	b.SetSpeed(constants.JUMP_SPEED)
	return b
}

func (b *bird) SetX(x int) Player {
	b.pos.X = x
	return b
}

func (b *bird) SetY(y int) Player {
	b.pos.Y = y
	return b
}

func (b *bird) X() int {
	return b.pos.X
}

func (b *bird) Y() int {
	return b.pos.Y
}

func (b *bird) SetSpeed(speed float64) Player {
	b.speed = speed
	return b
}

func (b *bird) Speed() float64 {
	return b.speed
}

func (b *bird) Update(t time.Duration) Player {
	delta := float64(t.Microseconds()) / 100.0
	b.pos.Y += int(b.Speed() * delta)
	b.speed -= constants.FALL_SPEED * delta
	return b
}

func NewBirdPlayer(pos Point, speed float64) Player {
	return &bird{pos, speed}
}
