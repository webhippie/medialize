package main

import (
	"os"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/webhippie/medialize/cmd"
	"github.com/webhippie/medialize/config"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if env := os.Getenv("MEDIALIZE_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:      "medialize",
		Version:   config.Version,
		Usage:     "Sort and filter your media files based on meta infos",
		Compiled:  time.Now(),
		ArgsUsage: "",

		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},

		Flags:    cmd.Flags(),
		Before:   cmd.Before(),
		Commands: cmd.Commands(),
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
