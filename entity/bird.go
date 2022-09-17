package entity

type Player interface {
	SetX(x int) Player
	SetY(y int) Player
	X() int
	Y() int
	SetSpeed(speed float64) Player
	AddSpeed(speed float64) Player
	Speed() float64
}

type bird struct {
	position Point
	speed    float64
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

func (b *bird) SetSpeed(speed float64) Player {
	b.speed = speed
	return b
}

func (b *bird) Speed() float64 {
	return b.speed
}

func (b *bird) AddSpeed(speed float64) Player {
	return b
}

func NewBirdPlayer(pos Point, speed float64) Player {
	return &bird{pos, speed}
}
