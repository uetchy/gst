package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/daviddengcn/go-colortext"
)

var flagsOfFetch = []cli.Flag{
	cli.BoolFlag{
		Name:  "short, s",
		Usage: "shorten result for pipeline processing",
	},
}

var commandFetch = cli.Command{
	Name:   "fetch",
	Action: doFetch,
	Flags:  flagsOfFetch,
}

func doFetch(c *cli.Context) error {
	ghqPath := verifyGhqPath()
	repos := searchForRepos(ghqPath)

	// Listing repos
	for repo := range repos {
		printlnWithColor(repo.Path, ct.Cyan)
		out, err := GitFetch(repo.Path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if out != "" {
			fmt.Println(err)
		}
	}
	return nil
}
