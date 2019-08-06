package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/urfave/cli"
)

var commandDoctor = cli.Command{
	Name:   "doctor",
	Action: doDoctor,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "f, fix",
			Usage: "automatically fix issues",
		},
	},
}

func doDoctor(c *cli.Context) error {
	fixupIssues := c.Bool("fix")
	ghqPath := verifyGhqPath()
	reposChannel := searchForRepos(ghqPath)

	// Listing repos
	for repo := range reposChannel {
		remoteOriginURL, _ := GitConfigGet(repo.Path, "remote.origin.url")
		trimmedRemote := compileTargetPathFromURL(remoteOriginURL)
		trimmedLocal := strings.TrimPrefix(repo.Path, ghqPath+"/")

		if remoteOriginURL == "" {
			fmt.Println("===> 'remote.origin.url' does not exist in", trimmedLocal)
			if fixupIssues {
				okToChange := prompter.YN("===> Add "+trimmedLocal+" to 'remote.origin' ?", true)
				if okToChange {
					slp := strings.Split(trimmedLocal, "/")
					remotePathFromLocal := fmt.Sprintf("git@%s:%s/%s.git", slp[0], slp[1], slp[2])
					fmt.Println(remotePathFromLocal)
					err := GitRemoteAdd(repo.Path, "origin", remotePathFromLocal)
					if err != nil {
						fmt.Println("===> Fix failed because of", err)
					} else {
						fmt.Println("===> Add", remotePathFromLocal, "to", repo.Path)
					}
				} else {
					fmt.Println()
				}
			}
		} else if trimmedRemote != trimmedLocal && !strings.Contains(trimmedLocal, "golang.org/x/") {
			fmt.Println("===> 'remote.origin.url' has been changed")
			fmt.Println("===> local ", trimmedLocal)
			fmt.Println("===> remote", trimmedRemote)
			if fixupIssues {
				fmt.Println("===> Choose the right location for" + trimmedLocal)
				fmt.Println("[1] " + trimmedLocal)
				fmt.Println("[2] " + trimmedRemote)
				choice := prompter.Choose("===>", []string{"1", "2"}, "1")
				if choice == "1" {
					// Change remote.origin
					slp := strings.Split(trimmedLocal, "/")
					remotePathFromLocal := fmt.Sprintf("git@%s:%s/%s.git", slp[0], slp[1], slp[2])
					err := GitRemoteSetURL(repo.Path, "origin", remotePathFromLocal)
					if err != nil {
						fmt.Println("===> Failed because of", err)
						continue
					}
					fmt.Println("===> Change remote.origin.url to", remotePathFromLocal)
				} else {
					// Move directory
					localPathFromRemote := filepath.Join(ghqPath, trimmedRemote)
					fmt.Println(localPathFromRemote)
					if _, err := os.Stat(localPathFromRemote); os.IsExist(err) {
						fmt.Println("===> Fix failed because", localPathFromRemote, "already exist")
						continue
					}

					if err := os.Rename(repo.Path, localPathFromRemote); err != nil {
						fmt.Println("===> Fix failed because of", err)
						continue
					}
					fmt.Println("===> Move repository from", repo.Path, "to", localPathFromRemote)
				}
			}
		}
	}
	return nil
}
