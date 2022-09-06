package model

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/PieroNarciso/flapie/entity"
)

const (
	ACCEL = 2
	SPEED = 20
)

type TickMsg time.Time

type model struct {
	bird      entity.Player
	obstacles []entity.Obstacle
	height    int
	width     int
	tickRate  time.Time
}

func NewModel() tea.Model {
	point := entity.Point{X: 2, Y: 20}
	return &model{
		bird:      entity.NewBirdPlayer(point, SPEED, ACCEL),
		obstacles: make([]entity.Obstacle, 0),
		height:    0,
		width:     0,
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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case " ":
			m.bird.SetY(m.bird.Y() - 1)
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case TickMsg:
		m.tickRate = time.Now()
		return m, m.tick()
	}
	return m, nil
}

func (m model) View() string {
	view := strings.Builder{}

	for y := 0; y < m.height; y++ {
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
