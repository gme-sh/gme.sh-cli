package interaction

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func (c *CLI) RunApp() (err error) {
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
				Aliases: []string{"s", "sh", "c", "create", "new"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "URL",
						Aliases: []string{"u"},
					},
					&cli.StringFlag{
						Name:    "alias",
						Aliases: []string{"a"},
					},
					&cli.DurationFlag{
						Name:    "expire",
						Aliases: []string{"expires", "ex", "e", "duration", "dura", "d"},
					},
					&cli.BoolFlag{
						Name:    "hide-secret",
						Aliases: []string{"x"},
					},

					&cli.BoolFlag{
						Name:    "qr-code",
						Aliases: []string{"qr", "q"},
					},
				},
				Action: c.ActionShortURL,
			},
			{
				Name:    "delete",
				Usage:   "Deletes a short URL",
				Aliases: []string{"del", "remove", "rem"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "secret",
						Aliases:  []string{"s"},
						Required: true,
					},
				},
				Action: c.ActionDeleteURL,
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
			{
				Name:  "test",
				Usage: "Test input",
				Action: func(ctx *cli.Context) error {
					u := c.FindUrl(ctx)
					if u == "" {
						return errors.New("no url given")
					}
					log.Println("Testing:", u)
					log.Println("👉 IsURL?", IsURL(u))
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
