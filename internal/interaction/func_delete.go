package interaction

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/shortreq"
	"github.com/mgutz/ansi"
	"github.com/urfave/cli/v2"
)

const (
	DeleteReenterSecret = "Re-enter secret"
	DeleteRetry         = "Retry"
	DeleteNothing       = "Nothing"
)

func (c *CLI) ActionDeleteURL(ctx *cli.Context) (err error) {
	u := c.FindUrl(ctx)
	s := ctx.String("secret")
	err = c._actionDeleteURL(u, s)
	return
}

func (c *CLI) _actionDeleteURL(u, sec string) (err error) {
	if u == "" {
		return errors.New("no url provided")
	} else if sec == "" {
		return errors.New("no secret provided")
	}
	s, err := c.API.DeleteShortURL(u, sec)
	if err != nil {
		return err
	}
	if !s.Success {
		fmt.Println(ansi.Red+"ERROR:", ansi.White+s.Message, ansi.Reset)
		return c._recoverDeleteURL(s, u, sec)
	}
	fmt.Println(ansi.Green+"OKAY:", ansi.White+"Deleted Alias", ansi.Reset)
	return nil
}

func (c *CLI) _recoverDeleteURL(s *shortreq.Successable, u, sec string) (err error) {
	switch s.Code {

	case shortreq.ResponseErrLocked.InternalCode:
	case shortreq.ResponseErrURLNotFound.InternalCode:
	case shortreq.ResponseErrEmptyID.InternalCode:
		fmt.Println(ansi.Red+"ERROR:", ansi.White+s.Message, ansi.Reset)
		return nil

	case shortreq.ResponseErrSecretMismatch.InternalCode:
		// ask to retry
		var answer string
		q := &survey.Select{
			Message: "What do we do about it?",
			Options: []string{
				DeleteReenterSecret,
				DeleteRetry,
				DeleteNothing,
			},
		}
		if err = survey.AskOne(q, &answer); err != nil {
			return err
		}

		switch answer {
		case DeleteReenterSecret:
			// ask for new secret
			var secret string
			p := &survey.Password{
				Message: "Enter secret",
			}
			if err = survey.AskOne(p, &secret); err != nil {
				return err
			}
			return c._actionDeleteURL(u, secret)

		case DeleteRetry:
			return c._actionDeleteURL(u, sec)

		case DeleteNothing:
			return nil

		default:
			return errors.New("wth")
		}

	default:
		return errors.New("nothing to recover")
	}

	return nil
}
