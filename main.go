package main

import (
    "os"
    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()

    app.Commands = []cli.Command{
        {
            Name: "add",
            Aliases:    []string{"a"},
            Usage:  "add local note",
            Action:  addNote,
            Flags: []cli.Flag{
                cli.StringFlag{
                    Name: "enc",
                    Value: false,
                    Usage: "encrypt this note",
            },
                cli.StringFlag{
                    Name: "subject, s",
                    Value: "",
                    Usage: "Subject of this node",
            },
                cli.StringFlag{
                    Name: "body, b",
                    Value: "",
                    Usage: "Body of the note",
            },

            Name: "del",
            Aliases:    []string{"d"},
            Usage:  "delete a local note",
            Value: "",
            Action:  delNote,
},
}
