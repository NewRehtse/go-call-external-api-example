package main

import (
	"github.com/spf13/cobra"
	githubcli "newrehtse.info/go-training/github/internal"
	"newrehtse.info/go-training/github/internal/cli"
	csv "newrehtse.info/go-training/github/internal/storage/csv"
	github "newrehtse.info/go-training/github/internal/storage/github"
)

func main() {
	var input githubcli.GithubRepo
	var output githubcli.SaveRepo

	input = github.NewGithubRepository()
	output = csv.NewSaveCsvRepository()

	rootCmd := &cobra.Command{Use: "github-cli"}
	rootCmd.AddCommand(cli.SearchRepoCmd(input, output))
	rootCmd.Execute()
}
