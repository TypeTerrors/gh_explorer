package commands

import (
	"explorer/services"
	"explorer/tools"
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const divisor = 1

const (
	ChooseRepo status = iota
	Loading
)

/* STYLING */
var (
	pink             = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	UnFocusRepoStyle = lipgloss.NewStyle().
				Padding(0, 2).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62"))
	FocusRepoStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Border(lipgloss.HiddenBorder())
	StatusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
	helpStyle = func(s string) string {
		return lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).
			Margin(0, 4, 0, 4).
			Padding(1).
			Italic(true).
			Render(s)
	}
)

type Load struct{}
type Model struct {
	Focused  status
	Lists    []list.Model
	RepoList []list.Item

	LoadingView viewport.Model

	LoadingMessage string

	Quitting  bool
	DidItWork string
	Height    int
	Width     int
}

func (m *Model) Init() tea.Cmd {
	m.Focused = Loading
	m.LoadingView.SetContent(m.LoadingMessage)
	m.LoadingView.Style = pink
	return tea.Batch(tea.Cmd(func() tea.Msg {
		return Load{}
	}))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case Load:
		m.Focused = Loading
		var cmd tea.Cmd
		m.LoadingView, cmd = m.LoadingView.Update(msg)
		return &m, tea.Batch(cmd, services.LoadApplication())

	case tea.WindowSizeMsg:
		m.Height = msg.Height
		m.Width = msg.Width

		if m.Focused != Loading {
			FocusRepoStyle.
				Width(msg.Width).
				Height(msg.Height - 6)
			UnFocusRepoStyle.
				Width(msg.Width).
				Height(msg.Height - 6)
			m.Focused = ChooseRepo
			m.Lists[ChooseRepo].SetSize(msg.Width, msg.Height-6)
		} else {
			m.LoadingView.Width = msg.Width
			m.LoadingView.Height = msg.Height
			m.Focused = Loading
			var cmd tea.Cmd
			m.LoadingView, cmd = m.LoadingView.Update(msg)
			return &m, cmd
		}

	case services.ApplicationLoadedMsg:
		var cmd tea.Cmd
		if len(msg.Repos) > 0 {

			m.Focused = ChooseRepo

			FocusRepoStyle.
				Width(m.Width).
				Height(m.Height - 6)
			UnFocusRepoStyle.
				Width(m.Width).
				Height(m.Height - 6)

			defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
			defaultList.SetShowHelp(false)

			m.Lists = []list.Model{defaultList}

			m.Lists[ChooseRepo].Title = "Choose your Repo: "

			cmd := m.Lists[ChooseRepo].SetItems(msg.Repos)
			cmds = append(cmds, cmd)

			m.Lists[ChooseRepo].SetSize(m.Width/divisor, m.Height-6)

			m.Lists[ChooseRepo], cmd = m.Lists[ChooseRepo].Update(msg)
			cmds = append(cmds, cmd)

			cmds = append(cmds, m.Lists[m.Focused].FilterInput.Focus())
		} else {
			m.Focused = Loading
			m.LoadingView, cmd = m.LoadingView.Update(msg)
			return &m, tea.Batch(cmd)
		}
	case OpenEditorMsg:
		if msg.err != nil {
			m.DidItWork = "Oops!"
		} else {
			m.DidItWork = "code ."
		}
		showSelection := m.Lists[m.Focused].NewStatusMessage(StatusMessageStyle(m.DidItWork))
		cmds = append(cmds, showSelection)
	case GitPulledMessage:
		if msg.err != nil {
			m.DidItWork = "Oops!"
		} else {
			m.DidItWork = "Git pull"
		}
		showSelection := m.Lists[m.Focused].NewStatusMessage(StatusMessageStyle(m.DidItWork))
		cmds = append(cmds, showSelection)
	case CloneEditorMsg:

		if msg.err != nil {
			m.DidItWork = "Oops!"
		} else {
			m.DidItWork = "Git clone"
		}
		if m.Focused != Loading {
			showSelection := m.Lists[m.Focused].NewStatusMessage(StatusMessageStyle(m.DidItWork))
			cmds = append(cmds, showSelection)
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c" + "q":
			m.Quitting = true
			cmds = append(cmds, tea.Quit)
			//return m, tea.Quit
		case "ctrl+r":

			m.RepoList = []list.Item{}

			var res list.Model

			cmd := m.Lists[ChooseRepo].SetItems(m.RepoList)
			m.Lists[ChooseRepo] = res
			cmds = append(cmds, cmd)

			m.Focused = Loading
			m.LoadingView.Height = m.Height
			m.LoadingView.Width = m.Width
			cmds = append(cmds, m.Init())
			m.LoadingMessage = tools.RadnomnLoadingMessage()
			return &m, tea.Batch(cmds...)

		case "ctrl+p":
			if m.Focused == ChooseRepo {
				// git pull selected repo
				repoName := m.Lists[m.Focused].SelectedItem().FilterValue()
				folders, err := services.GetDirectories()
				if err != nil {
					m.DidItWork = "Oops!"
					cmds = append(cmds, tea.Quit)
					//return m, tea.Quit
				}
				exists := services.CheckDirExist(repoName, folders)
				if exists {
					m.DidItWork = "Git Pull"
					showSelection := m.Lists[m.Focused].NewStatusMessage(StatusMessageStyle(m.DidItWork))
					cmds = append(cmds, GitPull(repoName), showSelection)
					//return m, tea.Batch(showSelection, GitPull(repoName))
				} else {
					m.DidItWork = "Cloned! Repo already exists."
					var repoLink string

					// if user is specified with an organization than the user takes precedence
					if services.GITHUB_ORG != "" && services.GITHUB_USER != "" {
						repoLink = "https://github.com/" + services.GITHUB_USER + "/" + repoName + ".git"
						// If an organization is specified without a user, use the organization name
					} else if services.GITHUB_ORG == "" && services.GITHUB_USER != "" {
						repoLink = "https://github.com/" + services.GITHUB_USER + "/" + repoName + ".git"
					} else if services.GITHUB_ORG != "" && services.GITHUB_USER == "" {
						repoLink = "https://github.com/" + services.GITHUB_ORG + "/" + repoName + ".git"
					}

					showSelection := m.Lists[m.Focused].NewStatusMessage(StatusMessageStyle(m.DidItWork))
					cmds = append(cmds, CloneEditor(repoLink, repoName), showSelection)
					//return m, tea.Batch(showSelection, CloneEditor(repoLink, repoName))
				}
			}
		case "enter":
			if m.Focused == ChooseRepo {
				repoName := m.Lists[m.Focused].SelectedItem().FilterValue()
				var repoLink string

				// if user is specified with an organization than the user takes precedence
				if services.GITHUB_ORG != "" && services.GITHUB_USER != "" {
					repoLink = "https://github.com/" + services.GITHUB_USER + "/" + repoName + ".git"
					// If an organization is specified without a user, use the organization name
				} else if services.GITHUB_ORG == "" && services.GITHUB_USER != "" {
					repoLink = "https://github.com/" + services.GITHUB_USER + "/" + repoName + ".git"
				} else if services.GITHUB_ORG != "" && services.GITHUB_USER == "" {
					repoLink = "https://github.com/" + services.GITHUB_ORG + "/" + repoName + ".git"
				} else {
					m.DidItWork = "Oops!"
					// cmds = append(cmds, tea.Quit)
				}

				folders, err := services.GetDirectories()
				if err != nil {
					m.DidItWork = "Oops!"
					cmds = append(cmds, tea.Quit)
					//return m, tea.Quit
				}
				exists := services.CheckDirExist(repoName, folders)
				if exists {
					m.DidItWork = "code ."
					fmt.Println(m.DidItWork)
					showSelection := m.Lists[m.Focused].NewStatusMessage(StatusMessageStyle(m.DidItWork))
					cmds = append(cmds, OpenEditor(repoName), showSelection)
					// return m, tea.Batch(showSelection, OpenEditor(repoName))
				} else {
					m.DidItWork = "Git Clone"
					fmt.Println(m.DidItWork)
					showSelection := m.Lists[m.Focused].NewStatusMessage(StatusMessageStyle(m.DidItWork))
					cmds = append(cmds, CloneEditor(repoLink, repoName), showSelection)
					//return m, tea.Batch(showSelection, CloneEditor(repoLink, repoName))
				}
			} else {
				var cmd tea.Cmd
				m.Focused = Loading
				m.LoadingView, cmd = m.LoadingView.Update(msg)
				return &m, tea.Batch(cmd)
			}
		}
	}

	if m.Focused == ChooseRepo {
		var cmd tea.Cmd
		m.Lists[m.Focused], cmd = m.Lists[m.Focused].Update(msg)
		cmds = append(cmds, cmd, m.Lists[m.Focused].FilterInput.Focus())
		return &m, tea.Batch(cmds...)
	} else {
		var cmd tea.Cmd
		m.LoadingView, cmd = m.LoadingView.Update(msg)
		cmds = append(cmds, cmd)
		return &m, tea.Batch(cmds...)
	}
}

