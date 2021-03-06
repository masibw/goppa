package main

import (
	"fmt"
	"github.com/masibw/goppa/infrastructure/loader"
	"github.com/masibw/goppa/usecase"
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
				Usage:    "specify previous test output file.(with -v option)",
				Required: false,
			},
			&cli.StringFlag{
				Name: "current",
				Aliases: []string{
					"c",
				},
				Usage:    "specify current test output file.(with -v option)",
				Required: true,
			},
			&cli.Float64Flag{
				Name: "border",
				Aliases: []string{
					"b",
				},
				Usage:    "specify how many times slower than before to be detected.",
				Required: false,
				Value:    1.5,
			},
		},
		Action: func(c *cli.Context) error {
			l := loader.NewVerboseLoader()
			diff := usecase.CompareWithPrev(c.String("previous"), c.String("current"), l, c.Float64("border"))

			if diff == nil {
				return nil
			}
			fmt.Println("some tests are slower than previous.")
			for _, output := range diff {
				fmt.Println(output)
			}
			return cli.Exit("", 1)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
