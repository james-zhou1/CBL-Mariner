// Copyright Microsoft Corporation.
// Licensed under the MIT License.

package timestamp

import (
	"fmt"
	"os"

	// "regexp"
	// "strings"
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
	time.Sleep(30 * time.Millisecond)
	info2.RecordToFile("test step", "test action", os.Stdout)
}

func Test_WritetoCSV_range(t *testing.T) {
	InitCSV("build-time", true)
	Stamp.Start()
	time.Sleep(10 * time.Millisecond)
	Stamp.RecordToCSV("step 1", "action 1")
	Stamp.Start()
	time.Sleep(20 * time.Millisecond)
	Stamp.RecordToCSV("step 2", "action 1")
	Stamp.Start()
	time.Sleep(10 * time.Millisecond)
	Stamp.RecordToCSV("step 2", "action 2")
	Stamp.Start()
	time.Sleep(30 * time.Millisecond)
	Stamp.RecordToCSV("step 3", "action 1")
}

func Test_getHomeDir(t * testing.T) {
	home, _ := os.UserHomeDir()
	fmt.Printf("%s\n", home)
	curr, _ := os.Getwd()
	fmt.Printf("%s\n", curr)
	testDir := "/home/xuanchen/repos/pod_repo/CBL-Mariner/toolkit"
	dirLen := len(testDir)
	fmt.Printf("%s\n", testDir[dirLen - 19: ])
}

func Test_WritetoCSV_noRange(t *testing.T) {
	InitCSV("build-time", false)
	time.Sleep(10 * time.Millisecond) // extra sleep
	// info2.Start()
	time.Sleep(20 * time.Millisecond)
	Stamp.RecordToCSV("step 1", "action 1")
	Stamp.Start()
	time.Sleep(20 * time.Millisecond)
	Stamp.RecordToCSV("step 2", "action 1")
	time.Sleep(10 * time.Millisecond) // extra sleep
	// info2.Start()
	time.Sleep(20 * time.Millisecond)
	Stamp.RecordToCSV("step 2", "action 2")
	Stamp.Start()
	time.Sleep(20 * time.Millisecond)
	Stamp.RecordToCSV("step 3", "action 1")
}

func Test_WritetoFile_noRange(t *testing.T) {
	info2.Start()
	time.Sleep(10 * time.Millisecond)
	info2.RecordToFile("step 1", "action 1", os.Stdout)
	info2.Start()
	time.Sleep(10 * time.Millisecond)
	info2.RecordToFile("step 2", "action 1", os.Stdout)
	time.Sleep(20 * time.Millisecond)
	info2.Start()
	time.Sleep(10 * time.Millisecond)
	info2.RecordToFile("step 2", "action 2", os.Stdout)
	info2.Start()
	time.Sleep(10 * time.Millisecond)
	info2.RecordToFile("step 3", "action 1", os.Stdout)
}

// func Test_roast(t *testing.T) {
// 	// info1.InitCSV("toolkit/tools/internal/timestamp/results/roast_test")
// 	info1.InitCSV("roast_test")
// 	info1.RecordToCSV("step", "action")
// }

func WritetoCSV(info *TimeInfo, seconds time.Duration) {
	Stamp.Start()
	time.Sleep(seconds * time.Millisecond)
	Stamp.RecordToCSV("test tool", "test step")
}

// func Test_WritetoCSV_Delay(t *testing.T) {
// 	info1.InitCSV("build-time")
// 	WritetoCSV(info1, 0)
// 	WritetoCSV(info1, 1)
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

// func WritetoCSV_MultipleLines(count int, t *testing.T) {
// 	oldLines := NumberOfLines()
// 	for i := 0; i < count; i++ {
// 		WritetoCSV(info1, 0)
// 	}
// 	newLines := NumberOfLines() - oldLines
// 	assert.Equal(t, newLines, count)
// }

// // //	Run debug test to see print output in debug console.
// func Test_WritetoCSV_MultipleLines(t *testing.T) {
// 	info1.InitCSV("build-time")
// 	WritetoCSV_MultipleLines(1, t)
// }

// func WritetoCSV_timingTest(time time.Duration, t *testing.T) {
// 	WritetoCSV(time)
// 	latestTimestamp := GetLatestTimestamp()
// 	data := strings.Split(latestTimestamp, ",")
// 	match, err := regexp.MatchString("1.[0-9]{9}s", data[2]) // TODO: Make the timing test work for non-three second intervals
// 	assert.NoError(err)
// 	assert.True(match)
// }

// func Test_WritetoCSV_timingTest(t *testing.T) {
// 	WritetoCSV_timingTest(1, t)
// }

// func Test_WritetoCSV_formatTest(t *testing.T) {
// 	WritetoCSV(0)
// func Test_WritetoCSV_Delay(t *testing.T) {
// 	writetoCSV(0)
// 	writetoCSV(1)
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

// func WritetoCSV_MultipleLines(count int, t *testing.T) {
// 	oldLines := NumberOfLines()
// 	for i := 0; i < count; i++ {
// 		writetoCSV(0)
// 	}
// 	newLines := NumberOfLines() - oldLines
// 	assert.Equal(newLines, count)
// }

// //	Run debug test to see print output in debug console.
// func Test_WritetoCSV_MultipleLines(t *testing.T) {
// 	WritetoCSV_MultipleLines(1, t)
// }

// func WritetoCSV_timingTest(time time.Duration, t *testing.T) {
// 	writetoCSV(time)
// 	latestTimestamp := GetLatestTimestamp()
// 	data := strings.Split(latestTimestamp, ",")
// 	match, err := regexp.MatchString("1.[0-9]{9}s", data[2]) // TODO: Make the timing test work for non-three second intervals
// 	assert.NoError(err)
// 	assert.True(match)
// }

// func Test_WritetoCSV_timingTest(t *testing.T) {
// 	WritetoCSV_timingTest(1, t)
// }

// func Test_WritetoCSV_formatTest(t *testing.T) {
// 	writetoCSV(0)
// 	latestTimestamp := GetLatestTimestamp()
// 	data := strings.Split(latestTimestamp, ",")
// 	exp := [7]string{
// 		".+",
// 		".+",
// 		"[0-9]+[.][0-9]+[(µs)(s)]",
// 		"[A-Za-z]{3}",
// 		"[0-9]{2}\\s[A-Za-z]{3}\\s[0-9]{4}\\s[0-9]{2}[:][0-9]{2}[:][0-9]{2}\\s[A-Z]{3}",
// 		"[A-Za-z]{3}",
// 		"[0-9]{2}\\s[A-Za-z]{3}\\s[0-9]{4}\\s[0-9]{2}[:][0-9]{2}[:][0-9]{2}\\s[A-Z]{3}",
// 	}
// 	for i := 0; i < 7; i++ {
// 		match, err := regexp.MatchString(exp[i], data[i])
// 		assert.NoError(err)
// 		assert.True(match)
// 	}
// }
