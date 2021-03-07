package gmesh

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func (c *CLI) ParseArgs() (err error) {
	cli.VersionPrinter = func(context *cli.Context) {
		fmt.Printf("GME-CLI-Version 🚀 %s\n", context.App.Version)
	}

	app := &cli.App{
		Name:    "gme",
		Version: VERSION,
		Usage:   "urls goes brrrrrr",
		Commands: []*cli.Command{
			{
				Name:    "sh",
				Usage:   "Short an URL",
				Aliases: []string{"short", "s"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "URL",
						Aliases: []string{"u"},
						Value:   "",
					},
					&cli.BoolFlag{
						Name:    "qrcode",
						Aliases: []string{"qr", "q"},
					},
					&cli.StringFlag{
						Name:    "file",
						Aliases: []string{"f"},
						Value:   "",
					},
				},
				Action: c.ActionShortURL,
			},
			{
				Name:  "stats",
				Usage: "Shows you the stats of a shortened URL",
				Action: func(context *cli.Context) (err error) {
					return nil
				},
			},
			{
				Name:  "inspect",
				Usage: "Inspects an URL",
				Action: func(context *cli.Context) (err error) {
					return nil
				},
			},
		},
	}
	err = app.Run(os.Args)
	return
}
