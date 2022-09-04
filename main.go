package main

import (

	"github.com/PieroNarciso/flapie/model"
	tea "github.com/charmbracelet/bubbletea"
)


func main() {
	p := tea.NewProgram(model.NewModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		panic(err.Error())
	}
}
