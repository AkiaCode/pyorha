package cli

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	tarfs "github.com/nlepage/go-tarfs"
	"github.com/spf13/cobra"
	gozstd "github.com/valyala/gozstd"
)

func serveCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "serve [path (without file extension)] [port=3000]",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			zstd, err := os.ReadFile(args[0] + ".star")
			if err != nil {
				return err
			}

			data, err := gozstd.Decompress(nil, zstd)
			if err != nil {
				return err
			}
			tempFile, err := os.CreateTemp("", "*")
			if err != nil {
				return err
			}
			defer os.RemoveAll(tempFile.Name())
			tempFile.Write(data)

			tf, err := os.Open(tempFile.Name())
			if err != nil {
				return err
			}
			defer tf.Close()

			tfs, err := tarfs.New(tf)
			if err != nil {
				return err
			}

			app := fiber.New(fiber.Config{
				DisableStartupMessage: true,
			})

			app.Use(filesystem.New(filesystem.Config{
				Root:         http.FS(tfs),
				NotFoundFile: "404.html",
			}))

			if len(args) == 1 {
				fmt.Println("started server on 0.0.0.0:3000, url: http://localhost:3000")
				app.Listen(":3000")
			} else {
				fmt.Println("started server on 0.0.0.0:" + args[1] + ", url: http://localhost:" + args[1])
				app.Listen(args[1])
			}

			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(serveCmd())
}
