// Copyright Microsoft Corporation.
// Licensed under the MIT License.

package timestamp

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

//TestMain found in configuration_test.go.

func Test_WritetoFile_range_instant(t *testing.T) {
	TrackToFile(time.Now(), "test tool", "test step", true, os.Stdout)
}

func Test_WritetoFile_noRange_instant(t *testing.T) {
	TrackToFile(time.Now(), "test tool", "test step", false, os.Stdout)
}

func Test_WritetoCSV_range_instant(t *testing.T) {
	TrackToCSV(time.Now(), "test tool", "test step", true)
}

func Test_WritetoCSV_noRange_instant(t *testing.T) {
	TrackToCSV(time.Now(), "test tool", "test step", false)
}

func Test_WritetoFile_range_sleeps(t *testing.T) {
	defer TrackToFile(time.Now(), "tool 1", "step 1", true, os.Stdout)
	time.Sleep(3 * time.Second)
}

func Test_WritetoCSV_range_sleeps(t *testing.T) {
	defer TrackToCSV(time.Now(), "test tool", "test step", true)
	time.Sleep(3 * time.Second)
}
func NumberOfLines() int {
	file, _ := os.Open("build-time.csv")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount
}

//	Run debug test to see print output in debug console.
func Test_WritetoCSV_threeTimes(t *testing.T) {
	oldLines := NumberOfLines()
	for i := 0; i < 3; i++ {
		TrackToCSV(time.Now(), "test tool", "test step", true)
	}
	newLines := NumberOfLines() - oldLines
	fmt.Println("Number of new lines:", newLines)
	if newLines != 3 {
		t.Fail()
	}
}

//	Tests between 20 to 40 times
func Test_WritetoCSV_nTimes(t *testing.T) {
	numTests := rand.Intn(20) + 20
	oldLines := NumberOfLines()
	for i := 0; i < numTests; i++ {
		TrackToCSV(time.Now(), "test tool", "test step", true)
	}
	newLines := NumberOfLines() - oldLines
	fmt.Println("Number of new lines:", newLines)
	if newLines != numTests {
		t.Fail()
	}
}
