package services

import (
	"explorer/tools"
	"explorer/types"
	"io/ioutil"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ApplicationLoadedMsg struct {
	Err          error
	Repos        []list.Item
	Environments []list.Item
}

func LoadApplication() tea.Cmd {

	repo, err := ParseFileToRepos()
	if err != nil {
		tools.AnyCmd(ApplicationLoadedMsg{
			Err:          err,
			Repos:        nil,
		})
	}

	return tools.AnyCmd(ApplicationLoadedMsg{
		Err:          err,
		Repos:        repo,
	})
}

func GetDirectories() ([]string, error) {

	res := []string{}
	userName := GetUser()
	dirname := "/Users/" + userName + "/Documents/"+GITHUB_USER+"/"

	entries, err := ioutil.ReadDir(dirname)
	for _, entry := range entries {
		if entry.IsDir() {
			res = append(res, entry.Name())
		}
	}
	return res, err
}

func CheckDirExist(str string, s []string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

/* Functionality */
func ParseFileToRepos() ([]list.Item, error) {
	var repos []list.Item

	asdf, err := GetRepoList()
	if err != nil {
		return nil, err
	}

	for _, item := range asdf {

		repo := types.Item{}
		repo.SetTitle(item.Name)
		repo.SetDescription(item.HTMLURL)
		repos = append(repos, &repo)
	}

	return repos, nil
}
