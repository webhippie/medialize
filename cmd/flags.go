package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/webhippie/medialize/config"
)

func Flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolTFlag{
			Name:        "update, u",
			Usage:       "Enable auto updates",
			EnvVar:      "MEDIALIZE_UPDATE",
			Destination: &config.Update,
		},
	}
}
