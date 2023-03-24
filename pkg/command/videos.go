package command

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webhippie/medialize/pkg/video"
)

var (
	videosCmd = &cobra.Command{
		Use:   "videos",
		Short: "Sort videos",
		Run:   videosAction,
		Args:  cobra.NoArgs,
	}

	defaultVideosSource     = ""
	defaultVideosTarget     = ""
	defaultVideosRename     = false
	defaultVideosByChecksum = true
)

func init() {
	rootCmd.AddCommand(videosCmd)

	videosCmd.PersistentFlags().String("source", defaultVideosSource, "Path to source directory")
	viper.SetDefault("videos.source", defaultVideosSource)
	viper.BindPFlag("videos.source", videosCmd.PersistentFlags().Lookup("source"))

	videosCmd.PersistentFlags().String("target", defaultVideosTarget, "Path to target directory")
	viper.SetDefault("videos.target", defaultVideosTarget)
	viper.BindPFlag("videos.target", videosCmd.PersistentFlags().Lookup("target"))

	videosCmd.PersistentFlags().Bool("rename", defaultVideosRename, "Rename the source instead of copying")
	viper.SetDefault("videos.rename", defaultVideosRename)
	viper.BindPFlag("videos.rename", videosCmd.PersistentFlags().Lookup("rename"))

	videosCmd.PersistentFlags().Bool("by-checksum", defaultVideosByChecksum, "Rename by checksum as fallback")
	viper.SetDefault("videos.by_checksum", defaultVideosByChecksum)
	viper.BindPFlag("videos.by_checksum", videosCmd.PersistentFlags().Lookup("by-checksum"))
}

func videosAction(_ *cobra.Command, _ []string) {
	source := viper.GetString("videos.source")
	if source == "" {
		log.Error().
			Msg("Missing source flag")

		os.Exit(1)
	}

	target := viper.GetString("videos.target")
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
		Msg("Starting scan for videos")

	if err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		file := video.New(
			video.WithPath(path),
			video.WithInfo(info),
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

		if viper.GetBool("videos.rename") {
			file.Move(target, viper.GetBool("videos.by_checksum"))
		} else {
			file.Copy(target, viper.GetBool("videos.by_checksum"))
		}

		return err
	}); err != nil {
		log.Error().
			Err(err).
			Str("source", source).
			Str("target", target).
			Msg("Failed to scan videos")

		os.Exit(1)
	}

	log.Info().
		Str("source", source).
		Str("target", target).
		Msg("Finished video processing")
}
