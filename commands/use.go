package commands

import (
	"github.com/urfave/cli"
)

// UseCommand returns use cli.command
func UseCommand() cli.Command {
	return cli.Command{
		Name:        "use",
		Aliases:     []string{"u"},
		Usage:       "use command",
		Category:    "MANAGEMENT COMMANDS",
		Hidden:      false,
		ArgsUsage:   "<command>",
		Subcommands: GetCommands(UseCmds),
	}
}
