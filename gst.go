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

const Version string = "0.1.0"
type Repository struct {
  Type string
  Path string
}

func main() {
  app := cli.NewApp()
  app.Name = "gst"
  app.Version = Version
  app.Usage = ""
  app.Author = "Yasuaki Uechi"
  app.Email = "uetchy@randompaper.co"
  app.Action = func(c *cli.Context) {
    out, err := exec.Command("ghq", "root").Output()
    if err != nil {
      fmt.Println("You must setup ghq first")
      os.Exit(1)
    }

    ghqPath := string(out)[:len(out)-1]
    repos := searchForRepos(ghqPath)

    for repo := range repos {
      status, err := gitStatus(repo.Path)
      if err != nil {
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

  app.Run(os.Args)
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
