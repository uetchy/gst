package main

import (
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/daviddengcn/go-colortext"
)

var flagsOfUpdate = []cli.Flag{
	cli.BoolFlag{
		Name:  "short, s",
		Usage: "shorten result for pipeline processing",
	},
}

var commandUpdate = cli.Command{
	Name:   "update",
	Action: doUpdate,
	Flags:  flagsOfUpdate,
}

func doUpdate(c *cli.Context) error {
	ghqPath := verifyGhqPath()
	repos := searchForRepos(ghqPath)

	// Listing repos
	for repo := range repos {
		printlnWithColor(repo.Path, ct.Cyan)
		out, err := GitPull(repo.Path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if strings.Contains(out, "Already up-to-date") != true {
			fmt.Println(out)
		}
	}

	return nil
}
