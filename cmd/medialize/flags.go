package main

import (
	"github.com/webhippie/medialize/pkg/config"
	"gopkg.in/urfave/cli.v2"
)

// Flags defines all available flags for this command.
func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Value:       false,
			Usage:       "Enable debugging output",
			EnvVars:     []string{"MEDIALIZE_DEBUG"},
			Destination: &config.Debug,
		},
	}
}
