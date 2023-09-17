package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "leaderboard service"
	cliApp.Version = "1.0.0"

	cliApp.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"start, serve", "s"},
			Usage:   "start server",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}

	fmt.Println("Starting application...")
}
