// Copyright Microsoft Corporation.
// Licensed under the MIT License.

package timestamp

import (
	"os"
	"regexp"
	"strings"
	"testing"
	"time"
)

var (
	info1 = New("tool 1", true)
	info2 = New("tool 2", false)
)

//TestMain found in configuration_test.go.

func Test_WritetoFile_range_instant(t *testing.T) {
	info1.Start()
	info1.RecordToFile("test step", "test action", os.Stdout)
}

func Test_WritetoFile_noRange_instant(t *testing.T) {
	info2.Start()
	time.Sleep(3 * time.Second)
	info2.RecordToFile("test step", "test action", os.Stdout)
}

func Test_WritetoCSV_range(t *testing.T) {
	info1.InitCSV("build-time")
	info1.Start()
	time.Sleep(1 * time.Second)
	info1.RecordToCSV("step 1", "action 1")
	info1.Start()
	time.Sleep(2 * time.Second)
	info1.RecordToCSV("step 2", "action 1")
	info1.Start()
	time.Sleep(1 * time.Second)
	info1.RecordToCSV("step 2", "action 2")
	info1.Start()
	time.Sleep(3 * time.Second)
	info1.RecordToCSV("step 3", "action 1")
}

func Test_WritetoCSV_noRange(t *testing.T) {
	info2.InitCSV("build-time")
	time.Sleep(1 * time.Second) // extra sleep
	// info2.Start()
	time.Sleep(2 * time.Second)
	info2.RecordToCSV("step 1", "action 1")
	info2.Start()
	time.Sleep(2 * time.Second)
	info2.RecordToCSV("step 2", "action 1")
	time.Sleep(1 * time.Second) // extra sleep
	// info2.Start()
	time.Sleep(2 * time.Second)
	info2.RecordToCSV("step 2", "action 2")
	info2.Start()
	time.Sleep(2 * time.Second)
	info2.RecordToCSV("step 3", "action 1")
}

func Test_WritetoFile_noRange(t *testing.T) {
	info2.Start()
	time.Sleep(1 * time.Second)
	info2.RecordToFile("step 1", "action 1", os.Stdout)
	info2.Start()
	time.Sleep(1 * time.Second)
	info2.RecordToFile("step 2", "action 1", os.Stdout)
	time.Sleep(2 * time.Second)
	info2.Start()
	time.Sleep(1 * time.Second)
	info2.RecordToFile("step 2", "action 2", os.Stdout)
	info2.Start()
	time.Sleep(1 * time.Second)
	info2.RecordToFile("step 3", "action 1", os.Stdout)
}

// func Test_WritetoFile_range_sleeps(t *testing.T) {
// 	defer TrackToFile(time.Now(), "tool 1", "step 1", true, os.Stdout)
// 	time.Sleep(3 * time.Second)
// }

// func Test_WritetoCSV_range_sleeps(t *testing.T) {
// 	defer TrackToCSV(time.Now(), "test tool", "test step", true)
// 	time.Sleep(3 * time.Second)
// }

// func NumberOfLines() int {
// 	file, _ := os.Open("build-time.csv")
// 	fileScanner := bufio.NewScanner(file)
// 	lineCount := 0
// 	for fileScanner.Scan() {
// 		lineCount++
// 	}
// 	return lineCount
// }

// //	Run debug test to see print output in debug console.
// func Test_WritetoCSV_threeTimes(t *testing.T) {
// 	oldLines := NumberOfLines()
// 	for i := 0; i < 3; i++ {
// 		TrackToCSV(time.Now(), "test tool", "test step", true)
// 	}
// 	newLines := NumberOfLines() - oldLines
// 	fmt.Println("Number of new lines:", newLines)
// 	if newLines != 3 {
// 		t.Fail()
// 	}
// }

// //	Tests between 20 to 40 times
// func Test_WritetoCSV_nTimes(t *testing.T) {
// 	rand.Seed(time.Now().UnixNano())
// 	numTests := rand.Intn(21) + 20
// 	oldLines := NumberOfLines()
// 	for i := 0; i < numTests; i++ {
// 		TrackToCSV(time.Now(), "test tool", "test step", true)
// 	}
// 	newLines := NumberOfLines() - oldLines
// 	fmt.Println("Number of new lines:", newLines)
// 	if newLines != numTests {
// 		t.Fail()
// 	}
// }

// func GetLatestTimestamp() string {
// 	file, _ := os.Open("build-time.csv")
// 	fileScanner := bufio.NewScanner(file)
// 	lastLine := ""
// 	for fileScanner.Scan() {
// 		lastLine = fileScanner.Text()
// 	}
// 	return lastLine
// }

// func Test_WritetoCSV_timingTest(t *testing.T) {
// 	fmt.Println(GetLatestTimestamp())
// }
