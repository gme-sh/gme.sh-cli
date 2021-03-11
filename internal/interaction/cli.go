package interaction

import (
	"github.com/briandowns/spinner"
	"github.com/full-stack-gods/gme.sh-cli/internal/api"
	"github.com/full-stack-gods/gme.sh-cli/internal/config"
	"time"
)

const (
	// Version -> Muss noch in ne Config
	Version = "0.1.0-alpha"
)

type CLI struct {
	Pipe   string
	API    *api.API
	Config *config.Config
}

func New(api *api.API, cfg *config.Config) *CLI {
	return &CLI{
		API:    api,
		Config: cfg,
	}
}

func (c *CLI) Run() (err error) {
	// check pipe
	err = c.ReadPipe()
	if err != nil {
		return
	}
	// show help
	return c.RunApp()
}

func newSpinner(message string) (sp *spinner.Spinner) {
	sp = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	sp.Prefix = message + " ["
	sp.Suffix = "]"
	sp.FinalMSG = sp.Prefix + "done" + sp.Suffix + "\n"
	return
}
