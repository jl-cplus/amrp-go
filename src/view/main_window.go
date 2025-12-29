// Copyright 2025 The AMRP Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func MainWindow() {
	app := app.NewWithID("amrp.go")
	app.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})

	win := app.NewWindow("AMRP")

	win.Resize(fyne.NewSize(600, 450))
	win.SetFixedSize(true)
	win.SetMainMenu(MainMenubar(win))

	win.ShowAndRun()
}

// function below is directly copied from the Fyne demo project.
// Source: https://github.com/fyne-io/demo
// License: BSD 3-Clause "New" or "Revised" License
type forcedVariant struct {
	fyne.Theme
	variant fyne.ThemeVariant
}

func (f *forcedVariant) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return f.Theme.Color(name, f.variant)
}
