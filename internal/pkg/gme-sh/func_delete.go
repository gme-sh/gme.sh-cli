package gmesh

import (
	"errors"
	"fmt"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/shortreq"
	"github.com/imroc/req"
	"github.com/urfave/cli/v2"
	"strings"
)

func (c *CLI) _actionDeleteURL(u, s string) (err error) {
	if u == "" {
		return errors.New("no url given")
	}
	// extract id
	for strings.HasSuffix(u, "/") {
		u = u[:len(u)-1]
	}
	if strings.Contains(u, "/") {
		u = u[strings.LastIndex(u, "/")+1:]
	}

	sp := newSpinner(fmt.Sprintf("📉 %s (with secret '%s')", u, s))
	// start spinner
	sp.Start()

	var res *req.Resp
	res, err = req.Delete(fmt.Sprintf("%s%s/%s", ApiUrl, u, s))

	// stop spinner
	sp.Stop()

	if err != nil {
		return
	}

	su := new(shortreq.Successable)
	if err = res.ToJSON(su); err != nil {
		return
	}

	if su.Success {
		fmt.Println("🖕", "Deleted", u, "(", su.Message, ")")
	} else {
		return errors.New(su.Message)
	}

	return
}

func (c *CLI) ActionDeleteURL(ctx *cli.Context) (err error) {
	u := c.FindUrl(ctx)
	s := ctx.String("secret")
	err = c._actionDeleteURL(u, s)
	return
}
