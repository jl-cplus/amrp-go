// Copyright 2025 The AMRP Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package view

import (
	"amrp-go/controller"
	"amrp-go/util"
	"amrp-go/view/extension"
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MainWindow() {
	app := app.NewWithID("amrp.go")
	app.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})

	win := app.NewWindow("AMRP")

	win.Resize(fyne.NewSize(600, 450))
	win.SetFixedSize(true)
	win.SetMainMenu(MainMenubar(win))
	win.SetContent(mainContent(win))

	win.ShowAndRun()
}

func mainContent(win fyne.Window) *fyne.Container {
	var fpath string = ""
	var proclog []string = []string{}
	var await bool = false
	// reset function
	reset := func(w *OpenFileDialog) {
		fpath = ""
		await = false
		// update on ui thread
		fyne.Do(func() {
			w.SetAwait(false)
			w.Label.SetText(util.NO_FILE_SELECTED)
		})
	}

	// top: openfile widget
	openFile := &OpenFileDialog{}
	openFile.SetVerbose(true)
	openFile.OpenFileWidget(win, func(path string, err error) {
		if err != nil {
			dialog.ShowError(err, win)
		}
		// get file path
		if path != "" {
			fpath = path
		}
	})

	// center: output textarea
	textarea := widget.NewMultiLineEntry()
	textarea.TextStyle = fyne.TextStyle{Monospace: true}
	// widget's disable function causes its text to be unreadable on light themes.
	// Workaround: prevent user data modification by listening to the onchange event
	// and overriding the user input.
	textarea.OnChanged = func(s string) {
		textarea.SetText(strings.Join(proclog, ""))
	}

	// bottom: action button
	actionBtn := widget.NewButtonWithIcon("Patch", theme.MailSendIcon(), func() {
		if await {
			dialog.ShowInformation("", util.AwaitMessageRandom(), win)
			return
		}

		if fpath == "" {
			extension.ShowWarning("", util.NO_FILE_SELECTED, win)
			return
		}

		// generate output filename
		opath := util.RenameAppendSuffix(fpath, "new")
		// update proclog
		proclog = append(proclog, fmt.Sprintf("%-9s%-8s%-8s%s\n", util.HHMMSS(), "[INFO ]", "Source:", fpath))
		proclog = append(proclog, fmt.Sprintf("%-9s%-8s%-8s%s\n", util.HHMMSS(), "[INFO ]", "Output:", opath))
		// update ui
		textarea.SetText(strings.Join(proclog, ""))

		// prevent new processes from starting until the current one finishes
		await = true
		openFile.SetAwait(await)

		go func() {
			// receive progress updates
			progressListener := func(num int) {
				last := proclog[len(proclog)-1]
				if strings.Contains(last, "Patching") {
					proclog = proclog[:len(proclog)-1]
				}
				proclog = append(proclog, fmt.Sprintf("%-9s%-8sPatching total record(s) %d...\n", util.HHMMSS(), "[INFO ]", num))
				// update on ui thread
				fyne.Do(func() {
					textarea.SetText(strings.Join(proclog, ""))
				})
			}
			// start csv path process
			err := controller.ProcessCSV(fpath, opath, progressListener)
			if err != nil {
				dialog.ShowError(err, win)
				// update on ui thread
				fyne.Do(func() {
					proclog = append(proclog, fmt.Sprintf("%-9s%-8s%s\n", util.HHMMSS(), "[ERROR]", err.Error()))
					textarea.SetText(strings.Join(proclog, ""))
				})
			} else {
				// update on ui thread
				fyne.Do(func() {
					proclog = append(proclog, fmt.Sprintf("%-9s%-8s%s\n", util.HHMMSS(), "[INFO ]", "Completed successfully..."))
					textarea.SetText(strings.Join(proclog, ""))
				})
			}
			// post output process
			controller.PostOutputInterceptor(err, opath, func(message *string) {
				if message != nil {
					// update on ui thread
					fyne.Do(func() {
						proclog = append(proclog, fmt.Sprintf("%-9s%-8s%s\n", util.HHMMSS(), "[INFO ]", *message))
						textarea.SetText(strings.Join(proclog, ""))
					})
				}
			})
			// reset
			reset(openFile)
		}()
	})
	actionBtn.Importance = widget.HighImportance

	// border layout
	border := container.NewBorder(
		openFile.Container, // top
		actionBtn,          // bottom
		nil,                // left
		nil,                // right
		textarea,           // center
	)

	return container.NewPadded(border)
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