func (m *Model) View() string {
	if m.Quitting {
		return ""
	}

	helpView := "ctrl+c to quit ◦ ctrl+p to git pull ◦ ctrl+r refresh ◦ / to search ◦ Esc to cancel search"
	switch m.Focused {
	case Loading:
		return m.LoadingView.View()
	case ChooseRepo:
		repoView := m.Lists[ChooseRepo].View()
		return lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.JoinHorizontal(lipgloss.Left, FocusRepoStyle.Render(repoView)),
			helpStyle(helpView),
		)
	}
	return "Didn't load..."
}

func Start() {

	// Define flags
	patPtr := flag.String("PAT", "", "Personal Access Token")
	userPtr := flag.String("USER", "", "github user account")
	orgPtr := flag.String("ORG", "", "github user account")
	flag.Parse()

	services.GITHUB_PERSONAL_ACCESS_TOKEN = *patPtr
	services.GITHUB_USER = *userPtr
	services.GITHUB_ORG = *orgPtr

	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	defaultList.SetShowHelp(false)

	app := &Model{
		LoadingMessage: tools.RadnomnLoadingMessage(),
		Lists:          []list.Model{defaultList},
		Focused:        Loading,
		LoadingView: viewport.Model{
			Style: pink,
		},
	}

	if _, err := tea.NewProgram(app, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
