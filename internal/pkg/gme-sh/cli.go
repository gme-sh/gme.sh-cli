package gmesh

const (
	// VERSION -> Muss noch in ne Config
	VERSION = "0.0.1-alpha"
	APIURL  = "http://127.0.0.1:80/"
	PREFIX  = "https://gme.sh/"
)

type CLI struct {
	Pipe string
}

func New() *CLI {
	return &CLI{}
}

func (c *CLI) Run() (err error) {
	// check pipe
	err = c.ReadPipe()
	if err != nil {
		return
	}
	// show help
	return c.ParseArgs()
}
