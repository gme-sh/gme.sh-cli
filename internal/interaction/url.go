package interaction

import (
	"github.com/urfave/cli/v2"
	"net/url"
	"strings"
)

func (c *CLI) FindUrl(ctx *cli.Context) (u string) {
	u = ctx.String("URL")
	if u == "" {
		// url from args
		if ctx.Args().Len() > 0 {
			u = strings.Join(ctx.Args().Slice(), " ")
			// url from pipe
		} else if c.Pipe != "" {
			u = c.Pipe
		}
	}
	return
}

func IsURL(str string) bool {
	if len(str) > 1000 {
		return false
	}
	if strings.Contains(str, " ") {
		return false
	}
	// try to parse url
	_, err := url.Parse(str)
	return err == nil
}
