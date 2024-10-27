package cmd

import (
	"github.com/theiskaa/nt/lib/commands"
)

// RunApp executes appCommand.
// It'd be happen only once, on starting program at [main.go].
func RunApp() {
	commands.ExecuteApp()
}
