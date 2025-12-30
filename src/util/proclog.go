// Copyright 2025 The AMRP Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"strings"
)

func ProclogInfo(message ...string) string {
	return fmt.Sprintf("%-9s%-8s%s\n", HHMMSS(), "[INFO ]", strings.Join(message, " "))
}

func ProclogError(err error) string {
	return fmt.Sprintf("%-9s%-8s%s\n", HHMMSS(), "[ERROR]", err.Error())
}
