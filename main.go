package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/metakeule/fmtdate"
	"github.com/rwcarlsen/goexif/exif"
)

var (
	buildDate string
)

func main() {
	app := cli.NewApp()
	app.Name = "medialize"
	app.Version = buildDate
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "Sort and filter your media files"

	app.Before = func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)

		return nil
	}

	var source string
	var destination string
	var format string
	var rename bool

	app.Commands = []cli.Command{
		{
			Name:  "photos",
			Usage: "Sort photos",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "dest",
					Value:       "",
					Destination: &destination,
					Usage:       "Destination folder for sorted files",
				},
				cli.StringFlag{
					Name:        "format",
					Value:       "YYYY/MM",
					Destination: &format,
					Usage:       "Output naming to store the files",
				},
				cli.BoolFlag{
					Name:        "rename",
					Destination: &rename,
					Usage:       "Rename the source insted of copying",
				},
			},
			Action: func(c *cli.Context) {
				source = c.Args().First()

				if len(source) == 0 {
					logrus.Error("Please provide a source folder")
					return
				}

				if len(destination) > 0 {
					if _, err := os.Stat(destination); os.IsNotExist(err) {
						if err := os.MkdirAll(destination, 0755); err != nil {
							logrus.Error("Failed to create destination directory")
							return
						} else {
							logrus.Debug("Created destination folder")
						}
					}
				} else {
					destination, _ = os.Getwd()
				}

				logrus.Info("Starting scan for photos")
				fileList := []string{}

				err := filepath.Walk(
					source,
					func(path string, f os.FileInfo, err error) error {
						if f.IsDir() {
							return nil
						}

						fileList = append(
							fileList,
							path)

						return nil
					})

				if err != nil {
					logrus.Error("Failed to scan for files")
					return
				} else {
					logrus.Infof("Finished scan for %d files", len(fileList))
				}

				for _, file := range fileList {
					logrus.Infof("Parsing of %s in progress", file)

					handle, err := os.Open(file)

					if err != nil {
						logrus.Error("Failed to open file")
						continue
					}

					info, err := exif.Decode(handle)

					if err != nil {
						logrus.Error("Failed to parse file")
						continue
					}

					taken, err := info.DateTime()

					if err != nil {
						logrus.Error("Failed to get time")
						continue
					}

					for i := 0; i < 1000; i++ {
						name := NextName(file, destination, format, taken, i)

						if _, err := os.Stat(name); err == nil {
							logrus.Debugf("File already exists, increment suffix from %d", i)
							continue
						}

						if _, err := os.Stat(filepath.Dir(name)); os.IsNotExist(err) {
							if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
								logrus.Error("Failed to create formatted directory")
								break
							} else {
								logrus.Debug("Created formatted folder")
							}
						}

						if rename {
							if err := os.Rename(file, name); err != nil {
								logrus.Error("Failed to move file ", file, " - ", name, " - ", err)
								break
							} else {
								logrus.Debug("Moved file successfully")
							}
						} else {
							if err := os.Link(file, name); err != nil {
								logrus.Error("Failed to copy file ", file, " - ", name, " - ", err)
								break
							} else {
								logrus.Debug("Copied file successfully")
							}
						}

						break
					}
				}

				logrus.Info("Finished processing!")
			},
		},
	}

	app.Run(os.Args)
}

func NextName(file, dest, format string, taken time.Time, suffix int) string {
	_ext := strings.ToLower(filepath.Ext(file))
	_dest := TrimSuffix(dest, "/")
	_taken := taken.Format(time.RFC3339)
	_format := fmtdate.Format(format, taken)
	_suffix := fmt.Sprintf("%03d", suffix)

	return fmt.Sprintf("%s/%s/%s-%s%s", _dest, _format, _taken, _suffix, _ext)
}

func TrimPrefix(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		s = s[len(prefix):]
	}
	return s
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
