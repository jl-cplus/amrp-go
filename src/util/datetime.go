// Copyright 2025 The AMRP Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package util

import (
	"time"
)

func HHMMSS() string {
	return time.Now().Format("15:04:05")
}
