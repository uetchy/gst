package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	// "github.com/daviddengcn/go-colortext"
	// "github.com/dustin/go-humanize"
	"strings"
	"os"
)

var commandDoctor = cli.Command{
	Name:   "doctor",
	Action: doDoctor,
	// Flags: []cli.Flag{
  //   cli.BoolFlag{
  // 		Name:  "fixup",
  // 		Usage: "automatically fix issues",
  // 	},
  // },
}

func doDoctor(c *cli.Context) {
  // fixupIssues := c.Bool("fixup")

	ghqPath, err := getGhqPath()
	if err != nil {
		fmt.Println("You must setup ghq first")
		os.Exit(1)
	}

	reposChannel := searchForRepos(ghqPath)

	// Listing repos
	for repo := range reposChannel {
    remoteOriginURL, _ := GitConfigGet(repo.Path, "remote.origin.url")
    target := compileTargetPathFromURL(remoteOriginURL)
    source := strings.TrimPrefix(repo.Path, ghqPath+"/")

    if target != source && !strings.Contains(source, "golang.org/x/") {
      fmt.Println("[bitbucket.org/uetchy/scent] git remote origin has changed:")
      fmt.Println("   Expected:\t", target)
      fmt.Println("   Actual:\t", source)
      // if fixupIssues {
      //   fmt.Println("   [Fixup]")
      // }
      fmt.Println()
    }

		// printlnWithColor(repo.Path, ct.Cyan)
		// printlnWithColor("-- "+humanize.Time(repo.ModTime), ct.Blue)
    //
		// for _, change := range changes[:len(changes)-1] {
		// 	staged := change[:1]
		// 	unstaged := change[1:2]
		// 	filename := change[3:]
    //
		// 	if staged == "?" {
		// 		printWithColor(staged, ct.Red)
		// 	} else {
		// 		printWithColor(staged, ct.Green)
		// 	}
		// 	printWithColor(unstaged, ct.Red)
		// 	fmt.Println("", filename)
		// }
	}
}
