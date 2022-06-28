// Copyright Microsoft Corporation.
// Licensed under the MIT License.

package timestamp

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
func WritetoCSV(seconds time.Duration) {
	defer TrackToCSV(time.Now(), "test tool", "test step", true)
	time.Sleep(seconds * time.Second)
}

func Test_WritetoCSV_Delay(t *testing.T) {
	WritetoCSV(0)
	WritetoCSV(1)
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

func WritetoCSV_MultipleLines(count int, t *testing.T) {
	oldLines := NumberOfLines()
	for i := 0; i < count; i++ {
		WritetoCSV(0)
	}
	newLines := NumberOfLines() - oldLines
	assert.Equal(newLines, count)
}

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

func WritetoCSV_timingTest(time time.Duration, t *testing.T) {
	WritetoCSV(time)
	latestTimestamp := GetLatestTimestamp()
	data := strings.Split(latestTimestamp, ",")
	match, err := regexp.MatchString("1.[0-9]{9}s", data[2]) // TODO: Make the timing test work for non-three second intervals
	assert.NoError(err)
	assert.True(match)
}

// func GetLatestTimestamp() string {
// 	file, _ := os.Open("build-time.csv")
// 	fileScanner := bufio.NewScanner(file)
// 	lastLine := ""
// 	for fileScanner.Scan() {
// 		lastLine = fileScanner.Text()
// 	}
// 	return lastLine
// }

func Test_WritetoCSV_formatTest(t *testing.T) {
	WritetoCSV(0)
	latestTimestamp := GetLatestTimestamp()
	data := strings.Split(latestTimestamp, ",")
	exp := [7]string{
		".+",
		".+",
		"[0-9]+[.][0-9]+[(µs)(s)]",
		"[A-Za-z]{3}",
		"[0-9]{2}\\s[A-Za-z]{3}\\s[0-9]{4}\\s[0-9]{2}[:][0-9]{2}[:][0-9]{2}\\s[A-Z]{3}",
		"[A-Za-z]{3}",
		"[0-9]{2}\\s[A-Za-z]{3}\\s[0-9]{4}\\s[0-9]{2}[:][0-9]{2}[:][0-9]{2}\\s[A-Z]{3}",
	}
	for i := 0; i < 7; i++ {
		match, err := regexp.MatchString(exp[i], data[i])
		assert.NoError(err)
		assert.True(match)
	}
}
