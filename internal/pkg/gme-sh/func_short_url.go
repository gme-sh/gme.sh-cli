package gmesh

import (
	"errors"
	"fmt"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/short"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/shortreq"
	"github.com/imroc/req"
	"github.com/mdp/qrterminal"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"time"
)

type SuccessableCreate struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    *short.ShortURL `json:"data"`
}

func (c *CLI) ActionShortURL(ctx *cli.Context) (err error) {
	u := c.FindUrl(ctx)
	if u == "" {
		return errors.New("no url given")
	}

	// parse duration
	ex := ctx.Duration("expire")
	alias := ctx.String("alias")

	// create payload
	payload := &shortreq.CreateShortURLPayload{
		FullURL:            u,
		ExpireAfterSeconds: int(ex.Seconds()),
		PreferredAlias:     short.ShortID(alias),
	}

	// alias?
	if alias != "" {
		fmt.Println("🚀", u, "->", Prefix+alias, "...")
	} else {
		fmt.Println("🚀", u, "...")
	}

	// make request
	var res *req.Resp
	res, err = req.Post(
		ApiUrl+"create",
		req.BodyJSON(payload),
	)
	if res == nil {
		return errors.New("response was null")
	}

	if err != nil {
		return
	}

	// parse response
	s := new(SuccessableCreate)
	err = res.ToJSON(s)
	if err != nil {
		return
	}
	if !s.Success {
		return errors.New(s.Message)
	}

	sh := s.Data

	var secret string
	if ctx.Bool("hide-secret") {
		secret = repeat("*", len(sh.Secret))
	} else {
		secret = sh.Secret
	}

	// output url and expiration
	if ex != 0 {
		fmt.Println("⏰", "Expires at", time.Now().Add(ex).Format("02.01.2006 15:04:05"))
	}
	url := Prefix + sh.ID.String()
	fmt.Println("🦾", url, "[Secret: "+secret+"]")

	// display qr code
	if ctx.Bool("qr-code") {
		config := qrterminal.Config{
			BlackChar: qrterminal.WHITE,
			WhiteChar: qrterminal.BLACK,
			QuietZone: 1,
			Writer:    os.Stdout,
			Level:     qrterminal.L,
		}
		qrterminal.GenerateWithConfig(url, config)
	}

	return
}

func repeat(str string, num int) string {
	var builder strings.Builder
	for i := 0; i < num; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
