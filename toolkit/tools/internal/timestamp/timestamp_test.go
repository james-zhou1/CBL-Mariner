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
)

//TestMain found in configuration_test.go.

func Test_WritetoFile_range_instant(t *testing.T) {
	TrackToFile(time.Now(), "test tool", "test step", true, os.Stdout)
}

func Test_WritetoFile_noRange_instant(t *testing.T) {
	TrackToFile(time.Now(), "test tool", "test step", false, os.Stdout)
}

func Test_WritetoFile_range_sleeps(t *testing.T) {
	defer TrackToFile(time.Now(), "tool 1", "step 1", true, os.Stdout)
	time.Sleep(3 * time.Second)
}

func WritetoCSV(seconds time.Duration) {
	defer TrackToCSV(time.Now(), "test tool", "test step", true)
	time.Sleep(seconds * time.Second)
}

func Test_WritetoCSV_Delay(t *testing.T) {
	WritetoCSV(0)
	WritetoCSV(1)
	WritetoCSV(3)
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
	if newLines != count {
		t.Fail()
	}
}

//	Run debug test to see print output in debug console.
func Test_WritetoCSV_MultipleLines(t *testing.T) {
	WritetoCSV_MultipleLines(1, t)
	WritetoCSV_MultipleLines(3, t)
}

func GetLatestTimestamp() string {
	file, _ := os.Open("build-time.csv")
	fileScanner := bufio.NewScanner(file)
	lastLine := ""
	for fileScanner.Scan() {
		lastLine = fileScanner.Text()
	}
	return lastLine
}

func WritetoCSV_timingTest(time time.Duration, t *testing.T) {
	WritetoCSV(time)
	latestTimestamp := GetLatestTimestamp()
	data := strings.Split(latestTimestamp, ",")
	match, err := regexp.MatchString("3.[0-9]{9}s", data[2]) // TODO: Make the timing test work for non-three second intervals
	if !match || err != nil {
		t.Fail()
	}
}

func Test_WritetoCSV_timingTest(t *testing.T) {
	WritetoCSV_timingTest(3, t)
}

func Test_WritetoCSV_formatTest(t *testing.T) {
	WritetoCSV(3)
	latestTimestamp := GetLatestTimestamp()
	data := strings.Split(latestTimestamp, ",")
	match, err := regexp.MatchString(".+", data[0])
	if !match || err != nil {
		t.Fail()
	}
	match, err = regexp.MatchString(".+", data[1])
	if !match || err != nil {
		t.Fail()
	}
	match, err = regexp.MatchString("[0-9]+[.][0-9]+[(µs)(s)]", data[2])
	if !match || err != nil {
		t.Fail()
	}
	match, err = regexp.MatchString("[A-Za-z]{3}", data[3])
	if !match || err != nil {
		t.Fail()
	}
	match, err = regexp.MatchString("[0-9]{2}\\s[A-Za-z]{3}\\s[0-9]{4}\\s[0-9]{2}[:][0-9]{2}[:][0-9]{2}\\s[A-Z]{3}", data[4])
	if !match || err != nil {
		t.Fail()
	}
	match, err = regexp.MatchString("[A-Za-z]{3}", data[5])
	if !match || err != nil {
		t.Fail()
	}
	match, err = regexp.MatchString("[0-9]{2}\\s[A-Za-z]{3}\\s[0-9]{4}\\s[0-9]{2}[:][0-9]{2}[:][0-9]{2}\\s[A-Z]{3}", data[6])
	if !match || err != nil {
		t.Fail()
	}
}
