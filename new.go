package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var commandNew = cli.Command{
	Name:   "new",
	Action: doNew,
}

func doNew(c *cli.Context) {
	ghqPath, err := getGhqPath()
	if err != nil {
		fmt.Println("You must setup 'ghq' command")
		os.Exit(1)
	}

	res := strings.Split(c.Args().Get(0), "/")
	var targetUser string
	var targetRepo string

	if len(res) == 1 {
		targetRepo = res[0]
		targetUser, err = getGithubUser()
		if err != nil {
			fmt.Println("You must set github.user first")
			fmt.Println("> git config --global github.user <name>")
			os.Exit(1)
		}
	} else {
		targetUser = res[0]
		targetRepo = res[1]
	}

	targetPath := filepath.Join(ghqPath, "github.com", targetUser, targetRepo)

	err = exec.Command("mkdir", "-p", targetPath).Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = os.Chdir(targetPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = exec.Command("git", "init").Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// if err = exec.Command("touch", "README.md", "CHANGELOG.md", "LICENSE").Run(); err != nil {
	//   fmt.Println(err)
	//   os.Exit(1)
	// }

	fmt.Println(targetPath)
}
