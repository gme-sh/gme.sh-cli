package main

import (
	surveyCore "github.com/AlecAivazis/survey/v2/core"
	"github.com/full-stack-gods/gme.sh-cli/internal/api"
	"github.com/full-stack-gods/gme.sh-cli/internal/config"
	"github.com/full-stack-gods/gme.sh-cli/internal/interaction"
	"github.com/mgutz/ansi"
)

func main() {
	// override survey's poor choice of color
	// https://github.com/cli/cli
	surveyCore.TemplateFuncsWithColor["color"] = func(style string) string {
		switch style {
		case "white":
			//if cmdFactory.IOStreams.ColorSupport256() {
			//	return fmt.Sprintf("\x1b[%d;5;%dm", 38, 242)
			//}
			return ansi.ColorCode("default")
		default:
			return ansi.ColorCode(style)
		}
	}

	// read config
	cfg := config.ReadConfig()

	// start app
	cli := interaction.New(api.NewApi(cfg), cfg)
	_ = cli.Run()
}
