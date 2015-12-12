package main

import (
	"github.com/codegangsta/cli"
	"os"
)

const Version string = "HEAD"

var Commands = []cli.Command{
	commandList,
	commandNew,
	commandRemove,
	commandDoctor,
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
