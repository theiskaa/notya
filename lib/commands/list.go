package commands

import (
	"github.com/spf13/cobra"
	"github.com/theiskaa/nt/lib/models"
	"github.com/theiskaa/nt/pkg"
)

// listCommand is a command that used to list all exiting nodes.
var listCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all nt nodes(files & folders)",
	Run:     runListCommand,
}

// initListCommand adds listCommand to main application command.
func initListCommand() {
	appCommand.AddCommand(listCommand)
}

// runListCommand runs appropriate service functionalities to log all nodes.
func runListCommand(cmd *cobra.Command, args []string) {
	determineService()

	var additional string
	if len(args) > 0 {
		additional = args[0]
	}

	loading.Start()

	// Generate a list of nodes.
	nodes, _, err := service.GetAll(additional, "", models.NotyaIgnoreFiles)

	loading.Stop()
	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
		return
	}

	pkg.PrintNodes(nodes)
}
