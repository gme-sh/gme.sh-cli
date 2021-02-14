package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	// VERSION -> Muss noch in ne Config
	VERSION = "0.0.1-alpha"
)

func main() {

	cli.VersionPrinter = func(context *cli.Context) {
		fmt.Printf("GME-CLI-Version %s🚀\n", context.App.Version)
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
				Action: func(context *cli.Context) (err error) {
					log.Println(context.String("file"), context.Bool("qrcode"))
					return nil
				},
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

	err := app.Run(os.Args)

	if err != nil {
		log.Fatalln("An error has occurred:", err)
	}

}
