package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/webhippie/medialize/pkg/config"
	"gopkg.in/urfave/cli.v2"
)

// Before gets called before any action on every execution.
func Before() cli.BeforeFunc {
	return func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)

		if config.Debug {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}

		return nil
	}
}
