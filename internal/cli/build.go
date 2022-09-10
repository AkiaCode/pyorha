package cli

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	gozstd "github.com/valyala/gozstd"
)

func buildCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "build [buildPath] [outPath]",
		Args: cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var buf bytes.Buffer
			err := compress(args[0], &buf)

			if err != nil {
				return err
			}

			fileWrite, err := os.OpenFile(args[1]+".star", os.O_CREATE|os.O_RDWR, os.FileMode(0600))

			if err != nil {
				return err
			}

			if _, err := io.Copy(fileWrite, &buf); err != nil {
				return err
			}
			return nil
		},
	}
}

// ref: https://gist.github.com/mimoo/25fc9716e0f1353791f5908f94d6e726
func compress(base string, buf io.Writer) error {
	zstdWriter := gozstd.NewWriter(buf)
	tarWriter := tar.NewWriter(zstdWriter)

	filepath.Walk(base, func(file string, fileInfo os.FileInfo, _ error) error {
		rel, err := filepath.Rel(base, file)

		if rel == "." {
			return nil
		}

		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fileInfo, rel)

		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(rel)
		header.ModTime = fileInfo.ModTime().Local()

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			data, err := os.Open(file)

			if err != nil {
				return err
			}

			if _, err := io.Copy(tarWriter, data); err != nil {
				return err
			}
		}
		return nil
	})

	if err := zstdWriter.Close(); err != nil {
		return err
	}

	if err := tarWriter.Close(); err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(buildCmd())
}
