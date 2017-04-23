package cmd

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/webhippie/medialize/config"
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

		if config.Update {
			Update()
		}

		return nil
	}
}
