// Copyright 2025 The AMRP Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MainMenubar(win fyne.Window) *fyne.MainMenu {
	// file menu
	file := fyne.NewMenu("File")

	// tool menu
	tool := fyne.NewMenu("Tool",
		fyne.NewMenuItem("CSV Viewer", func() {
			// note: this feature is excluded from the MVP scope
			dialog.ShowInformation("", "Feature not implemented yet.", win)
		}),
	)

	// help menu
	help := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			const license = "GNU GPLv3"
			version := fmt.Sprintf("%s-%d",
				fyne.CurrentApp().Metadata().Version,
				fyne.CurrentApp().Metadata().Build)
			content := container.NewGridWrap(fyne.NewSize(240, 20),
				widget.NewLabel(fmt.Sprintf("Version: %s", version)),
				widget.NewLabel(fmt.Sprintf("License: %s", license)),
				widget.NewLabel(""),
			)

			d := dialog.NewCustom("About", "Close", content, win)
			d.SetIcon(theme.InfoIcon())
			d.Show()

		}),
	)

	// create the main menu bar
	return fyne.NewMainMenu(
		file,
		tool,
		help,
	)
}
