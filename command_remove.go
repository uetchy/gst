package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Songmu/prompter"
	"github.com/urfave/cli"
)

var commandRemove = cli.Command{
	Name:   "remove",
	Action: doRemove,
}

// IsEmpty checks if directory is empty
func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

func doRemove(c *cli.Context) error {
	target := compileTargetPath(c.Args().Get(0))

	if _, err := os.Stat(target); err == nil {
		if !prompter.YN("Remove? "+target, true) {
			os.Exit(0)
		}

		// Remove specified directory
		err := os.RemoveAll(target)
		if err == nil {
			fmt.Println("Removed: " + target)
		} else {
			fmt.Println(err)
		}

		// Remove parent dirs if empty
		ghqPath, _ := getGhqPath()
		target = filepath.Dir(target)
		for target != ghqPath {
			if e, _ := IsEmpty(target); !e {
				break
			}
			fmt.Println(target, ghqPath)
			err := os.RemoveAll(target)
			if err == nil {
				fmt.Println("Removed: " + target)
			} else {
				fmt.Println(err)
				break
			}
			target = filepath.Dir(target)
		}
	} else {
		fmt.Println("Doesn't exist: " + target)
	}

	return nil
}
