package command

import (
	"os"
	"path/filepath"

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

	defaultPhotosSource     = ""
	defaultPhotosTarget     = ""
	defaultPhotosRename     = false
	defaultPhotosByChecksum = true
)

func init() {
	rootCmd.AddCommand(photosCmd)

	photosCmd.PersistentFlags().String("source", defaultPhotosSource, "Path to source directory")
	viper.SetDefault("photos.source", defaultPhotosSource)
	_ = viper.BindPFlag("photos.source", photosCmd.PersistentFlags().Lookup("source"))

	photosCmd.PersistentFlags().String("target", defaultPhotosTarget, "Path to target directory")
	viper.SetDefault("photos.target", defaultPhotosTarget)
	_ = viper.BindPFlag("photos.target", photosCmd.PersistentFlags().Lookup("target"))

	photosCmd.PersistentFlags().Bool("rename", defaultPhotosRename, "Rename the source instead of copying")
	viper.SetDefault("photos.rename", defaultPhotosRename)
	_ = viper.BindPFlag("photos.rename", photosCmd.PersistentFlags().Lookup("rename"))

	photosCmd.PersistentFlags().Bool("by-checksum", defaultPhotosByChecksum, "Rename by checksum as fallback")
	viper.SetDefault("photos.by_checksum", defaultPhotosByChecksum)
	_ = viper.BindPFlag("photos.by_checksum", photosCmd.PersistentFlags().Lookup("by-checksum"))
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

	if err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		file := photo.New(
			photo.WithPath(path),
			photo.WithInfo(info),
		)

		if file.Valid() {
			log.Info().
				Str("path", file.String()).
				Msg("Starting to process")
		} else {
			log.Warn().
				Str("path", file.String()).
				Msg("Skipping as invalid")

			return nil
		}

		if viper.GetBool("photos.rename") {
			file.Move(target, viper.GetBool("photos.by_checksum"))
		} else {
			file.Copy(target, viper.GetBool("photos.by_checksum"))
		}

		return err
	}); err != nil {
		log.Error().
			Err(err).
			Str("source", source).
			Str("target", target).
			Msg("Failed to scan photos")

		os.Exit(1)
	}

	log.Info().
		Str("source", source).
		Str("target", target).
		Msg("Finished photo processing")
}
