package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main11() {
	app := cli.NewApp()
	app.Action = action1
	app.Version = "0.0.1"
	app.UsageText = "My Test Cli Project"
	app.Authors = []cli.Author{}
	runCommand := cli.Command{
		Name:      "run",
		ShortName: "r",
		Usage:     "For test",
		Action:    action1,
		HelpName:  "help",
	}
	app.Commands = []cli.Command{
		runCommand,
	}
	app.Run(os.Args)
}

func action1(ctx *cli.Context) {
	fmt.Println("start")
}
