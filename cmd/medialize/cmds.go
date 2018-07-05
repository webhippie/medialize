package main

import (
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/Unknwon/com"
	"github.com/webhippie/medialize/pkg/config"
	"github.com/webhippie/medialize/pkg/photo"
	"github.com/webhippie/medialize/pkg/util"
	"gopkg.in/urfave/cli.v2"
)

// Commands defines all available sub-commands for this tool.
func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:      "photos",
			Usage:     "Sort photos",
			ArgsUsage: "<source> <destination>",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "rename",
					Value:       false,
					Usage:       "Rename the source insted of copying",
					EnvVars:     []string{"MEDIALIZE_PHOTOS_RENAME"},
					Destination: &config.Rename,
				},
			},
			Action: func(c *cli.Context) error {
				source := c.Args().Get(0)

				if source == "" {
					return fmt.Errorf("Please provide a source folder")
				}

				dest := c.Args().Get(1)

				if dest == "" {
					return fmt.Errorf("Please provide a dest folder")
				}

				if _, err := os.Stat(dest); os.IsNotExist(err) {
					if err := os.MkdirAll(dest, 0755); err != nil {
						return fmt.Errorf("Failed to create %s directory", dest)
					}

					logrus.Debugf("Created %s folder", dest)
				}

				logrus.Infof("Starting scan for photos")
				list, err := photo.Find(source)

				if err != nil {
					return fmt.Errorf("Failed to scan for files")
				}

				logrus.Infof("Found %d files within scan", len(list))

				for _, file := range list {
					if file.Valid() {
						logrus.Debugf("Trying to process %s", file)
					} else {
						logrus.Infof("Skipping %s as invalid", file)
						continue
					}

					creation, err := file.Creation()

					if err == nil {
						logrus.Debugf("Detected creation time as %s for %s", creation, file)

						if err := copyPhotoByCreation(source, dest, file); err != nil {
							logrus.Errorf("%s", err)
						}

						continue
					} else {
						logrus.Errorf("Failed to detect creation time for %s. %s", file, err)
					}

					checksum, err := file.Checksum()

					if err == nil {
						logrus.Debugf("Detected checksum as %s for %s", checksum, file)

						if err := copyPhotoByChecksum(source, dest, file); err != nil {
							logrus.Errorf("%s", err)
						}

						continue
					} else {
						logrus.Errorf("Failed to detect checksum for %s. %s", file, err)
					}
				}

				logrus.Infof("Finished processing!")
				return nil
			},
		},
	}
}

func copyPhotoByCreation(source, dest string, file *photo.File) error {
	dest = path.Join(
		dest,
		file.CalculatedCreation.Format("2006/01"),
	)

	if !com.IsDir(dest) {
		logrus.Debugf("Creating %s destination directory", dest)

		if err := os.MkdirAll(dest, 0755); err != nil {
			return fmt.Errorf("Failed to create destination directory. %s", err)
		}
	}

	for i := 0; i < 100000; i++ {
		destFile := path.Join(
			dest,
			fmt.Sprintf(
				"%s-%05d%s",
				file.CalculatedCreation.Format("20060102-150405-0700"),
				i,
				file.Ext(),
			),
		)

		if com.IsExist(destFile) {
			sourceChecksum, err := file.Checksum()

			if err != nil {
				return fmt.Errorf("Failed to process source checksum. %s", err)
			}

			destChecksum, err := util.Checksum(destFile)

			if err != nil {
				return fmt.Errorf("Failed to process destination checksum. %s", err)
			}

			if sourceChecksum == destChecksum {
				if config.Rename {
					logrus.Debugf("Dropping %s as it already exists as %s", file, destFile)
					return os.Remove(file.Path)
				} else {
					logrus.Debugf("Skipping %s as it already exists as %s", file, destFile)
					return nil
				}
			}

			logrus.Debugf("Next name %s already exists", destFile)
			continue
		}

		if config.Rename {
			if err := os.Rename(file.Path, destFile); err != nil {
				return fmt.Errorf("Failed to move %s.", file)
			} else {
				logrus.Debugf("Moved %s successfully", file)
			}

			return nil
		} else {
			if err := com.Copy(file.Path, destFile); err != nil {
				return fmt.Errorf("Failed to copy %s", file)
			} else {
				logrus.Debugf("Copied %s successfully", file)
			}

			return nil
		}
	}

	return fmt.Errorf("Failed to detect a destination for %s", file)
}

func copyPhotoByChecksum(source, dest string, file *photo.File) error {
	dest = path.Join(
		dest,
		"0000",
	)

	if !com.IsDir(dest) {
		logrus.Debugf("Creating %s destination directory", dest)

		if err := os.MkdirAll(dest, 0755); err != nil {
			return fmt.Errorf("Failed to create destination directory. %s", err)
		}
	}

	destFile := path.Join(
		dest,
		fmt.Sprintf(
			"%s%s",
			file.CalculatedChecksum,
			file.Ext(),
		),
	)

	if com.IsExist(destFile) {
		if config.Rename {
			logrus.Debugf("Dropping %s as it already exists as %s", file, destFile)
			return os.Remove(file.Path)
		} else {
			logrus.Debugf("Skipping %s as it already exists as %s", file, destFile)
			return nil
		}
	}

	if config.Rename {
		if err := os.Rename(file.Path, destFile); err != nil {
			return fmt.Errorf("Failed to move %s.", file)
		} else {
			logrus.Debugf("Moved %s successfully", file)
		}
	} else {
		if err := com.Copy(file.Path, destFile); err != nil {
			return fmt.Errorf("Failed to copy %s", file)
		} else {
			logrus.Debugf("Copied %s successfully", file)
		}
	}

	return nil
}
