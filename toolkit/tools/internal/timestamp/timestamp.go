package timestamp

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
	"time"
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
func (info *TimeInfo) InitCSV(filePath string) {
	// Path subject to change later (to build folder?).
	completePath := "tools/internal/timestamp/results/" + filePath + ".csv" // this line is the actual completePath
	// completePath := filePath + ".csv" // this line is tor testing only
	// Create file.
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	err := os.MkdirAll(filepath.Dir(completePath), 0777)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(completePath)
	if err != nil {
		panic(err)
	}

	// Store file path information.
	info.filePath = completePath
	file.Close()

	// file, err := os.OpenFile(filePath + ".csv", os.O_CREATE | os.O_RDWR, 0644) // not sure what 0644 means but it works
	// if err != nil {
	// 	fmt.Printf("Failed to open the csv file. %s\n", err)
	// 	return
	// }
	// file.Close()
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

func (info *TimeInfo) track() {
	info.endTime = time.Now()
	info.duration = info.endTime.Sub(info.startTime)
	// result := timeInfo{toolName, stepName, "", diff.String(), start.Format(time.RFC1123), end.Format(time.RFC1123), timeRange}
}

// output as a string
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
// create a global instance of timestamp (initialized along with the logger)
// features: initialize timestamp, flag each run (start & end, wipe out? )
// make each of these function a method of the timeInfo struct
