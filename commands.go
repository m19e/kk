package main

import (
	"fmt"
	"os"

	"github.com/m19e/kk/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []*cli.Command{
	{
		Name:   "read",
		Usage:  "",
		Action: command.CmdRead,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "write",
		Usage:  "",
		Action: command.CmdWrite,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "list",
		Usage:  "",
		Action: command.CmdList,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
