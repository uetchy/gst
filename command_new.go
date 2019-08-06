package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli"
)

var commandNew = cli.Command{
	Name:   "new",
	Action: doNew,
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func doNew(c *cli.Context) error {
	name := c.Args().Get(0)
	target := compileTargetPath(name)

	err := exec.Command("mkdir", "-p", target).Run()
	check(err)

	err = os.Chdir(target)
	check(err)

	err = exec.Command("git", "init").Run()
	check(err)

	f, err := os.Create(filepath.Join(target, "README.md"))
	check(err)
	defer f.Close()
	_, err = f.WriteString("# " + name + "\n")
	f.Sync()

	fmt.Println(target)
	return nil
}
