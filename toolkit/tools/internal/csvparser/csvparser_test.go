package csvparser

import (
	"fmt"
	"os"
	"testing"
)

func Test_OutputCSVLog(t *testing.T) {
	wd, _ := os.Getwd()
	fmt.Printf("%s \n", wd)
	files := FilepathsToArray(wd + "/results_test")
	for _, str := range files {
		fmt.Printf("%s \n", str)
	}
	OutputCSVLog(files)
}
