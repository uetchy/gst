package main

import (
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "strings"
  "errors"
  "github.com/codegangsta/cli"
  "github.com/daviddengcn/go-colortext"
)

const Version string = "1.0.0"

type Repository struct {
  Type string
  Path string
}

func main() {
  app := cli.NewApp()
  app.Name = "gst"
  app.Version = Version
  app.Usage = "gst"
  app.Author = "Yasuaki Uechi"
  app.Email = "uetchy@randompaper.co"
  app.HideHelp = true
  app.Flags = []cli.Flag {
    cli.BoolFlag{
      Name: "short, s",
      Usage: "shorten result for pipeline processing",
    },
  }
  app.Action = commandList
  app.Commands = []cli.Command {
    {
      Name: "new",
      Action: commandNew,
    },
  }

  app.Run(os.Args)
}

func commandNew(c *cli.Context) {
  ghqPath, err := getGhqPath()
  if err != nil {
    fmt.Println("You must setup ghq first")
    os.Exit(1)
  }

  githubUser, err := getGithubUser()
  if err != nil {
    fmt.Println("You must set github.user first")
    os.Exit(1)
  }

  newPath := filepath.Join(ghqPath, "github.com", githubUser, c.Args().Get(0))

  err = exec.Command("mkdir", "-p", newPath).Run()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  if err = os.Chdir(newPath); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  if err = exec.Command("git", "init").Run(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  if err = exec.Command("touch", "README.md", "CHANGELOG.md", "LICENSE").Run(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println(newPath)
}

func commandRemove(c *cli.Context) {
  ghqPath, err := getGhqPath()
  if err != nil {
    fmt.Println("You must setup ghq first")
    os.Exit(1)
  }
}

func commandList(c *cli.Context) {
  ghqPath, err := getGhqPath()
  if err != nil {
    fmt.Println("You must setup ghq first")
    os.Exit(1)
  }

  shortExpression := c.Bool("short")
  repos := searchForRepos(ghqPath)

  for repo := range repos {
    status, err := gitStatus(repo.Path)
    if err != nil {
      continue
    }

    if shortExpression {
      fmt.Println(repo.Path)
      continue
    }

    printlnWithColor(repo.Path, ct.Cyan)

    changes := strings.Split(status, "\n")
    for _, change := range changes[:len(changes)-1] {
      staged := change[:1]
      unstaged := change[1:2]
      filename := change[3:]

      if staged == "?" {
        printWithColor(staged, ct.Red)
      } else {
        printWithColor(staged, ct.Green)
      }
      printWithColor(unstaged, ct.Red)
      fmt.Println("", filename)
    }

    fmt.Println()
  }
}

func getGhqPath() (string, error) {
  out, err := exec.Command("ghq", "root").Output()
  if err != nil {
    return "", err
  }
  return string(out)[:len(out)-1], nil
}

func getGithubUser() (string, error) {
  out, err := exec.Command("git", "config", "--get", "github.user").Output()
  if err != nil {
    return "", err
  }
  return string(out)[:len(out)-1], nil
}

func printWithColor(str string, color ct.Color) {
  ct.ChangeColor(color, false, ct.None, false)
  fmt.Print(str)
  ct.ResetColor()
}

func printlnWithColor(str string, color ct.Color) {
  printWithColor(str + "\n", color)
}

func gitStatus(targetPath string) (status string, err error) {
  if err := os.Chdir(targetPath); err != nil {
    return "", err
  }

  out, _ := exec.Command("git", "status", "-s").Output()
  if len(out) == 0 {
    return "", errors.New("No status changed")
  }

  return string(out), nil
}

func searchForRepos(rootPath string) <-chan Repository {
  repos := make(chan Repository)

  go func() {
    filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
      // skip file
      if !info.IsDir() {
        return nil
      }

      // skip directories which is not a repository
      if _, err := os.Stat(filepath.Join(path, ".git")); err != nil {
        return nil
      }

      repository := Repository{"git", path}
      repos <- repository

      return filepath.SkipDir
    })
    close(repos)
  }()

  return repos
}
