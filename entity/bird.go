package entity

type Player interface {
	SetX(x int) Player
	SetY(y int) Player
	Position() Point
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
	b.position.x = x
	return b
}

func (b *bird) SetY(y int) Player {
	b.position.y = y
	return b
}

func (b *bird) SetSpeed(speed int) Player {
	b.speed = speed
	return b
}

func (b *bird) Position() Point {
	return b.position
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

func NewBirdPlayer() Player {
	return &bird{}
}
