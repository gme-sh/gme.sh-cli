package gmesh

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func (c *CLI) ReadPipe() error {
	info, err := os.Stdin.Stat()
	if err != nil {
		return err
	}
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return nil
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	c.Pipe = strings.TrimSpace(string(output))

	return nil
}
