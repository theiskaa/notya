package commands

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/theiskaa/nt/assets"
	"github.com/theiskaa/nt/lib/models"
	"github.com/theiskaa/nt/pkg"
)

// editCommand is a command model which used to overwrite body of notes or files.
var editCommand = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"overwrite", "update"},
	Short:   "Edit/Update note data",
	Run:     runEditCommand,
}

// initEditCommand adds editCommand to main application command.
func initEditCommand() {
	appCommand.AddCommand(editCommand)
}

// runEditCommand runs appropriate service commands to edit/overwrite note data.
func runEditCommand(cmd *cobra.Command, args []string) {
	determineService()

	// Take note title from arguments. If it's provided.
	if len(args) > 0 {
		editAndFinish(models.Node{Title: args[0]})
		return
	}

	// Generate all node names.
	loading.Start()
	_, nodeNames, err := service.GetAll("", "file", models.NotyaIgnoreFiles)
	loading.Stop()
	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
		return
	}

	// Ask for note selection.
	var selected string
	survey.AskOne(
		assets.ChooseNodePrompt("note", "edit", nodeNames),
		&selected,
	)

	// Open selected note-file.
	editAndFinish(models.Node{Title: selected})
}

func editAndFinish(note models.Node) {
	if len(note.Title) == 0 {
		os.Exit(-1)
		return
	}

	if err := service.Open(note); err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
	}
}
