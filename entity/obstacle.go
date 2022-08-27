package entity


type Obstacle interface {
	Location() int
	SetLocation(l int) Obstacle
}

type barObstacle struct {
	// Column location
	location int
}

func (b *barObstacle) Location() int {
	return b.location
}

func (b *barObstacle) SetLocation(l int) Obstacle {
	b.location = l
	return b
}

func NewBarObstacle() Obstacle {
	return &barObstacle{}
}
