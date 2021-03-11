package gmesh

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/short"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/shortreq"
	"github.com/imroc/req"
	"github.com/mdp/qrterminal"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
	"time"
)

type SuccessableCreate struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    *short.ShortURL `json:"data"`
}

const (
	AliasOccupiedNewAlias       = "Enter new alias"
	AliasOccupiedDeleteOld      = "Delete old alias"
	AliasOccupiedGenerateRandom = "Generate random alias"
	AliasOccupiedNothing        = "Nothing"
)

func (c *CLI) _recoverShortURL(s *SuccessableCreate, u, alias string, ex time.Duration, hideSecret bool, qrCode bool) (err error) {
	if s.Message == "alias not available" {
		// ask to delete
		fmt.Println("🏷 Alias occupied.")

		sel := &survey.Select{
			Message: "What do we do about it?",
			Options: []string{
				AliasOccupiedNewAlias,
				AliasOccupiedDeleteOld,
				AliasOccupiedGenerateRandom,
				AliasOccupiedNothing,
			},
		}
		wtd := ""
		err = survey.AskOne(sel, &wtd)
		if err != nil {
			return
		}

		switch wtd {
		case AliasOccupiedNothing:
			log.Println("🤬 Alias occupied.")
			return nil
		case AliasOccupiedGenerateRandom:
			return c._actionShortURL(u, "", ex, hideSecret, qrCode)
		case AliasOccupiedNewAlias:
			q := &survey.Input{
				Message: "New alias",
				Default: "",
				Help:    "<empty> = generate",
			}
			err = survey.AskOne(q, &alias)
			if err != nil {
				return
			}
			return c._actionShortURL(u, alias, ex, hideSecret, qrCode)
		case AliasOccupiedDeleteOld:
			for {
				// delete
				input := &survey.Password{
					Message: "Enter Secret",
				}
				secret := ""
				err = survey.AskOne(input, &secret)
				if err != nil {
					return
				}
				// try to delete
				res := c._actionDeleteURL(alias, secret)
				if res != nil {
					input := &survey.Confirm{
						Renderer: survey.Renderer{},
						Message:  "Retry?",
						Default:  true,
					}
					retry := false
					err = survey.AskOne(input, &retry)
					if err != nil {
						return
					}
					if retry {
						continue
					} else {
						return errors.New("alias occupied")
					}
				}

				return c._actionShortURL(u, alias, ex, hideSecret, qrCode)
			}

		default:
			log.Println("wth?!")
			return nil
		}
	}

	return errors.New(s.Message)
}

func (c *CLI) _actionShortURL(u, alias string, ex time.Duration, hideSecret bool, qrCode bool) (err error) {
	// create payload
	payload := &shortreq.CreateShortURLPayload{
		FullURL:            u,
		ExpireAfterSeconds: int(ex.Seconds()),
		PreferredAlias:     short.ShortID(alias),
	}

	// alias?
	var prefix string
	if alias != "" {
		prefix = fmt.Sprintf("🚀 %s -> %s", u, Prefix+alias)
	} else {
		prefix = fmt.Sprintf("🚀 %s", u)
	}
	sp := newSpinner(prefix)
	sp.Start()

	// make request
	var res *req.Resp
	res, err = req.Post(
		ApiUrl+"create",
		req.BodyJSON(payload),
	)
	// stop spinner
	sp.Stop()
	if res == nil {
		return errors.New("response was null")
	}

	// parse response
	s := new(SuccessableCreate)
	err = res.ToJSON(s)
	if err != nil {
		return
	}
	if !s.Success {
		return c._recoverShortURL(s, u, alias, ex, hideSecret, qrCode)
	}

	sh := s.Data
	var secret string
	if hideSecret {
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
	if qrCode {
		config := qrterminal.Config{
			BlackChar: qrterminal.WHITE,
			WhiteChar: qrterminal.BLACK,
			QuietZone: 1,
			Writer:    os.Stdout,
			Level:     qrterminal.L,
		}
		qrterminal.GenerateWithConfig(url, config)
	}

	return nil
}

func (c *CLI) ActionShortURL(ctx *cli.Context) (err error) {
	u := c.FindUrl(ctx)
	if u == "" {
		return errors.New("no url given")
	}

	// parse duration
	ex := ctx.Duration("expire")
	alias := ctx.String("alias")
	hideSecret := ctx.Bool("hide-secret")
	qrCode := ctx.Bool("qr-code")

	return c._actionShortURL(u, alias, ex, hideSecret, qrCode)
}

func repeat(str string, num int) string {
	var builder strings.Builder
	for i := 0; i < num; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
