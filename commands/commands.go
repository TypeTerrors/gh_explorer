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
	var repoPath string
	// if user is specified with an organization than the user takes precedence
	if services.GITHUB_ORG != "" && services.GITHUB_USER != "" {
		repoPath = "/Users/" + userName + "/Documents/" + services.GITHUB_USER + "/" + repoName
		// If an organization is specified without a user, use the organization name
	} else if services.GITHUB_ORG == "" && services.GITHUB_USER != "" {
		repoPath = "/Users/" + userName + "/Documents/" + services.GITHUB_USER + "/" + repoName
	} else if services.GITHUB_ORG != "" && services.GITHUB_USER == "" {
		repoPath = "/Users/" + userName + "/Documents/" + services.GITHUB_ORG + "/" + repoName
	}

	c := exec.Command("code", "--new-window", repoPath) //nolint:gosec
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return OpenEditorMsg{err}
	})
}
func CloneEditor(repoLink string, repoName string) tea.Cmd {

	userName := services.GetUser()
	var repoPath string

	// if user is specified with an organization than the user takes precedence
	if services.GITHUB_ORG != "" && services.GITHUB_USER != "" {
		exec.Command("gh", "auth", "switch", "--user", services.GITHUB_USER)
		repoPath = "/Users/" + userName + "/Documents/" + services.GITHUB_USER + "/" + repoName
		// If an organization is specified without a user, use the organization name
	} else if services.GITHUB_ORG == "" && services.GITHUB_USER != "" {
		repoPath = "/Users/" + userName + "/Documents/" + services.GITHUB_USER + "/" + repoName
		exec.Command("gh", "auth", "switch", "--user", services.GITHUB_USER)
	} else if services.GITHUB_ORG != "" && services.GITHUB_USER == "" {
		repoPath = "/Users/" + userName + "/Documents/" + services.GITHUB_ORG + "/" + repoName
	}

	c := exec.Command("git", "clone", repoLink, repoPath)

	return tea.ExecProcess(c, func(err error) tea.Msg {
		return CloneEditorMsg{err}
	})
}
func GitPull(repoName string) tea.Cmd {

	userName := services.GetUser()
	var repoPath string

	// if user is specified with an organization than the user takes precedence
	if services.GITHUB_ORG != "" && services.GITHUB_USER != "" {
		exec.Command("gh", "auth", "switch", "--user", services.GITHUB_USER)
		repoPath = "/Users/" + userName + "/Documents/" + services.GITHUB_USER + "/" + repoName
		// If an organization is specified without a user, use the organization name
	} else if services.GITHUB_ORG == "" && services.GITHUB_USER != "" {
		repoPath = "/Users/" + userName + "/Documents/" + services.GITHUB_USER + "/" + repoName
		exec.Command("gh", "auth", "switch", "--user", services.GITHUB_USER)
	} else if services.GITHUB_ORG != "" && services.GITHUB_USER == "" {
		repoPath = "/Users/" + userName + "/Documents/" + services.GITHUB_ORG + "/" + repoName
	}
	c := exec.Command("git", "-C", repoPath, "pull")

	return tea.ExecProcess(c, func(err error) tea.Msg {
		return GitPulledMessage{err}
	})
}
