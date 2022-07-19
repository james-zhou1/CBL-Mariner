// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Test cases for csvparser.go.

package csvparser

import (
	"os"
	"testing"
)

func Test_OutputCSVLog(t *testing.T) {
	wd, _ := os.Getwd()
	OutputCSVLog(wd + "/results_test")
}
