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
		bird: entity.NewBirdPlayer(),
		obstacles: make([]entity.Obstacle, 0),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m model) View() string {
	return ""
}
