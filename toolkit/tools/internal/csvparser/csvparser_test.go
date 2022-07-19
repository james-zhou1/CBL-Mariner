package csvparser

import (
	"os"
	"testing"
)

func Test_OutputCSVLog(t *testing.T) {
	wd, _ := os.Getwd()
	OutputCSVLog(wd + "/results_test")
}
