package commands

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/theiskaa/nt/pkg"
)

// initCommand is a setup command of nt.
var initCommand = &cobra.Command{
	Use:     "init",
	Aliases: []string{"setup"},
	Short:   "Initialize application related files/folders",
	Run:     runInitCommand,
}

// initSetupCommand adds initCommand to main application command.
func initSetupCommand() {
	appCommand.AddCommand(initCommand)
}

// runInitCommand runs appropriate functionalities to setup nt and make it ready-to-use.
func runInitCommand(cmd *cobra.Command, args []string) {
	determineService()

	loading.Start()
	err := service.Init(nil)
	loading.Stop()

	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
		return
	}

	pkg.Alert(pkg.SuccessL, `Application initialized successfully`)
	pkg.Print(" > [nt -h/help] for help", color.FgBlue)
}
