// +build !windows

package app

import (
	"github.com/Benbentwo/ui-survey/pkg/cmd"
	"github.com/spf13/cobra"
	"os"
)

// Run runs the command, if args are not nil they will be set on the command
func Get(args []string) *cobra.Command {
	cmd := cmd.NewMainCmd(os.Stdin, os.Stdout, os.Stderr, nil)
	if args != nil {
		args = args[1:]
		cmd.SetArgs(args)
	}
	return cmd
}
