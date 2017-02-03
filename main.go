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
                    Name: "dec",
                    Value: false,
                    Usage: "decrypt this note",
            },
                cli.BoolFlag{
                    Name: "pass, p",
                    Usage: "password for [enc|dec]cryption",
            },

            Name: "del",
            Aliases:    []string{"d"},
            Usage:  "del local note",
            Action:  delNote,
            Flags: []cli.Flag{
                cli.StringFlag{
                    Name: "enc",
                    Value: false,
                    Usage: "encrypt this note",
            },
                cli.StringFlag{
                    Name: "dec",
                    Value: false,
                    Usage: "decrypt this note",
            },
                cli.BoolFlag{
                    Name: "pass, p",
                    Usage: "password for [enc|dec]cryption",
            },
            Name: "del"
},
}
