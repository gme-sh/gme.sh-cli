package interaction

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/atotto/clipboard"
	"github.com/full-stack-gods/gme.sh-cli/internal/api"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/short"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/shortreq"
	"github.com/urfave/cli/v2"
	"time"
)

func (c *CLI) ActionWatchPool(ctx *cli.Context) (err error) {
	var (
		secret string
		poolID short.PoolID
		name   string
	)
	// read name
	if n := ctx.String("name"); n != "" {
		name = n
	} else {
		q := &survey.Input{
			Message: "Computer-Name",
		}
		if err = survey.AskOne(q, &name); err != nil {
			return
		}
	}
	// read pool id
	if i := ctx.String("pool-id"); i != "" {
		poolID = short.PoolID(i)
	} else {
		var poolIDstr string
		q := &survey.Input{
			Message: "Pool-ID",
		}
		if err = survey.AskOne(q, &poolIDstr); err != nil {
			return
		}
		poolID = short.PoolID(poolIDstr)
	}
	// read secret
	if s := ctx.String("secret"); s != "" {
		secret = s
	} else {
		// ask for secret
		q := &survey.Password{
			Message: "Pool-Secret",
		}
		if err = survey.AskOne(q, &secret); err != nil {
			return
		}
	}
	// try to fetch pool
	var s *api.SuccessablePool
	if s, err = c.API.GetPool(&poolID, secret); err != nil {
		return
	}
	if !s.Success {
		return errors.New(s.Message)
	}

	pool := s.Pool
	fmt.Println("🏊 Found pool:", pool.ID)
	fmt.Println()
	fmt.Println("-- Entries for", name, "--")
	found := false
	if pool.Entries != nil {
		if entries, ok := pool.Entries[name]; ok {
			found = ok
			fmt.Println(len(entries), "entries ...")
			for _, e := range entries {
				fmt.Println("  *", e.Time.Format("02.01.2006 15:04:05"), "->", e.URL)
			}
		}
	}
	if !found {
		fmt.Println("😭 No Entries")
	}
	fmt.Println("--")

	// watch
	if !ctx.Bool("watch") {
		return
	}

	fmt.Println("👉 Watching Clipboard. Press CTRL-C to stop.")

	// start watching
	var oldClipboard string
	for {
		time.Sleep(500 * time.Millisecond)
		text, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("🧐 ERROR:", err)
			continue
		}
		if oldClipboard == text {
			continue
		}
		oldClipboard = text

		// check if url
		if !shortreq.UrlRegex.MatchString(text) {
			continue
		}

		fmt.Print("NEW: ", text, " 🚀")

		s := new(shortreq.Successable)
		var e error
		if s, e = c.API.AppendPool(&poolID, name, secret, text); e != nil {
			fmt.Println(" ... Errored:", e)
			continue
		}

		if !s.Success {
			fmt.Println(" ...  Unsuccessful:", s.Message)
			continue
		}

		fmt.Println(" ... Success!")
	}
}
