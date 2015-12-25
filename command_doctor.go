package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	// "github.com/daviddengcn/go-colortext"
	// "github.com/dustin/go-humanize"
	"strings"
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
	ghqPath := verifyGhqPath()
	reposChannel := searchForRepos(ghqPath)

	// Listing repos
	for repo := range reposChannel {
		remoteOriginURL, _ := GitConfigGet(repo.Path, "remote.origin.url")
		target := compileTargetPathFromURL(remoteOriginURL)
		source := strings.TrimPrefix(repo.Path, ghqPath+"/")

		if remoteOriginURL == "" {
			fmt.Println("[" + source + "] 'remote.origin' doesn't exist:")
			fmt.Println("   Expected:\t", source)
			fmt.Println("   Actual:\t (no remote)")
			fmt.Println()
		} else if target != source && !strings.Contains(source, "golang.org/x/") {
			fmt.Println("[" + source + "] 'remote.origin' has changed:")
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
