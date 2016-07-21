package cmd

import (
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/webhippie/medialize/photo"
)

// Commands defines all available sub-commands for this tool.
func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:      "photos",
			Usage:     "Sort photos",
			ArgsUsage: "<source> <destination>",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "rename",
					Usage: "Rename the source insted of copying",
				},
			},
			Action: func(c *cli.Context) {
				source := c.Args().Get(0)

				if len(source) == 0 {
					logrus.Error("Please provide a source folder")
					return
				}

				dest := c.Args().Get(1)

				if len(source) == 0 {
					logrus.Error("Please provide a dest folder")
					return
				}

				if len(dest) > 0 {
					if _, err := os.Stat(dest); os.IsNotExist(err) {
						if err := os.MkdirAll(dest, 0755); err != nil {
							logrus.Errorf(
								"Failed to create %s directory",
								dest,
							)

							return
						}

						logrus.Debugf(
							"Created %s folder",
							dest,
						)
					}
				} else {
					dest, _ = os.Getwd()
				}

				logrus.Info("Starting scan for photos")
				fileList, err := photo.FindFiles(source)

				if err != nil {
					logrus.Error("Failed to scan for files")
					return
				}

				logrus.Infof("Finished scan for %d files", len(fileList))

				for _, file := range fileList {
					if photo.ValidExtension(file) {
						logrus.Infof(
							"Parsing of %s in progress",
							file)
					} else {
						logrus.Infof(
							"Skipping %s, invalid ext",
							file)

						continue
					}

					for i := 0; i < 100000; i++ {
						name, _ := photo.NextName(file, dest, i)

						if _, err := os.Stat(name); err == nil {
							continue
						}

						if _, err := os.Stat(filepath.Dir(name)); os.IsNotExist(err) {
							if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
								logrus.Error("Failed to create formatted directory")
								break
							}
						}

						if c.Bool("rename") {
							if err := os.Rename(file, name); err != nil {
								logrus.Errorf("Failed to move %s", file)
								break
							} else {
								logrus.Debugf("Moved %s successfully", file)
							}
						} else {
							if err := os.Link(file, name); err != nil {
								logrus.Errorf("Failed to copy %s", file)
								break
							} else {
								logrus.Debugf("Copied %s successfully", file)
							}
						}

						break
					}
				}

				logrus.Info("Finished processing!")
			},
		},
	}
}
