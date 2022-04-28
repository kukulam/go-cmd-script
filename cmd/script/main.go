package main

import (
	"go-cmd-script/internal/script"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
logger := script.NewLogger()

app := cli.NewApp()
	app.Name = "go-cmd-script"
	app.Usage = ""
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "name",
			Usage:    "usage",
			Required: false,
		},
	}
	app.Action = func (c *cli.Context) error {
		return script.Command(c, logger)
	}
	
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
