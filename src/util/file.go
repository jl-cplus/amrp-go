// Copyright 2025 The AMRP Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"path/filepath"
)

func RenameAppendSuffix(file string, suffix string) string {
	fpath := filepath.Dir(file)  // directory part
	fname := filepath.Base(file) // filename with extension
	fext := filepath.Ext(fname)  // extension (".csv")

	name := fname[:len(fname)-len(fext)]               // filename without extension
	name = fmt.Sprintf("%s(%s)%s", name, suffix, fext) // add suffix before extension

	return filepath.Join(fpath, name) // rebuild full path
}
