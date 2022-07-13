package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	menu Menu
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			m.menu, cmd = m.menu.Update(msg)
			cmds = append(cmds, cmd)
		}
	}
	return m, tea.Batch(cmds...)
}
func (m model) View() string {
	// The root model is in charge of rendering two components: the menu, and
	// the main view screen that reflects the corresponding items in the menu
	return m.menu.View()
}

func newModel() model {
	m := NewMenu()
	return model{menu: m}

}

func main() {

	m := newModel()

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
