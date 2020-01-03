package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	githubcli "newrehtse.info/go-training/github/internal"
)

const (
	searchRepositoriesEndpoint = "/search/repositories"
	apiGithubURL               = "https://api.github.com"
)

type githubRepo struct {
	url string
}

func NewGithubRepository() githubcli.GithubRepo {
	return &githubRepo{url: apiGithubURL}
}

func (g *githubRepo) SearchRepositories(q string) (repos githubcli.RepositoriesResponse, err error) {
	url := fmt.Sprintf("%v%v?q=%v", g.url, searchRepositoriesEndpoint, q)

	response, err := http.Get(url)
	if err != nil {
		return repos, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return repos, err
	}

	err = json.Unmarshal(contents, &repos)
	if err != nil {
		return repos, err
	}

	return
}
