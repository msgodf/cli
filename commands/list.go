package commands

import (
	"github.com/urfave/cli"
)

// ListCommand returns list cli.command
func ListCommand() cli.Command {
	return cli.Command{
		Name:        "list",
		Aliases:     []string{"ls"},
		Usage:       "list command",
		Category:    "MANAGEMENT COMMANDS",
		Hidden:      false,
		ArgsUsage:   "<command>",
		Subcommands: GetCommands(ListCmds),
	}
}
