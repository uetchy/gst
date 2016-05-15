package main

import (
	"fmt"
	"github.com/Songmu/prompter"
	"github.com/codegangsta/cli"
	"os"
)

var commandRemove = cli.Command{
	Name:   "remove",
	Action: doRemove,
}

func doRemove(c *cli.Context) error {
	target := compileTargetPath(c.Args().Get(0))

	if _, err := os.Stat(target); err == nil {
		if !prompter.YN("Remove? "+target, true) {
			os.Exit(0)
		}

		err := os.RemoveAll(target)
		if err == nil {
			fmt.Println("Removed: " + target)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Doesn't exist: " + target)
	}

	return nil
}
