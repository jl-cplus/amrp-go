// Copyright 2025 The AMRP Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package extension

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewWarning(title, message string, parent fyne.Window) *dialog.CustomDialog {
	content := container.NewVBox(widget.NewLabel(message))
	d := dialog.NewCustom(title, lang.L("OK"), content, parent)
	d.SetIcon(theme.WarningIcon())
	return d
}

func ShowWarning(title, message string, parent fyne.Window) {
	NewWarning(title, message, parent).Show()
}
