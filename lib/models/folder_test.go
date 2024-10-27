package models_test

import (
	"testing"

	"github.com/theiskaa/nt/lib/models"
	"github.com/theiskaa/nt/lib/services"
)

func TestFolderToNode(t *testing.T) {
	tests := []struct {
		dir      models.Folder
		expected models.Node
	}{
		{
			dir:      models.Folder{},
			expected: models.Node{},
		},
		{
			dir:      models.Folder{Title: "folder/", Path: map[string]string{services.LOCAL.ToStr(): "~/folder"}},
			expected: models.Node{Title: "folder/", Path: map[string]string{services.LOCAL.ToStr(): "~/folder"}},
		},
	}

	for _, td := range tests {
		got := td.dir.ToNode()
		path := got.GetPath(services.LOCAL.ToStr())
		if got.Title != td.expected.Title || path != td.expected.GetPath(services.LOCAL.ToStr()) {
			t.Errorf("Sum was different of [Folder-To-Node] function: Want: %v | Got: %v", td.expected, got)
		}
	}
}
