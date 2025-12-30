// Copyright 2025 The AMRP Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package controller

import (
	"amrp-go/util"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// progress listener for ui to receive updates
type ProgressListener func(num int)

func ProcessCSV(inputPath, outputPath string, progressListener ProgressListener) error {
	// open input file
	inFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// buffer reader
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	lineNum := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// patch a record per line
		if err, line = processInterceptor(line); err != nil {
			return err
		}

		// write to output
		if _, err := writer.WriteString(line); err != nil {
			return err
		}

		// update progress
		lineNum++
		progressListener(lineNum)
	}
	return nil
}

// patching logic
func processInterceptor(line string) (error, string) {
	const size = 653
	schema := []int{5, 28, 10, 10, 6, 2, 30, 140, 40, 140, 17, 17, 17, 6, 10, 12, 25, 25, 25, 25, 25, 17}

	// patch header
	if len(line) < size {
		// validate schema
		if len(strings.Split(line, ",")) != len(schema) {
			return errors.New(util.INVALID_CSV_SCHEMA), ""
		}
		node := strings.Split(strings.TrimSuffix(line, "\n"), ",")
		for i := range node {
			node[i] = strconv.Quote(node[i])
		}
		return nil, strings.Join(node, ",") + "\n"
	}

	// patch data
	str := []string{}
	begin := 0
	for i, length := range schema {
		end := begin + length
		substr := line[begin:end]

		// patch number format
		if i == 10 || i == 11 {
			substr = formatNumberString(substr)
		}

		str = append(str, "\""+substr+"\"")
		begin = end + 1
	}

	return nil, strings.Join(str, ",") + "\n"
}

func formatNumberString(str string) string {
	if strings.HasSuffix(str, "-") {
		str = str[:len(str)-1]
		num := strings.TrimSpace(str)
		space := str[:len(str)-len(num)]
		str = space[:len(space)-1] + "-" + num + " "
	}
	return str
}

func PostOutputInterceptor(err error, outputPath string, callback func(message *string)) {
	if err == nil {
		return
	}

	if err.Error() == util.INVALID_CSV_SCHEMA {
		// remove output
		os.Remove(outputPath)

		msg := fmt.Sprintf("Remove: %s", outputPath)
		callback(&msg)
	}
}
