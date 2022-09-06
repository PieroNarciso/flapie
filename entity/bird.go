package entity

type Player interface {
	SetX(x int) Player
	SetY(y int) Player
	X() int
	Y() int
	SetSpeed(speed int) Player
	SetAccel(accel int) Player
	Speed() int
	Accel() int
}

type bird struct {
	position Point
	speed    int
	accel    int
}

func (b *bird) SetX(x int) Player {
	b.position.X = x
	return b
}

func (b *bird) SetY(y int) Player {
	b.position.Y = y
	return b
}

func (b *bird) X() int {
	return b.position.X
}

func (b *bird) Y() int {
	return b.position.Y
}

func (b *bird) SetSpeed(speed int) Player {
	b.speed = speed
	return b
}

func (b *bird) Speed() int {
	return b.speed
}

func (b *bird) SetAccel(accel int) Player {
	b.accel = accel
	return b
}

func (b *bird) Accel() int {
	return b.accel
}

func NewBirdPlayer(pos Point, speed, accel int) Player {
	return &bird{pos, speed, accel}
}
