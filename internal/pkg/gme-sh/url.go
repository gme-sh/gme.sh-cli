package gmesh

import (
	"github.com/urfave/cli/v2"
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
