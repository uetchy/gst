package main

import (
  "os"
  "fmt"
  "strings"
  "path/filepath"
  "github.com/codegangsta/cli"
)

var commandRemove = cli.Command{
  Name: "remove",
  Action: doRemove,
}

func doRemove(c *cli.Context) {
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

  if _, err := os.Stat(targetPath); err == nil {
    err := os.RemoveAll(targetPath)
    if err == nil {
      fmt.Println("Removed: "+targetPath)
    } else {
      fmt.Println(err)
    }
  } else {
    fmt.Println("Doesn't exist: "+targetPath)
  }
}
