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
		Name:    "🚀 GME.sh",
		Usage:   "long url goes brrr",
		Version: Version,
		Commands: []*cli.Command{
			{
				Name:    "short",
				Usage:   "Short an loooong URL",
				Aliases: []string{"s", "sh"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "URL",
						Aliases: []string{"u"},
					},
					&cli.BoolFlag{
						Name:    "show-secret",
						Aliases: []string{"s"},
					},
					&cli.StringFlag{
						Name:    "alias",
						Aliases: []string{"a"},
					},
					&cli.BoolFlag{
						Name:    "qr-code",
						Aliases: []string{"qr", "q"},
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
		ExitErrHandler: func(context *cli.Context, err error) {
			fmt.Println("🤬📉 ERROR:", err)
		},
	}
	err = app.Run(os.Args)
	return
}
