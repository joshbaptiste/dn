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
					Name:  "local, l",
					Value: "",
					Usage: "Local note",
				},
				cli.StringFlag{
					Name:  "prehook",
					Value: "",
					Usage: "Command to run before adding note",
				},
				cli.StringFlag{
					Name:  "posthook",
					Value: "",
					Usage: "Command to run after adding note",
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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "local, l",
					Value: "",
					Usage: "Local note",
				},
				cli.StringFlag{
					Name:  "prehook",
					Value: "",
					Usage: "Command to run before adding note",
				},
				cli.StringFlag{
					Name:  "posthook",
					Value: "",
					Usage: "Command to run after adding note",
				},
			},
		},
		{
			Name:    "edit",
			Aliases: []string{"e"},
			Usage:   "Edit a note",
			Action: func(c *cli.Context) error {
				fmt.Println("Editing note: ", c.Args())
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "local, l",
					Value: "",
					Usage: "Edit local note",
				},
			},
		},
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search a note",
			Action: func(c *cli.Context) error {
				fmt.Println("Searching note: ", c.Args())
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "body, b",
					Value: "",
					Usage: "Search body of note",
				},
				cli.StringFlag{
					Name:  "server, s",
					Value: "",
					Usage: "Search server note",
				},
				cli.StringFlag{
					Name:  "local, l",
					Value: "",
					Usage: "Search local note",
				},
			},
		},
		{
			Name:    "push",
			Aliases: []string{"p"},
			Usage:   "Push a note to a server",
			Action: func(c *cli.Context) error {
				fmt.Println("Pushing note to server: ", c.Args())
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "all",
					Value: "",
					Usage: "Push all notes",
				},
				cli.StringFlag{
					Name:  "clobber",
					Value: "",
					Usage: "Overwrite server note owned by you",
				},
				cli.StringFlag{
					Name:  "server",
					Value: "",
					Usage: "Server ip:[port]",
				},
				cli.StringFlag{
					Name:  "local, l",
					Value: "",
					Usage: "Search local note",
				},
			},
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get a note to a server",
			Action: func(c *cli.Context) error {
				fmt.Println("Getting note from server: ", c.Args())
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "all",
					Value: "",
					Usage: "Get all notes from server",
				},
				cli.StringFlag{
					Name:  "server",
					Value: "",
					Usage: "Server ip:[port]",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
