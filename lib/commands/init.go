// Copyright 2021-present Anon. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package commands

import (
	"github.com/anonistas/notya/pkg"
	"github.com/spf13/cobra"
)

// initCommand is a setup command of notya.
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

// runInitCommand runs appropriate functionalities to setup notya and make it ready-to-use.
func runInitCommand(cmd *cobra.Command, args []string) {
	err := service.Init()
	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
		return
	}

	pkg.Alert(pkg.SuccessL, "Notya initializing completed successfully, It's ready to use now!")
}

// initializeIfNotExists is called on each main command execution.
// Checks if notya initialized or not, if not then initializes it automatically.
// If notya initialized already, returns nothing.
func initializeIfNotExists() error {
	notyaPath, err := pkg.NotyaPWD()
	if err != nil {
		return err
	}

	if !pkg.FileExists(*notyaPath) {
		_ = service.Init()
	}

	return nil
}
