package interaction

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/full-stack-gods/gme.sh-cli/internal/api"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/short"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/shortreq"
	"github.com/mdp/qrterminal"
	"github.com/mgutz/ansi"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

const (
	AliasOccupiedNewAlias       = "Enter new alias"
	AliasOccupiedDeleteOld      = "Delete old alias"
	AliasOccupiedGenerateRandom = "Generate random alias"
	AliasOccupiedNothing        = "Nothing"
)

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

func (c *CLI) _actionShortURL(u, alias string, ex time.Duration, hideSecret bool, qrCode bool) (err error) {
	// create payload
	payload := &shortreq.CreateShortURLPayload{
		FullURL:            u,
		ExpireAfterSeconds: int(ex.Seconds()),
		PreferredAlias:     short.ShortID(alias),
	}

	prefix := orElse(alias == "",
		fmt.Sprintf("🚀 %s", u),
		fmt.Sprintf("🚀 %s -> %s", u, c.API.GetURL(alias)))

	// make request
	sp := newSpinner(prefix)
	sp.Start()
	s, err := c.API.CreateShortURL(payload)
	sp.Stop()
	if err != nil {
		return err
	}

	// if s.Success is false, try to recover
	if !s.Success {
		return c._recoverShortURL(s, u, alias, ex, hideSecret, qrCode)
	}

	sh := s.Data
	secret := orElse(hideSecret,
		repeat("*", len(sh.Secret)),
		sh.Secret)

	// output url
	url := c.API.GetURL(sh.ID.String())
	fmt.Println("🔗", ansi.Cyan+url, ansi.LightGreen+"[Secret: "+secret+"]", ansi.Reset)

	// output expiration
	if ex != 0 {
		fmt.Println("⏰", "Expires at", time.Now().Add(ex).Format("02.01.2006 15:04:05"))
	}

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

func (c *CLI) _recoverShortURL(s *api.SuccessableCreate, u, alias string, ex time.Duration, hideSecret bool, qrCode bool) (err error) {
	switch s.Code {
	case shortreq.ResponseErrInvalidURL.InternalCode:
		fmt.Println(ansi.Red+"ERROR:", ansi.White+"Invalid URL", ansi.Reset)
		return nil

	case shortreq.ResponseErrDomainBlocked.InternalCode:
		fmt.Println(ansi.Red+"ERROR:", ansi.White+"Domain is blocked.", ansi.Reset)
		fmt.Println(">", ansi.Yellow+s.Message)
		return nil

	case shortreq.ResponseErrDatabaseSave.InternalCode:
	case shortreq.ResponseErrGeneratedAliasNotAvailable.InternalCode:
		fmt.Println(ansi.Red+"ERROR:", ansi.White+"Internal server error. Try again.", ansi.Reset)
		fmt.Println(">", ansi.Yellow+s.Message)
		return nil

	case shortreq.ResponseErrAliasOccupied.InternalCode:
		fmt.Println(ansi.Red+"ERROR:", ansi.White+"Alias occupied.", ansi.Reset)

		var wtd string
		sel := &survey.Select{
			Message: "What do we do about it?",
			Options: []string{
				AliasOccupiedNewAlias,
				AliasOccupiedDeleteOld,
				AliasOccupiedGenerateRandom,
				AliasOccupiedNothing,
			},
		}
		if err = survey.AskOne(sel, &wtd); err != nil {
			return
		}

		switch wtd {
		case AliasOccupiedNothing:
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
						Message: "Retry?",
						Default: true,
					}
					retry := false
					err = survey.AskOne(input, &retry)
					if err != nil {
						return
					}

					if retry {
						continue // ask again for an alias
					}
					return errors.New("alias occupied")
				}

				// after it was successfully deleted, create new short url
				return c._actionShortURL(u, alias, ex, hideSecret, qrCode)
			}

		default:
			log.Println("wth?!")
			return nil
		}

	case shortreq.ResponseOkCreate.InternalCode:
		fmt.Println("Nothing to recover ¯\\_(ツ)_/¯")
		return nil
	}

	return errors.New(s.Message)
}
