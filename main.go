package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/webhippie/medialize/cmd"
	"github.com/webhippie/medialize/config"
)

var (
	version    string
	versionSha string
)

func main() {
	app := cli.NewApp()
	app.Name = "medialize"
	app.Version = config.Version
	app.Usage = "Sort and filter your media files based on meta infos"
	app.ArgsUsage = ""

	app.Authors = []cli.Author{
		{"Thomas Boerger", "thomas@webhippie.de"},
	}

	app.HideHelp = false

	app.Before = cmd.Before()
	app.Flags = cmd.Flags()
	app.Commands = cmd.Commands()

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	app.Run(os.Args)
}
