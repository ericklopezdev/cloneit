package ui

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"cloneit/internal/types"
)

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			if i, ok := m.list.SelectedItem().(types.RepoItem); ok {
				// Clone
				repo := i.GetRepo()
				cmd := exec.Command("git", "clone", repo.SshUrl)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Fprintf(os.Stderr, "Error cloning: %v\n", err)
					os.Exit(1)
				}
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		h, v := lipgloss.NewStyle().GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}

// NewModel creates a new model with the given repos
func NewModel(repos []types.Repo) model {
	var items []list.Item
	for _, repo := range repos {
		items = append(items, types.NewRepoItem(repo))
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select a repository to clone"

	return model{list: l}
}

func Run(m model) error {
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	return err
}
