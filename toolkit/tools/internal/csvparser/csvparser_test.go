package csvparser

import "testing"

func Test_OutputCSVLog(t *testing.T) {
	files := FilepathsToArrayTest()
	OutputCSVLog(files)
}
