package cmd

import (
	"os"

	"github.com/Sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

// Before gets called before any action on every execution.
func Before() cli.BeforeFunc {
	return func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)

		if c.Bool("update") {
			Update()
		}

		return nil
	}
}
