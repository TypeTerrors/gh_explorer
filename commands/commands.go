package commands

import (
	"explorer/services"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type OpenEditorMsg struct{ err error }
type GitPulledMessage struct{ err error }
type CloneEditorMsg struct{ err error }

func OpenEditor(repoName string) tea.Cmd {

	userName := services.GetUser()

	c := exec.Command("code", "--new-window", "/Users/"+userName+"/Documents/"+services.GITHUB_USER+"/"+repoName) //nolint:gosec
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return OpenEditorMsg{err}
	})
}
func CloneEditor(repoLink string, repoName string) tea.Cmd {

	userName := services.GetUser()
	exec.Command("gh", "auth", "switch", "--user", services.GITHUB_USER)                                           
	c := exec.Command("git", "clone", repoLink, "/Users/"+userName+"/Documents/"+services.GITHUB_USER+"/"+repoName) 

	return tea.ExecProcess(c, func(err error) tea.Msg {
		return CloneEditorMsg{err}
	})
}
func GitPull(repoName string) tea.Cmd {

	userName := services.GetUser()
	exec.Command("gh", "auth", "switch", "--user", services.GITHUB_USER)
	c := exec.Command("git", "-C", "/Users/"+userName+"/Documents/"+services.GITHUB_USER+"/"+repoName, "pull")

	return tea.ExecProcess(c, func(err error) tea.Msg {
		return GitPulledMessage{err}
	})
}
