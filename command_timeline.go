package main

import (
	"sort"
	"time"

	"github.com/codegangsta/cli"
	"github.com/daviddengcn/go-colortext"
	"github.com/dustin/go-humanize"
)

var flagsOfTimeline = []cli.Flag{
	cli.BoolFlag{
		Name:  "short, s",
		Usage: "shorten result for pipeline processing",
	},
}

var commandTimeline = cli.Command{
	Name:   "timeline",
	Action: doTimeline,
	Flags:  flagsOfTimeline,
}

func doTimeline(c *cli.Context) error {
	ghqPath := verifyGhqPath()
	reposChannel := searchForRepos(ghqPath)

	// Sort by time
	repos := []Repository{}
	for repo := range reposChannel {
		repos = append(repos, repo)
	}
	sort.Sort(RepositoriesByModTime{repos})

	// Listing repos
	for _, repo := range repos {
		duration := time.Now().Sub(repo.ModTime).Hours()
		var timeColor ct.Color
		if duration > 4320 { // 6 months
			timeColor = ct.Red
		} else if duration > 2160 { // 3 months
			timeColor = ct.Yellow
		} else if duration > 720 { // 1 month
			timeColor = ct.Green
		} else if duration > 504 { // 3 weeks
			timeColor = ct.Blue
		} else if duration > 168 { // 1 week
			timeColor = ct.Magenta
		} else {
			timeColor = ct.White
		}
		printlnWithColor(repo.Path+" ("+humanize.Time(repo.ModTime)+")", timeColor)
	}
	return nil
}
