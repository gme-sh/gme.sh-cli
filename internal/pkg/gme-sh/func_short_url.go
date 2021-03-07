package gmesh

import (
	"errors"
	"fmt"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/short"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/shortreq"
	"github.com/imroc/req"
	"github.com/urfave/cli/v2"
)

type SuccessableCreate struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    *short.ShortURL `json:"data"`
}

func (c *CLI) ActionShortURL(ctx *cli.Context) (err error) {
	u := ctx.String("URL")
	// use pipe?
	if u == "" {
		if c.Pipe == "" {
			return errors.New("no url given")
		}
		u = c.Pipe
	}

	fmt.Println("🚀", u, "...")

	// make request
	var res *req.Resp
	res, err = req.Post(
		APIURL+"create",
		req.BodyJSON(&shortreq.CreateShortURLPayload{
			FullURL:            u,
			ExpireAfterSeconds: 0,
		}),
	)
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
	fmt.Println("🦾", PREFIX+sh.ID.String(), "[Secret: "+sh.Secret+"]")
	return
}
