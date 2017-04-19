package cmd

import (
	"github.com/webhippie/medialize/config"
	"gopkg.in/urfave/cli.v2"
)

// Flags defines all available flags for this command.
func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "update, u",
			Value:       true,
			Usage:       "Enable auto updates",
			EnvVars:     []string{"MEDIALIZE_UPDATE"},
			Destination: &config.Update,
		},
	}
}
