// Copyright 2025 The AMRP Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package view

import (
	"amrp-go/util"
	"image/color"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type OpenFileDialog struct {
	Container *fyne.Container
	Label     *widget.Label
	await     bool
	verbose   bool
}

func (w *OpenFileDialog) OpenFileWidget(win fyne.Window, callback func(fpath string, err error)) {
	w.Label = widget.NewLabel(util.NO_FILE_SELECTED)

	rec := canvas.NewRectangle(color.RGBA{R: 233, G: 233, B: 233, A: 255})
	lblBg := container.NewStack(rec, w.Label)

	btn := widget.NewButtonWithIcon("Browse", theme.DocumentIcon(), func() {
		if w.await {
			if w.verbose {
				dialog.ShowInformation("", util.AwaitMessageRandom(), win)
			}
			return
		}

		f := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				callback("", err)
				return
			}

			if reader == nil {
				w.Label.SetText(util.NO_FILE_SELECTED)
				return
			}
			defer reader.Close()

			path := reader.URI().Path()
			// convert to OS-specific separator
			path = filepath.FromSlash(path)

			w.Label.SetText(filepath.Base(path))
			callback(path, nil)
		}, win)
		// TODO: (enhancement) refactor to support configuration in future release
		f.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))
		f.Show()
	})
	btn.Importance = widget.HighImportance

	w.Container = container.New(layout.NewFormLayout(), btn, lblBg)
}

func (w *OpenFileDialog) SetAwait(await bool) {
	w.await = await
}

func (w *OpenFileDialog) SetVerbose(verbose bool) {
	w.verbose = verbose
}

/* // first version
func OpenFileWidget(win fyne.Window, callback func(fpath string, err error)) (*fyne.Container, *widget.Button, *widget.Label) {
	// label
	lbl := widget.NewLabel(NO_FILE_SELECTED)
	//lbl := widget.NewLabel("No file selected...")
	rec := canvas.NewRectangle(color.RGBA{R: 211, G: 211, B: 211, A: 255})
	lblBg := container.NewStack(rec, lbl)
	// button
	btn := widget.NewButtonWithIcon("Browse", theme.DocumentIcon(), func() {
		if of.locked {
			return
		}
		// open file dialog
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			// handle error
			if err != nil {
				callback("", err)
				return
			}
			// cancel, reset label
			if reader == nil {
				lbl.SetText(NO_FILE_SELECTED)
				return
			}
			defer reader.Close()
			path := reader.URI().Path()
			lbl.SetText(filepath.Base(path))
			callback(path, nil)
		}, win)
	})
	btn.Importance = widget.HighImportance
	// form layout
	ctn := container.New(layout.NewFormLayout(),
		btn, lblBg,
	)
	return ctn, btn, lbl
}
*/
