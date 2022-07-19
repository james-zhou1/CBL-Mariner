// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Records the run time for different parts of a go program
// and its nested calls to other go programs.

package timestamp

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/microsoft/CBL-Mariner/toolkit/tools/internal/logger"
)

var (
	Stamp = &TimeInfo{} // A shared TimeInfo object that is by default empty; will log an warning if the empty object is called.
)

// TimeInfo holds information needed for timestamping a go program.
type TimeInfo struct {
	filePath   string        // Path to store all timestamps
	toolName   string        // Name of the tool (consistent for all timestamps related to this object)
	stepName   string        // Name of the step
	actionName string        // Subaction within current step
	duration   time.Duration // Time to complete the step (ms)
	startTime  time.Time     // Start time of the step
	endTime    time.Time     // End time for the step
	timeRange  bool          // Whether to record start and end time
}

// Create a new instance of timeInfo struct.
func New(toolName string) *TimeInfo {
	return &TimeInfo{
		toolName:  toolName,
		timeRange: true,
		startTime: time.Now(),
	}
}

// Creates the file that every subsequent timestamp in this go program will write to.
func InitCSV(completePath string) {
	// Update the global object "Stamp".
	// Assume the base directory of completePath ends with .csv for now (possible to be .json later).
	fileName := filepath.Base(completePath)
	Stamp = New(fileName)

	file, err := os.Create(completePath)
	if err != nil {
		logger.Log.Warnf("Unable to create file %s: %s \n", completePath, err)
	}

	// Store file path information.
	Stamp.filePath = completePath
	file.Close()
}

// Start recording time for a new operation.
func (info *TimeInfo) Start() {
	// If we have not set up TimeInfo, log a warning and do nothing.
	if *info == (TimeInfo{}) {
		logger.Log.Warnf("Unable to record timestamp; 'timestamp.Stamp' has not been set up for this file.")
		return
	}
	info.startTime = time.Now()
}

// An internal function that helps record the timestamp.
func (info *TimeInfo) track() {
	info.endTime = time.Now()
	info.duration = info.endTime.Sub(info.startTime)
}

// Records a new timestamp and outputs it through the io.Writer specified in the input.
func (info *TimeInfo) RecordToFile(stepName string, actionName string, writer io.Writer) {
	info.track()
	info.stepName = stepName
	info.actionName = actionName

	// Generates the message.
	msg := info.stepName + " " + info.actionName + " in " + info.toolName + " took " + info.duration.String() + ". "
	if info.timeRange {
		msg += "Started at " + info.startTime.Format(time.UnixDate) + "; ended at " + info.endTime.Format(time.UnixDate) + ". \n"
	} else {
		msg += "\n"
	}
	_, err := io.WriteString(writer, msg)
	if err != nil {
		logger.Log.Warnf("Fail to write to file. %s\n", err)
	}

	// In case .start() is not called.
	info.startTime = info.endTime
}

// Records a new timestamp and writes it to the corresponding csv file.
func (info *TimeInfo) RecordToCSV(stepName string, actionName string) {
	fmt.Printf("executing recordToCSV() \n")
	// If we have not set up TimeInfo, log a warning and do nothing.
	if *info == (TimeInfo{}) {
		logger.Log.Warnf("Unable to record timestamp; 'timestamp.Stamp' has not been set up for this file.")
		return
	}
	// Create a new .csv file.
	file, err := os.OpenFile(info.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Log.Warnf("Failed to open the csv file. %s\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Writes the timestamp.
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
		logger.Log.Warnf("Fail to write to file. %s\n", err)
	}

	// In case .start() is not called.
	info.startTime = info.endTime
}
