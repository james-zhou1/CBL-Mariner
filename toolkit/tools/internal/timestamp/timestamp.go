package timestamp

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

type timeInfo struct {
	toolName  string // Name of the tool
	stepName  string // Name of the step
	duration  string // Time to complete the step (ms)
	start     string // Start time of the step
	end       string // End time for the step
	timeRange bool   // Whether to record start and end time
}

// var (
// 	data = []timeInfo{}
// )

// Call at the begining of the main() of a tool using "defer timestamp(time.Now(), "name_of_function")".
// The tool needs to import "time" too.
func track(start time.Time, toolName string, stepName string, timeRange bool) timeInfo {
	end := time.Now()
	diff := end.Sub(start)
	result := timeInfo{toolName, stepName, diff.String(), start.Format(time.RFC1123), end.Format(time.RFC1123), timeRange}
	return result
}

// output as a string
// make a class output io.Writer
func TrackToFile(start time.Time, toolName string, stepName string, timeRange bool, writer io.Writer) {
	curr := track(start, toolName, stepName, timeRange)
	msg := "Step " + stepName + " in " + toolName + " took " + curr.duration + ". "
	if timeRange {
		msg += "Started at " + curr.start + "; ended at " + curr.end + ". "
	}
	_, err := io.WriteString(writer, msg)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Wrote %d bytes\n", n)

}

// go tool for csv files (for future parsing), tool name, step name, time, flag for time range
func TrackToCSV(start time.Time, toolName string, stepName string, timeRange bool) {
	// Create a new .csv file.
	file, err := os.Create("build-time.csv") // this step will be moved to the init stage later
	if err != nil {
		fmt.Printf("Failed to create the csv file. %s\n", err)
	}
	defer file.Close()

	// Create a new csv writer.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Run
	curr := track(start, toolName, stepName, timeRange)
	if timeRange {
		err = writer.Write([]string{curr.toolName, curr.stepName, curr.duration, curr.start, curr.end})
	} else {
		err = writer.Write([]string{curr.toolName, curr.stepName, curr.duration})
	}
	if err != nil {
		fmt.Printf("Fail to write to file. %s\n", err)
	}
}

// output sth in the trace level?
// figure out logger package (how to call logger everywhere without passing a parameter)

// next step:
// create a global instance of timestamp (initialized along with the logger)
// features: initialize timestamp, flag each run (start & end, wipe out? )
// make each of these function a method of the timeInfo struct
