package model

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/PieroNarciso/flapie/entity"
)

const (
	FALL_SPEED = 40
	JUMP_SPEED = 23
	INIT_SPEED = 0
)

type TickMsg time.Time

type tickRate struct {
	lastTime time.Time
}

func (t *tickRate) UpdateLastTime() {
	t.lastTime = time.Now()
}

func (t *tickRate) Delta() float64 {
	delta := time.Since(t.lastTime).Seconds()
	return delta
}

type model struct {
	bird      entity.Player
	obstacles []entity.Obstacle
	height    int
	width     int
	tickRate  tickRate
	pressed   bool
}

func NewModel() tea.Model {
	point := entity.Point{X: 2, Y: 20}
	return &model{
		bird:      entity.NewBirdPlayer(point, INIT_SPEED),
		obstacles: make([]entity.Obstacle, 0),
		height:    0,
		width:     0,
		tickRate:  tickRate{lastTime: time.Now()},
		pressed:   false,
	}
}

func (m model) tick() tea.Cmd {
	return tea.Tick(time.Second/10, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return m.tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.pressed == true {
		m.bird.SetSpeed(JUMP_SPEED)
	}
	m.pressed = false
	newY := m.bird.Y() + int(m.bird.Speed()*m.tickRate.Delta())
	m.bird.SetY(newY)
	newSpeed := m.bird.Speed() - FALL_SPEED*m.tickRate.Delta()
	m.bird.SetSpeed(newSpeed)
	m.tickRate.UpdateLastTime()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case " ":
			m.pressed = true
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case TickMsg:
		return m, m.tick()
	}
	return m, nil
}

func (m model) View() string {

	view := strings.Builder{}

	// view.WriteString(fmt.Sprintf("Pos: %v  Speed: %v   Delta: %v\n", m.bird.Y(), m.bird.Speed(), m.tickRate.Delta()))

	for y := m.height-1; y >= 0; y-- {
		for x := 0; x < m.width; x++ {
			if y == m.bird.Y() && x == m.bird.X() {
				view.WriteRune('*')
			} else {
				view.WriteRune(' ')
			}
		}
		view.WriteRune('\n')
	}

	return view.String()
}
