package photo

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/webhippie/medialize/pkg/checksum"
)

var (
	validExtensions = []string{
		".png",
		".gif",
		".jpg",
		".jpeg",
	}
)

// New initializes a new file instance.
func New(opts ...Option) *File {
	options := newOptions(opts...)

	return &File{
		path: options.Path,
		info: options.Info,
	}
}

// File represents a single source file.
type File struct {
	path string
	info os.FileInfo
}

// String converts a file into a string representation.
func (h *File) String() string {
	return h.path
}

// Valid checks if the file got a valid extension.
func (h *File) Valid() bool {
	ext := strings.ToLower(filepath.Ext(h.info.Name()))

	for _, check := range validExtensions {
		if check == ext {
			return true
		}
	}

	return false
}

// Ext tries to return a cleaned file extension.
func (h *File) Ext() string {
	result := strings.ToLower(filepath.Ext(h.info.Name()))

	switch result {
	case ".jpeg":
		return ".jpg"
	default:
		return result
	}
}

// Checksum generates a SHA256 checksum from file content.
func (h *File) Checksum() (string, error) {
	return checksum.New(h.path)
}

// Creation tries to extract the correct creation time.
func (h *File) Creation() (time.Time, error) {
	handle, err := os.Open(h.path)

	if err != nil {
		return time.Now(), fmt.Errorf("failed to open file: %w", err)
	}

	info, err := exif.Decode(handle)

	if err != nil && exif.IsCriticalError(err) {
		return time.Now(), fmt.Errorf("failed to parse exif: %w", err)
	}

	parsed, err := info.DateTime()

	if err != nil {
		return time.Now(), fmt.Errorf("failed to parse time: %w", err)
	}

	return parsed, nil
}

// Move simply moves the file to defined target.
func (h *File) Move(target string, byChecksum bool) {
	if final, ok := h.handle(target, byChecksum); ok {
		if err := h.rename(final); err != nil {
			log.Error().
				Err(err).
				Str("source", h.path).
				Str("target", final).
				Msg("Failed move photo")

			return
		}

		log.Info().
			Str("source", h.path).
			Str("target", final).
			Msg("Finished to move")
	}
}

// Copy simply copies the file to defined target.
func (h *File) Copy(target string, byChecksum bool) {
	if final, ok := h.handle(target, byChecksum); ok {
		s, err := os.Open(h.path)

		if err != nil {
			log.Error().
				Err(err).
				Str("source", h.path).
				Msg("Failed open source")

			return
		}

		defer func() { _ = s.Close() }()

		d, err := os.Create(final)

		if err != nil {
			log.Error().
				Err(err).
				Str("target", final).
				Msg("Failed open target")

			return
		}

		defer func() { _ = d.Close() }()

		if _, err := io.Copy(d, s); err != nil {
			log.Error().
				Err(err).
				Str("source", h.path).
				Str("target", final).
				Msg("Failed copy photo")

			return
		}

		log.Info().
			Str("source", h.path).
			Str("target", final).
			Msg("Finished to copy")
	}
}

func (h *File) handle(target string, byChecksum bool) (string, bool) {
	creation, err := h.Creation()

	if err == nil {
		d := path.Join(
			target,
			creation.UTC().Format("2006/01"),
		)

		if _, err := os.Stat(d); os.IsNotExist(err) {
			log.Info().
				Str("path", h.path).
				Str("directory", d).
				Msg("Creating target directory")

			if err := os.MkdirAll(d, 0755); err != nil {
				log.Error().
					Err(err).
					Str("path", h.path).
					Str("directory", d).
					Msg("Failed to create directory")

				return "", false
			}
		}

		for i := 0; i < 100000; i++ {
			f := path.Join(
				d,
				fmt.Sprintf(
					"%s-%05d%s",
					creation.UTC().Format("20060102-150405"),
					i,
					h.Ext(),
				),
			)

			if _, err := os.Stat(f); os.IsExist(err) {
				sourceChecksum, err := checksum.New(h.path)

				if err != nil {
					log.Error().
						Err(err).
						Str("source", h.path).
						Str("target", f).
						Msg("Failed to detect checksum")

					return "", false
				}

				targetChecksum, err := checksum.New(f)

				if err != nil {
					log.Error().
						Err(err).
						Str("source", h.path).
						Str("target", f).
						Msg("Failed to detect checksum")

					return "", false
				}

				if sourceChecksum == targetChecksum {
					return f, true
				}

				log.Info().
					Err(err).
					Str("source", h.path).
					Str("target", f).
					Msg("Target already exists")

				continue
			}

			return f, true
		}

		return "", false
	}

	log.Warn().
		Err(err).
		Str("path", h.path).
		Msg("Failed to detect exif data")

	if !byChecksum {
		log.Info().
			Str("path", h.path).
			Msg("Skipping checksum detection")

		return "", false
	}

	checksum, err := h.Checksum()

	if err == nil {
		d := path.Join(
			target,
			"0000",
		)

		if _, err := os.Stat(d); os.IsNotExist(err) {
			log.Info().
				Str("path", h.path).
				Str("directory", d).
				Msg("Creating target directory")

			if err := os.MkdirAll(d, 0755); err != nil {
				log.Error().
					Err(err).
					Str("path", h.path).
					Str("directory", d).
					Msg("Failed to create directory")

				return "", false
			}
		}

		f := path.Join(
			d,
			fmt.Sprintf(
				"%s%s",
				checksum,
				h.Ext(),
			),
		)

		return f, true
	}

	log.Warn().
		Err(err).
		Str("path", h.path).
		Msg("Failed to generate checksum")

	return "", false
}

func (h *File) rename(target string) error {
	err := os.Rename(h.path, target)

	if err != nil && strings.Contains(err.Error(), "invalid cross-device link") {
		return h.renameCrossDevice(target)
	}

	return err
}

func (h *File) renameCrossDevice(target string) error {
	src, err := os.Open(h.path)

	if err != nil {
		return fmt.Errorf("failed to open source: %w", err)
	}

	defer func() { _ = src.Close() }()

	dst, err := os.Create(target)

	if err != nil {
		return fmt.Errorf("failed to create target: %w", err)
	}

	defer func() { _ = dst.Close() }()

	if _, err = io.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	fi, err := os.Stat(h.path)

	if err != nil {
		_ = os.Remove(target)
		return fmt.Errorf("failed to stat source: %w", err)
	}

	if err := os.Chmod(target, fi.Mode()); err != nil {
		_ = os.Remove(target)
		return fmt.Errorf("failed to chmod target: %w", err)
	}

	return os.Remove(h.path)
}
