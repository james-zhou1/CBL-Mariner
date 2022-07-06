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
	// currentProgress	int	// Proportion of the current task
	// maxProgress		int	// Maximum progress for parent level
}

// Create a new instance of timeInfo struct.
func New(toolName string, timeRange bool) *TimeInfo {
	return &TimeInfo{
		toolName:  toolName,
		timeRange: timeRange,
		startTime: time.Now(),
	}
}

// Creates the file that every preceding log in this go program will write to.
// Is this function necessary...?
func InitCSV(toolName string, timeRange bool) {
	// Path subject to change later (to build folder?).
	completePath := toolName + ".csv" 
	
	// Update the global object "Stamp".
	Stamp = New(toolName, timeRange)

	// mask := syscall.Umask(0)
	// defer syscall.Umask(mask)

	// If statement will not be triggered if the user is testing an executable file in a sub-folder.
	currDir, _ := os.Getwd()
	if currDir[len(currDir) - 19 : ] == "CBL-Mariner/toolkit" {
		// An image-build is probably running. 
		completePath = "tools/internal/timestamp/results/" + completePath
		err := os.MkdirAll(filepath.Dir(completePath), 0644)
		if err != nil {
		panic(err)
		}
	}

	file, err := os.Create(completePath)
	if err != nil {
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
	// curr := track(start, toolName, stepName, timeRange)
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
	// curr := track(start, toolName, stepName, timeRange)
	info.track()
	info.stepName = stepName
	info.actionName = actionName
	if info.timeRange {
		err = writer.Write([]string{info.toolName, info.stepName, info.actionName, info.duration.String(),
			info.startTime.Format(time.UnixDate), info.endTime.Format(time.UnixDate)})
		// err = writer.Write([]string{info.toolName, info.stepName, info.actionName, info.duration.String(),
		// 	info.startTime.Format(time.RFC1123), info.endTime.Format(time.RFC1123)})
	} else {
		err = writer.Write([]string{info.toolName, info.stepName, info.actionName, info.duration.String()})
	}
	if err != nil {
		fmt.Printf("Fail to write to file. %s\n", err)
	}

	// In case .start() is not called
	info.startTime = info.endTime
}

// output sth in the trace level?
// figure out logger package (how to call logger everywhere without passing a parameter)

// next step:
// features: initialize timestamp, flag each run (start & end, wipe out? )
// make each of these function a method of the timeInfo struct
// change csv destination to build/logs/csvlogs (?)
