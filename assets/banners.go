package assets

import "fmt"

// ShortSlog is a main slogan of nt.
var ShortSlog = "📝 Take notes quickly and expeditiously"

// MinimalisticBanner is a first banner of nt.
var MinimalisticBanner = `
  _   _   _____
 | \ | | |_   _|
 |  \| |   | |
 | |\  |   | |
 |_| \_|   |_|

`

// GenerateBanner merges slog and banner together, to get final result of application banner.
func GenerateBanner(banner string, slog string) string {
	template := `
  %v
 %v
   `

	return fmt.Sprintf(template, banner, slog)
}
