// Copyright 2025 The AMRP Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package util

import (
	"math/rand"
)

const NO_FILE_SELECTED = "No file selected..."
const INVALID_CSV_SCHEMA = "Invalid CSV schema detected..."

var AWAIT_MESSAGE = []string{
	"Data is processing... Please wait!",
	"Data is processing... Grab a coffee!",
	"Data is busy... Results coming soon!",
}

func AwaitMessageRandom() string {
	// generate a random number in the range 0, 1, 2
	n := rand.Intn(3)
	return AWAIT_MESSAGE[n]
}
