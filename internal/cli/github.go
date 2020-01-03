package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	githubcli "newrehtse.info/go-training/github/internal"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const searchFlag = "id"
const outputFlag = "output"

// InitBeersCmd initialize beers command
func SearchRepoCmd(input githubcli.GithubRepo, output githubcli.SaveRepo) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "search",
		Short: "Search some repository",
		Run:   runSearchFn(input, output),
	}

	beersCmd.Flags().StringP(searchFlag, "q", "", "query for the search")
	beersCmd.Flags().StringP(outputFlag, "o", "", "csv file to save the data")

	return beersCmd
}

func runSearchFn(input githubcli.GithubRepo, output githubcli.SaveRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		search, _ := cmd.Flags().GetString(searchFlag)
		file, _ := cmd.Flags().GetString(outputFlag)

		repositories, err := input.SearchRepositories(search)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(repositories)

		err = output.StoreRepositoriesResponse(repositories, file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
