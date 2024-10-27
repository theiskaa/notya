package commands

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/theiskaa/nt/assets"
	"github.com/theiskaa/nt/lib/models"
	"github.com/theiskaa/nt/pkg"
)

var whereCommand = &cobra.Command{
	Use:     "where",
	Aliases: []string{"path", "wh"},
	Short:   "View the path of file or folder",
	Run:     runWhereCommand,
}

func initWhereCommand() {
	appCommand.AddCommand(whereCommand)
}

func runWhereCommand(cmd *cobra.Command, args []string) {
	determineService()

	if len(args) > 0 {
		note, err := service.View(models.Note{Title: args[0]})
		loading.Stop()

		if err != nil {
			pkg.Alert(pkg.ErrorL, err.Error())
		} else {
			pkg.PrintPath((*note).ToNode())
		}

		return
	}

	nodes, noteNames, err := service.GetAll("", "", models.NotyaIgnoreFiles)
	loading.Stop()
	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
		return
	}

	// Ask for note selection.
	var selected string
	survey.AskOne(
		assets.ChooseNodePrompt("note", "view path", noteNames),
		&selected,
	)

	for _, n := range nodes {
		if n.Title == selected {
			pkg.PrintPath(n)
		}
	}
}
