package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

var commandNew = cli.Command{
	Name:   "new",
	Action: doNew,
}

func doNew(c *cli.Context) {
	target := compileTargetPath(c.Args().Get(0))

	err := exec.Command("mkdir", "-p", target).Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = os.Chdir(target); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = exec.Command("git", "init").Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = exec.Command("touch", "README.md", "CHANGELOG.md").Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(target)
}
