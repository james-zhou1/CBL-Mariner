package timestamp

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type timeInfo struct {
	toolName	string	// Name of the tool
	stepName	string	// Name of the step
	duration	int		// Time to complete the step (ms)
	start		string	// Start time of the step
	end		string		// End time for the step
	timeRange	bool	// Whether to record start and end time
}

var(
	// data = [][]string{
	// 	{"Tool Name", "Step Name", "Duration", "Start", "End"},
	// }
	data = []timeInfo{}
)

// Call at the begining of the main() of a tool using "defer timestamp(time.Now(), "name_of_function")".
// The tool needs to import "time" too.
func Track(start time.Time, name string) string {
	end := time.Now()
	diff := end.Sub(start)
	result := fmt.Sprintf("%s took %dms. Started at %s and ended at %s", name, diff.Nanoseconds()/1000, start, end)
	fmt.Printf(result)
	return result
}

// output as a string
// make a class output io.Writer
func TrackToFile(start time.Time, name string, writer io.Writer) {
	msg := track(start, name)
	n, err := io.WriteString(writer, msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote %d bytes\n", n)

}

// features: initialize timestamp, flag each run (start & end, wipe out? ),
// go tool for csv files (for future parsing), tool name, step name, time, flag for time range
func TrackToCSV(start time.Time, toolName string, stepName string, timeRange bool) {
	file, err := os.Create("build-time.csv")
	if err != nil {
		fmt.Println("Failed to create the csv file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.flush()

}
// output sth in the trace level?
// create struct that
// figure out logger package (how to call logger everywhere without passing a parameter)
