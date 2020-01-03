package githubcli

type RepositoriesResponse struct {
	Total        int          `json:"total_count"`
	Repositories []Repository `json:"items"`
}

type Repository struct {
	RepositoryId int     `json:"id"`
	Name         string  `json:"name"`
	FullName     string  `json:"full_name"`
	Private      bool    `json:"private"`
	Owner        Owner   `json:"owner"`
	HtmlUrl      string  `json:"html_url"`
	Description  string  `json:"description"`
	Fork         bool    `json:"fork"`
	ApiUrl       string  `json:"url"`
	License      License `json:"license"`
}

type Owner struct {
	Login   string `json:"login"`
	Type    string `json:"type"`
	HtmlUrl string `json:"html_url"`
	ApiUrl  string `json:"url"`
}
type License struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type GithubRepo interface {
	SearchRepositories(q string) (RepositoriesResponse, error)
}

type SaveRepo interface {
	StoreRepositoriesResponse(list RepositoriesResponse, dest string) error
}

func (repo RepositoriesResponse) GetHeaders() []string {
	return []string{"id", "name", "full_name", "private", "owner", "html_url", "description", "fork", "url", "license"}
}

func (repo Repository) ToSlice() []string {
	owner := repo.Owner
	license := repo.License
	private := "no"
	if repo.Private {
		private = "si"
	}

	fork := "no"
	if repo.Fork {
		fork = "si"
	}
	return []string{string(repo.RepositoryId), repo.Name, repo.FullName, private, owner.Login, repo.HtmlUrl, repo.Description, fork, repo.ApiUrl, license.Name}
}
