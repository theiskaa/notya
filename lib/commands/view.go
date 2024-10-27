package commands

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/theiskaa/nt/assets"
	"github.com/theiskaa/nt/lib/models"
	"github.com/theiskaa/nt/pkg"
)

// viewCommand is a command model which used to view metadata of note.
var viewCommand = &cobra.Command{
	Use:     "view",
	Aliases: []string{"show", "read"},
	Short:   "View full note data",
	Run:     runViewCommand,
}

// initViewCommand adds viewCommand to main application command.
func initViewCommand() {
	appCommand.AddCommand(viewCommand)
}

// runViewCommand runs appropriate service commands to log full note data.
func runViewCommand(cmd *cobra.Command, args []string) {
	determineService()

	loading.Start()

	// Take note title from arguments. If it's provided.
	if len(args) > 0 {
		note, err := service.View(models.Note{Title: args[0]})
		loading.Stop()

		if err != nil {
			pkg.Alert(pkg.ErrorL, err.Error())
		} else {
			pkg.PrintNote(*note, service.Type())
		}

		return
	}

	// Generate array of all note names.
	nodes, noteNames, err := service.GetAll("", "file", models.NotyaIgnoreFiles)
	loading.Stop()
	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
		return
	}

	// Ask for note selection.
	var selected string
	survey.AskOne(
		assets.ChooseNodePrompt("note", "view", noteNames),
		&selected,
	)

	for _, n := range nodes {
		if n.Title == selected {
			pkg.PrintNote(n.ToNote(), service.Type())
		}
	}
}
