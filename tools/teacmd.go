package tools

import (
	tea "github.com/charmbracelet/bubbletea"
)

func AnyCmd[T any](msg T) tea.Cmd {
	return tea.Cmd(func() tea.Msg {
		return msg
	})
}
