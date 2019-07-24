package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add local note",
			Action: func(c *cli.Context) error {
				fmt.Println("adding note: ")
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "enc",
					Value: "",
					Usage: "encrypt this note",
				},
				cli.StringFlag{
					Name:  "subject, s",
					Value: "",
					Usage: "Subject of this node",
				},
				cli.StringFlag{
					Name:  "body, b",
					Value: "",
					Usage: "Body of the note",
				},
			},
		},
		{
			Name:    "del",
			Aliases: []string{"d"},
			Usage:   "delete a local note",
			Action: func(c *cli.Context) error {
				fmt.Println("deleting note: ", c.Args())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
