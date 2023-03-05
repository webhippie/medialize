package command

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webhippie/medialize/pkg/photo"
)

var (
	photosCmd = &cobra.Command{
		Use:   "photos",
		Short: "Sort photos",
		Run:   photosAction,
		Args:  cobra.NoArgs,
	}

	defaultPhotosSource = ""
	defaultPhotosTarget = ""
	defaultPhotosRename = false
)

func init() {
	rootCmd.AddCommand(photosCmd)

	photosCmd.PersistentFlags().String("source", defaultPhotosSource, "Path to source directory")
	viper.SetDefault("photos.source", defaultPhotosSource)
	viper.BindPFlag("photos.source", photosCmd.PersistentFlags().Lookup("source"))

	photosCmd.PersistentFlags().String("target", defaultPhotosTarget, "Path to target directory")
	viper.SetDefault("photos.target", defaultPhotosTarget)
	viper.BindPFlag("photos.target", photosCmd.PersistentFlags().Lookup("target"))

	photosCmd.PersistentFlags().Bool("rename", defaultPhotosRename, "Rename the source instead of copying")
	viper.SetDefault("photos.rename", defaultPhotosRename)
	viper.BindPFlag("photos.rename", photosCmd.PersistentFlags().Lookup("rename"))
}

func photosAction(_ *cobra.Command, _ []string) {
	source := viper.GetString("photos.source")
	if source == "" {
		log.Error().
			Msg("Missing source flag")

		os.Exit(1)
	}

	target := viper.GetString("photos.target")
	if target == "" {
		log.Error().
			Msg("Missing target flag")

		os.Exit(1)
	}

	if _, err := os.Stat(target); os.IsNotExist(err) {
		if err := os.MkdirAll(target, 0755); err != nil {
			log.Error().
				Err(err).
				Str("target", target).
				Msg("Failed to create target")

			os.Exit(1)
		}

		log.Debug().
			Str("target", target).
			Msg("Created target directory")
	}

	log.Info().
		Str("source", source).
		Str("target", target).
		Msg("Starting scan for photos")

	list, err := photo.Find(source)

	if err != nil {
		log.Error().
			Err(err).
			Str("target", target).
			Msg("Failed to scan photos")

		os.Exit(1)
	}

	log.Info().
		Str("source", source).
		Str("target", target).
		Int("count", len(list)).
		Msg("Finished scan for photos")

	// for _, file := range list {
	// 	if file.Valid() {
	// 		logrus.Debugf("Trying to process %s", file)
	// 	} else {
	// 		logrus.Infof("Skipping %s as invalid", file)
	// 		continue
	// 	}

	// 	creation, err := file.Creation()

	// 	if err == nil {
	// 		logrus.Debugf("Detected creation time as %s for %s", creation, file)

	// 		if err := copyPhotoByCreation(source, dest, file); err != nil {
	// 			logrus.Errorf("%s", err)
	// 		}

	// 		continue
	// 	} else {
	// 		logrus.Errorf("Failed to detect creation time for %s. %s", file, err)
	// 	}

	// 	checksum, err := file.Checksum()

	// 	if err == nil {
	// 		logrus.Debugf("Detected checksum as %s for %s", checksum, file)

	// 		if err := copyPhotoByChecksum(source, dest, file); err != nil {
	// 			logrus.Errorf("%s", err)
	// 		}

	// 		continue
	// 	} else {
	// 		logrus.Errorf("Failed to detect checksum for %s. %s", file, err)
	// 	}
	// }

	log.Info().
		Msg("Finished photo processing")
}

// func copyPhotoByCreation(source, dest string, file *photo.File) error {
// 	dest = path.Join(
// 		dest,
// 		file.CalculatedCreation.Format("2006/01"),
// 	)

// 	if !com.IsDir(dest) {
// 		logrus.Debugf("Creating %s destination directory", dest)

// 		if err := os.MkdirAll(dest, 0755); err != nil {
// 			return fmt.Errorf("Failed to create destination directory. %s", err)
// 		}
// 	}

// 	for i := 0; i < 100000; i++ {
// 		destFile := path.Join(
// 			dest,
// 			fmt.Sprintf(
// 				"%s-%05d%s",
// 				file.CalculatedCreation.Format("20060102-150405-0700"),
// 				i,
// 				file.Ext(),
// 			),
// 		)

// 		if com.IsExist(destFile) {
// 			sourceChecksum, err := file.Checksum()

// 			if err != nil {
// 				return fmt.Errorf("Failed to process source checksum. %s", err)
// 			}

// 			destChecksum, err := util.Checksum(destFile)

// 			if err != nil {
// 				return fmt.Errorf("Failed to process destination checksum. %s", err)
// 			}

// 			if sourceChecksum == destChecksum {
// 				if config.Rename {
// 					logrus.Debugf("Dropping %s as it already exists as %s", file, destFile)
// 					return os.Remove(file.Path)
// 				}

// 				logrus.Debugf("Skipping %s as it already exists as %s", file, destFile)
// 				return nil
// 			}

// 			logrus.Debugf("Next name %s already exists", destFile)
// 			continue
// 		}

// 		if config.Rename {
// 			if err := os.Rename(file.Path, destFile); err != nil {
// 				return fmt.Errorf("Failed to move %s", file)
// 			}

// 			logrus.Debugf("Moved %s successfully", file)
// 			return nil
// 		}

// 		if err := com.Copy(file.Path, destFile); err != nil {
// 			return fmt.Errorf("Failed to copy %s", file)
// 		}

// 		logrus.Debugf("Copied %s successfully", file)
// 		return nil
// 	}

// 	return fmt.Errorf("Failed to detect a destination for %s", file)
// }

// func copyPhotoByChecksum(source, dest string, file *photo.File) error {
// 	dest = path.Join(
// 		dest,
// 		"0000",
// 	)

// 	if !com.IsDir(dest) {
// 		logrus.Debugf("Creating %s destination directory", dest)

// 		if err := os.MkdirAll(dest, 0755); err != nil {
// 			return fmt.Errorf("Failed to create destination directory. %s", err)
// 		}
// 	}

// 	destFile := path.Join(
// 		dest,
// 		fmt.Sprintf(
// 			"%s%s",
// 			file.CalculatedChecksum,
// 			file.Ext(),
// 		),
// 	)

// 	if com.IsExist(destFile) {
// 		if config.Rename {
// 			logrus.Debugf("Dropping %s as it already exists as %s", file, destFile)
// 			return os.Remove(file.Path)
// 		}

// 		logrus.Debugf("Skipping %s as it already exists as %s", file, destFile)
// 		return nil
// 	}

// 	if config.Rename {
// 		if err := os.Rename(file.Path, destFile); err != nil {
// 			return fmt.Errorf("Failed to move %s", file)
// 		}

// 		logrus.Debugf("Moved %s successfully", file)
// 	} else {
// 		if err := com.Copy(file.Path, destFile); err != nil {
// 			return fmt.Errorf("Failed to copy %s", file)
// 		}

// 		logrus.Debugf("Copied %s successfully", file)
// 	}

// 	return nil
// }
