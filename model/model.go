package model

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/PieroNarciso/flapie/constants"
	"github.com/PieroNarciso/flapie/entity"
)


type TickMsg time.Time

type tickRate struct {
	lastTime time.Time
	delta time.Duration
}

func (t *tickRate) UpdateDelta(tm time.Duration) {
	t.delta = tm
}

func (t *tickRate) Delta() time.Duration {
	return t.delta
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
		bird:      entity.NewBirdPlayer(point, constants.INIT_SPEED),
		obstacles: make([]entity.Obstacle, 0),
		height:    0,
		width:     0,
		tickRate:  tickRate{lastTime: time.Now()},
		pressed:   false,
	}
}

func (m model) tick() tea.Cmd {
	return tea.Tick(time.Second/constants.FPS, func(t time.Time) tea.Msg {
		return TickMsg(time.Now())
	})
}

func (m model) Init() tea.Cmd {
	return m.tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case " ":
			m.bird.Jump()
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case TickMsg:
		delta := time.Since(time.Time(msg))
		m.tickRate.UpdateDelta(delta)
		m.bird.Update(delta)
		return m, m.tick()
	}
	return m, nil
}

func (m model) View() string {

	view := strings.Builder{}

	view.WriteString(fmt.Sprintf("Pos: %v  Speed: %v   Delta: %v\n", m.bird.Y(), m.bird.Speed(), m.tickRate.Delta()/100.0))

	for y := m.height - 1; y >= 0; y-- {
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
