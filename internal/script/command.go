package script

import "github.com/urfave/cli"

func Command(c *cli.Context, l *Logger) error {
	l.debug("Hello world")
	return nil
}