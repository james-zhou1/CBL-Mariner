package timestamp

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

var (
	Stamp *TimeInfo
)

type TimeInfo struct {
	filePath   string        // Path to store all timestamps.
	toolName   string        // Name of the tool (consistent for all timestamps related to this object)
	stepName   string        // Name of the step
	actionName string        // Subaction within current step
	duration   time.Duration // Time to complete the step (ms)
	startTime  time.Time     // Start time of the step
	endTime    time.Time     // End time for the step
	timeRange  bool          // Whether to record start and end time
}

// Create a new instance of timeInfo struct.
func New(toolName string, timeRange bool) *TimeInfo {
	return &TimeInfo{
		toolName:  toolName,
		timeRange: timeRange,
		startTime: time.Now(),
	}
}

/*
 * Creates the file that every subsequent timestamp in this go program will write to.
 * Input:
 *	 completePath: A string representing the absolute path where all of the timestamps will be stored.
 *	 timeRange: A boolean that will record the start and end time of a timestamp interval if set to true.
 */
func InitCSV(completePath string, timeRange bool) {

	// Update the global object "Stamp".
	// assume the base directory of completePath ends with .csv for now (possible to be .json later).
	fileName := filepath.Base(completePath)
	fmt.Println(fileName)
	Stamp = New(fileName, timeRange)

	file, err := os.Create(completePath)
	if err != nil {
		fmt.Printf(completePath)
		panic(err)
	}

	// Store file path information.
	Stamp.filePath = completePath
	file.Close()
}

/*
 * Another possible option is to imput the step names at both .start() and .record().
 * If the two don't match, then wipe out the time recorded in timeInfo.startTime and
 * only record the finish time of the task.
 */
// Start recording time for a new operation.
func (info *TimeInfo) Start() {
	info.startTime = time.Now()
}

// An internal function that helps record the timestamp.
func (info *TimeInfo) track() {
	info.endTime = time.Now()
	info.duration = info.endTime.Sub(info.startTime)
}

// make a class output io.Writer
func (info *TimeInfo) RecordToFile(stepName string, actionName string, writer io.Writer) {
	info.track()
	info.stepName = stepName
	info.actionName = actionName
	msg := info.stepName + " " + info.actionName + " in " + info.toolName + " took " + info.duration.String() + ". "
	if info.timeRange {
		msg += "Started at " + info.startTime.Format(time.UnixDate) + "; ended at " + info.endTime.Format(time.UnixDate) + ". \n"
	} else {
		msg += "\n"
	}
	_, err := io.WriteString(writer, msg)
	if err != nil {
		panic(err)
	}

	// In case .start() is not called
	info.startTime = info.endTime
}

// go tool for csv files (for future parsing), tool name, step name, time, flag for time range
func (info *TimeInfo) RecordToCSV(stepName string, actionName string) {
	// Create a new .csv file. Should I add os.O_CREATE tag here?
	file, err := os.OpenFile(info.filePath, os.O_APPEND|os.O_WRONLY, 0777) // not sure what 0644 means but it works
	if err != nil {
		fmt.Printf("Failed to open the csv file. %s\n", err)
		return
	}
	defer file.Close()

	// Create a new csv writer.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Run
	info.track()
	info.stepName = stepName
	info.actionName = actionName
	if info.timeRange {
		err = writer.Write([]string{info.toolName, info.stepName, info.actionName, info.duration.String(),
			info.startTime.Format(time.UnixDate), info.endTime.Format(time.UnixDate)})
	} else {
		err = writer.Write([]string{info.toolName, info.stepName, info.actionName, info.duration.String()})
	}
	if err != nil {
		fmt.Printf("Fail to write to file. %s\n", err)
	}

	// In case .start() is not called
	info.startTime = info.endTime
}
