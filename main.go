package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "goppa",
		Usage: "compare test time.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "previous",
				Aliases: []string{
					"p",
				},
				Usage:    "specify previous test output file.(json)",
				Required: true,
			},
			&cli.StringFlag{
				Name: "current",
				Aliases: []string{
					"c",
				},
				Usage:    "specify current test output file.(json)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println(c.String("previous"), c.String("p"), c.String("current"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
