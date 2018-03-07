package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Action = action
	app.Version = "0.0.1"
	app.UsageText = "My Test Cli Project"
	app.Authors = []cli.Author{}
	runCommand := cli.Command{
		Name:      "run",
		ShortName: "r",
		Usage:     "For test",
		Action:    action,
		HelpName:  "help",
	}
	app.Commands = []cli.Command{
		runCommand,
	}
	app.Run(os.Args)
}

func action(ctx *cli.Context) {
	fmt.Println("start")
}
