package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gme",
		Usage: "urls goes brrrrrr",
		Commands: []*cli.Command{
			{
				Name:    "sh",
				Usage:   "Short an URL",
				Aliases: []string{"short", "s"},
				Flags: []cli.Flag{
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
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatalln("An error has occurred:", err)
	}
}
