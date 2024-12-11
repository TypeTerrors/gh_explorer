package services

import (
	"encoding/json"
	"explorer/tools"
	"explorer/types"
	"fmt"
)

var GITHUB_PERSONAL_ACCESS_TOKEN string
var GITHUB_USER string
var GITHUB_ORG string

// GetRepoList returns a list of the first 100 repos in the user account
func GetRepoList() (types.RepoList, error) {
	headers := map[string]string{
		"Authorization":        "Bearer " + GITHUB_PERSONAL_ACCESS_TOKEN,
		"X-GitHub-Api-Version": "2022-11-28",
		"Accept":               "application/vnd.github+json",
	}

	urls := []string{
		// "https://api.github.com/user/repos?&per_page=100",
	}

	// If both GITHUB_ORG and GITHUB_USER are specified, the user takes precedence
	if GITHUB_ORG != "" && GITHUB_USER != "" {
		urls = append(urls, "https://api.github.com/user/repos?&per_page=100")
		// If an organization is specified without a user, use the organization name
	} else if GITHUB_ORG == "" && GITHUB_USER != "" {
		urls = append(urls, "https://api.github.com/user/repos?&per_page=100")
	} else if GITHUB_ORG != "" && GITHUB_USER == "" {
		urls = append(urls, "https://api.github.com/orgs/"+GITHUB_ORG+"/repos?&per_page=100")
	}

	// Channels for results and errors
	resultsChan := make(chan []byte, len(urls))
	errChan := make(chan error, len(urls))

	// Start a goroutine for each URL
	for _, url := range urls {
		go func(url string) {
			page, err := tools.Get(url, headers)
			if err != nil {
				errChan <- err
				return
			}
			resultsChan <- page
		}(url)
	}

	var res types.RepoList
	for i := 0; i < len(urls); i++ {
		select {
		case page := <-resultsChan:
			var pageRes types.RepoList
			err := json.Unmarshal(page, &pageRes)
			if err != nil {
				fmt.Println("GET request", "Error unmarshalling GET request 😭", err)
				return nil, err
			}
			res = append(res, pageRes...)
		case err := <-errChan:
			return nil, err
		}
	}

	return res, nil
}
