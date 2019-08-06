package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/urfave/cli"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/dustin/go-humanize"
)

var flagsOfList = []cli.Flag{
	cli.BoolFlag{
		Name:  "short, s",
		Usage: "shorten result for pipeline processing",
	},
}

var commandList = cli.Command{
	Name:   "list",
	Action: doList,
	Flags:  flagsOfList,
}

func doList(c *cli.Context) error {
	ghqPath := verifyGhqPath()
	reposChannel := searchForRepos(ghqPath)

	shortExpression := c.Bool("short")

	// Sort by time
	repos := []Repository{}
	for repo := range reposChannel {
		repos = append(repos, repo)
	}
	sort.Sort(RepositoriesByModTime{repos})

	// Listing repos
	for _, repo := range repos {
		uncommitedChanges, ccErr := GitStatus(repo.Path)
		unpushedCommits, pcErr := GitLog(repo.Path)
		if ccErr != nil && pcErr != nil {
			continue
		}

		if shortExpression {
			fmt.Println(repo.Path)
			continue
		}

		printlnWithColor(repo.Path+" ("+humanize.Time(repo.ModTime)+")", ct.Cyan)

		// print uncommited changes
		if ccErr == nil {
			printlnWithColor("uncommitted changes", ct.Magenta)
			for _, changes := range uncommitedChanges {
				staged := changes[:1]
				unstaged := changes[1:2]
				filename := changes[3:]

				if staged == "?" {
					printWithColor(staged, ct.Red)
				} else {
					printWithColor(staged, ct.Green)
				}
				printWithColor(unstaged, ct.Red)
				fmt.Println("", filename)
			}
		}

		// print unpushed commits
		if pcErr == nil {
			printlnWithColor("unpushed commits", ct.Magenta)
			for _, commit := range unpushedCommits {
				line := strings.Split(commit, " ")
				printWithColor(line[0], ct.Yellow)
				fmt.Println(" " + strings.Join(line[1:], " "))
			}
		}

		fmt.Println()
	}
	return nil
}
