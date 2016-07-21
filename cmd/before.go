package cmd

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// Before gets called before any action on every execution.
func Before() BeforeFunc {
	return func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)

		if c.BoolT("update") {
			Update()
		}

		return nil
	}
}
