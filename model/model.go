package model

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/PieroNarciso/flapie/entity"
)

type model struct {
	bird      entity.Player
	obstacles []entity.Obstacle
}

func NewModel() tea.Model {
	return &model{
		bird:      entity.NewBirdPlayer(),
		obstacles: make([]entity.Obstacle, 0),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "space":
			p := m.bird.Position()
			m.bird.SetY(p.Y-1)
		}
	}
	return m, nil
}

func (m model) View() string {
	return "Hola Mundo!"
}
