package main

import (
  "os"
  "github.com/codegangsta/cli"
)

var Commands = []cli.Command{
  commandList,
  commandNew,
  commandRemove,
}

func main() {
  app := cli.NewApp()
  app.Name = "gst"
  app.Version = Version
  app.Usage = "gst"
  app.Author = "Yasuaki Uechi"
  app.Email = "uetchy@randompaper.co"
  app.Commands = Commands

  // Declare default action
  app.HideHelp = true
  app.Flags = flagsOfList
  app.Action = doList

  app.Run(os.Args)
}
