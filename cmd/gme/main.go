package main

import (
	surveyCore "github.com/AlecAivazis/survey/v2/core"
	"github.com/full-stack-gods/gme.sh-cli/internal/api"
	"github.com/full-stack-gods/gme.sh-cli/internal/interaction"
	"github.com/mgutz/ansi"
	"log"
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

	cli := interaction.New(&api.API{
		ApiUrl: "https://gme.sh",
	})
	if err := cli.Run(); err != nil {
		log.Fatalln("Error:", err)
		return
	}
}
