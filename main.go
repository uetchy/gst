package main

import (
	"os"

	"github.com/urfave/cli"
)

// Version of this program
var Version = "HEAD"

// Commands are list of available commands
var Commands = []cli.Command{
	commandList,
	commandNew,
	commandRemove,
	commandDoctor,
	commandUpdate,
	commandFetch,
	commandTimeline,
}

func main() {
	app := cli.NewApp()
	app.Name = "gst"
	app.Version = Version
	app.Usage = "gst"
	app.Author = "Yasuaki Uechi"
	app.Email = "y@uechi.io"
	app.Commands = Commands

	// Declare default action
	app.HideHelp = true
	app.Flags = flagsOfList
	app.Action = doList

	app.Run(os.Args)
}
